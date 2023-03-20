package server

import (
	"context"
	"encoding/json"
	"time"

	"github.com/FISCO-BCOS/go-sdk/canal"
	"github.com/FISCO-BCOS/go-sdk/redis"
	sql "github.com/FISCO-BCOS/go-sdk/sqlController"
	types "github.com/FISCO-BCOS/go-sdk/type"
	"github.com/sirupsen/logrus"
)

type Server struct {
	sql   *sql.SqlCtr
	redis *redis.RedisOperator
	canal *canal.Connector
}

// 初始化
func NewServer() *Server {
	sqlCtr := sql.NewSqlCtr()
	redis := redis.NewRedisOperator()
	canal := canal.NewConnector("db_node1\\.u_t_invoice_information")
	return &Server{
		sql:   sqlCtr,
		redis: redis,
		canal: canal,
	}
}

// 将redis的数据全部删除后，与mysql中的数据进行同步
func (s *Server) ForceSynchronous() {
	//首先清除redis数据
	ctx := context.Background()
	s.redis.FlushData(ctx)
	//获取mysql中的数据,获取的是解密后的明文
	plaintext := s.sql.QueryInvoiceInformation("")
	invoices := s.sql.InvoiceinfoToMap(plaintext)
	s.StoreInvoicesToRedis(invoices)
}

// 存储发票信息到redis数据库中
func (s *Server) StoreInvoicesToRedis(data []*types.InvoiceInformation) {
	ctx := context.Background()
	for _, invoice := range data {
		values := make(map[string]interface{})
		key := invoice.Customerid + ":" + invoice.Checkcode
		values["Certificateid"] = invoice.Certificateid
		values["Customerid"] = invoice.Customerid
		values["Corpname"] = invoice.Corpname
		values["Certificatetype"] = invoice.Certificatetype
		values["Intercustomerid"] = invoice.Intercustomerid
		values["Invoicenotaxamt"] = invoice.Invoicenotaxamt
		values["Invoiceccy"] = invoice.Invoiceccy
		values["Sellername"] = invoice.Sellername
		values["Invoicetype"] = invoice.Invoicetype
		values["Buyername"] = invoice.Buyername
		values["Buyerusccode"] = invoice.Buyerusccode
		values["Invoicedate"] = invoice.Invoicedate
		values["Sellerusccode"] = invoice.Sellerusccode
		values["Invoicecode"] = invoice.Invoicecode
		values["Invoicenum"] = invoice.Invoicenum
		values["Checkcode"] = invoice.Checkcode
		values["Invoiceamt"] = invoice.Invoiceamt
		err := s.redis.MultipleSet(ctx, key, values)
		if err != nil {
			logrus.Errorln(err)
		}
	}
}

// 从canal中进行同步消息到redis中
func (s *Server) DumpFromCanal() {
	go s.canal.Start()
	for {
		s.canal.Lock.RLock()
		if len(s.canal.RawData) != 0 {
			messages := make([]*types.RawSQLData, 0)
			messages = append(messages, s.canal.RawData...)
			s.canal.RawData = nil
			s.canal.Lock.RUnlock()
			s.StoreEncryptedToredis(messages)
		} else {
			s.canal.Lock.RUnlock()
			time.Sleep(1 * time.Second)
		}
	}
}

// 从数据库原始的数据，先解密，然后转换格式后存入redis中
func (s *Server) StoreEncryptedToredis(datas []*types.RawSQLData) {
	ret := make([]string, 0)
	for _, raw := range datas {
		symkey, err := s.sql.Decrypter.DecryptSymkey([]byte(raw.Key))
		if err != nil {
			logrus.Errorln("利用私钥解密对称密钥失败")
			continue
		}
		data, err := s.sql.Decrypter.DecryptData(raw.Data, symkey)
		if err != nil {
			logrus.Errorln("利用对称密钥解密数据失败")
			continue
		}
		if s.sql.Decrypter.ValidateHash([]byte(raw.Hash), data) {
			ret = append(ret, string(data))
		} else {
			logrus.Errorln("哈希值与数据对应错误")
			continue
		}
	}
	invoices := s.sql.InvoiceinfoToMap(ret)
	s.StoreInvoicesToRedis(invoices)
}
func (s *Server) SearchFromRedis(order map[string]string) []*types.InvoiceInformation {

	invoices := s.searchByIDFromRedis(order["id"])
	//redis未命中
	if len(invoices) == 0 {
		//同步mysql到redis
		s.DumpFromMysqlToRedis(order["id"])
		time.Sleep(500 * time.Millisecond)
		//二次查询
		invoices = s.searchByIDFromRedis(order["id"])
		if len(invoices) == 0 {
			return nil
		}
	}
	fliterBytype := s.fliterByInvoiceType(invoices, order["invoiceType"])
	fliterByTime := s.fliterByInvoiceTimeStamp(fliterBytype, order["time"])
	return fliterByTime
}

// 根据id的信息从redis中查询数据，如果结构体是空的，那么说明redis未命中，需要去mysql数据库中查询
func (s *Server) searchByIDFromRedis(id string) []*types.InvoiceInformation {
	ctx := context.Background()
	invoices := make([]*types.InvoiceInformation, 0)
	keys := s.GetMutipleKeys(id)
	//如果redis未命中,返回空的结构体
	if len(keys) == 0 {
		return nil
	}
	for _, key := range keys {
		resmap, err := s.redis.GetAll(ctx, key)
		if err != nil {
			logrus.Errorln(err)
			continue
		}
		invoice := packToInvoiceStruct(resmap)
		invoices = append(invoices, invoice)
	}
	return invoices
}

// 根据发票类型过滤，调用此函数前，需要先通过id进行第一次检索
func (s *Server) fliterByInvoiceType(messages []*types.InvoiceInformation, invocietype string) []*types.InvoiceInformation {
	if invocietype == "" {
		return messages
	}
	result := make([]*types.InvoiceInformation, 0)
	for _, message := range messages {
		if message.Invoicetype == invocietype {
			result = append(result, message)
		}
	}
	return result
}

// 根据发票的时间戳进行过滤，调用此函数前，需要先通过id进行第一次检索
func (s *Server) fliterByInvoiceTimeStamp(messages []*types.InvoiceInformation, invoiceTimeStamp string) []*types.InvoiceInformation {
	// time1 := "2015-03-20 08:50:29"
	// time2 := "2015-03-21 09:04:25"
	// //先把时间字符串格式化成相同的时间类型
	// t1, err := time.Parse("2006-01-02 15:04:05", time1)
	// t2, err := time.Parse("2006-01-02 15:04:05", time2)
	// if err == nil && t1.Before(t2) {
	// 	//处理逻辑
	// 	fmt.Println("true")
	// }
	//todo
	if invoiceTimeStamp == "" {
		return messages
	}
	return messages

}

// 将从redis查询出来的数据转换成结构体
func packToInvoiceStruct(message map[string]string) *types.InvoiceInformation {
	invoice := new(types.InvoiceInformation)
	invoice.Certificateid = message["Certificateid"]
	invoice.Customerid = message["Customerid"]
	invoice.Corpname = message["Corpname"]
	invoice.Certificatetype = message["Certificatetype"]
	invoice.Intercustomerid = message["Intercustomerid"]
	invoice.Invoicenotaxamt = message["Invoicenotaxamt"]
	invoice.Invoiceccy = message["Invoiceccy"]
	invoice.Sellername = message["Sellername"]
	invoice.Invoicetype = message["Invoicetype"]
	invoice.Buyername = message["Buyername"]
	invoice.Buyerusccode = message["Buyerusccode"]
	invoice.Invoicedate = message["Invoicedate"]
	invoice.Sellerusccode = message["Sellerusccode"]
	invoice.Invoicecode = message["Invoicecode"]
	invoice.Invoicenum = message["Invoicenum"]
	invoice.Checkcode = message["Checkcode"]
	invoice.Invoiceamt = message["Invoiceamt"]
	return invoice

}
func (s *Server) GetMutipleKeys(id string) []string {
	ctx := context.Background()
	order := id + "*"
	_, keys := s.redis.Scan(ctx, order)
	return keys
}

// redis未命中的情况下，去查询数据库中的数据，这种情况只适用于指定了id的情况，如果id未指定，则直接从redis数据库中返回信息
// 将mysql查询的数据首先存入redis，然后进行二次过滤
func (s *Server) DumpFromMysqlToRedis(id string) {
	plaintext := s.sql.QueryInvoiceInformation(id)
	invoices := s.sql.InvoiceinfoToMap(plaintext)
	s.StoreInvoicesToRedis(invoices)
}

func (s *Server) PackToJson(messages []*types.InvoiceInformation) string {
	ans, err := json.Marshal(messages)
	if err != nil {
		logrus.Errorln(ans)
	}
	return string(ans)

}
