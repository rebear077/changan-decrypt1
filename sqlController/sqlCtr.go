package sql

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/FISCO-BCOS/go-sdk/conf"
	types "github.com/FISCO-BCOS/go-sdk/type"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

type SqlCtr struct {
	db        *sql.DB
	Decrypter *Decrypter
}

func NewSqlCtr() *SqlCtr {
	configs, err := conf.ParseConfigFile("./configs/config.toml")
	if err != nil {
		logrus.Fatalln(err)
	}
	config := &configs[0]
	// db, err := sql.Open("mysql", "root:123456@/db_node0")
	str := config.MslUsername + ":" + config.MslPasswd + "@/" + config.MslName
	db, err := sql.Open("mysql", str)
	if err != nil {
		logrus.Fatalln(err)
	}
	de := NewDecrypter()
	return &SqlCtr{
		db:        db,
		Decrypter: de,
	}
}

// //////////////////////////////////////////////////////////////////////////////////////////////////
// 解析前端发来的URL请求，获取检索条件，结构体形式返回
func (s *SqlCtr) InvoiceInformationIndex(request *http.Request) *types.InvoiceInformationSearch {
	query := request.URL.Query()
	id := ""
	if len(query["id"]) > 0 {
		id = query["id"][0]
	}
	index := types.InvoiceInformationSearch{
		Id: id,
	}
	return &index
}

// 输入参数是解密后的发票信息，转换成redis存储所需要的数据结构
func (s *SqlCtr) InvoiceinfoToMap(ret []string) []*types.InvoiceInformation {

	// ans, err := json.Marshal(invoiceInfoStruct)
	// if err != nil {
	// 	return nil, err
	// }
	// input := string(ans)
	// var data []map[string]string
	// if err := json.Unmarshal([]byte(input), &data); err != nil {
	// 	return nil, err
	// }
	// result := make(map[string]map[string]string)

	// for _, item := range data {
	// 	key1 := item["certificateId"]
	// 	key2 := item["customerId"]
	// 	key3 := item["corpName"]
	// 	key4 := item["certificateType"]
	// 	key5 := item["interCustomerId"]
	// 	delete(item, "certificateId")
	// 	delete(item, "customerId")
	// 	delete(item, "corpName")
	// 	delete(item, "certificateType")
	// 	delete(item, "interCustomerId")
	// 	result[key1+"|"+key2+"|"+key3+"|"+key4+"|"+key5] = item
	// }
	return handleInvoiceInfo(ret)
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (s *SqlCtr) FinancingIntentionIndex(request *http.Request) *types.FinancingIntentionSearch {
	query := request.URL.Query()
	id := ""
	if len(query["id"]) > 0 {
		id = query["id"][0]
	}
	index := types.FinancingIntentionSearch{
		Id: id,
	}
	return &index
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (s *SqlCtr) CollectionAccountIndex(request *http.Request) *types.CollectionAccountSearch {
	query := request.URL.Query()
	id := ""
	if len(query["id"]) > 0 {
		id = query["id"][0]
	}
	index := types.CollectionAccountSearch{
		Id: id,
	}
	return &index
}

// 查询mysql数据库中加密后的发票信息，如果id为空，则查找全部的信息
func (s *SqlCtr) QueryInvoiceInformation(id string) []string {
	var ret []string
	if id == "" {
		ret, _ = s.QueryTablesByOrder("select * from u_t_invoice_information")
	} else {
		ret, _ = s.QueryTablesByOrder("select * from u_t_invoice_information where id = " + id)
	}
	return ret
}
func (s *SqlCtr) QueryInvoiceInformationLength() int {
	var length int
	err := s.db.QueryRow("select count(*) from u_t_invoice_information").Scan(&length)
	if err != nil {
		logrus.Fatalln(err)
	}
	return length
}

// 查询mysql数据库中融资意向信息，如果id为空，则查找全部的信息
func (s *SqlCtr) QueryFinancingIntention(id string) []string {
	var ret []string
	if id == "" {
		ret, _ = s.QueryTablesByOrder("select * from u_t_supplier_financing_application")
	} else {
		ret, _ = s.QueryTablesByOrder("select * from u_t_supplier_financing_application where id = " + id)
	}
	return ret
}

// 查询mysql数据库中回款账户信息，如果id为空，则查找全部的信息
func (s *SqlCtr) QueryCollectionAccount(id string) []string {
	var ret []string
	if id == "" {
		ret, _ = s.QueryTablesByOrder("select * from u_t_push_payment_accounts")
	} else {
		ret, _ = s.QueryTablesByOrder("select * from u_t_push_payment_accounts where id = " + id)
	}
	return ret
}

// 输入命令，比如“select * from u_t_push_payment_accounts”,查询出加密后的密文然后自动解密，返回明文[]string
func (s *SqlCtr) QueryTablesByOrder(order string) ([]string, error) {
	in_stmt, err := s.db.Prepare(order)
	if err != nil {
		return nil, err
	}
	rows, err := in_stmt.Query()
	if err != nil {
		return nil, err
	}
	ret := make([]string, 0)
	count := 0
	i := 0
	for rows.Next() {
		record := &types.RawSQLData{}
		err = rows.Scan(&record.SQLId, &record.Num, &record.Status, &record.ID, &record.Data, &record.Key, &record.Hash)
		if err != nil {
			fmt.Println(err)
			count++
			continue
		}
		symkey, err := s.Decrypter.DecryptSymkey([]byte(record.Key))
		if err != nil {
			logrus.Errorln("利用私钥解密对称密钥失败")
		}
		data, err := s.Decrypter.DecryptData(record.Data, symkey)
		if err != nil {
			logrus.Errorln("利用对称密钥解密数据失败")
		}
		if s.Decrypter.ValidateHash([]byte(record.Hash), data) {
			ret = append(ret, string(data))
		} else {
			logrus.Errorln("哈希值与数据对应错误")
		}
		i = i + 1
	}
	return ret, nil
}
