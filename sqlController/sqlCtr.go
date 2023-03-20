package sql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/FISCO-BCOS/go-sdk/conf"
	types "github.com/FISCO-BCOS/go-sdk/type"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

type SqlCtr struct {
	db    *sql.DB
	logDB *sql.DB
}

type LogData struct {
	Id        string
	Timestamp string
	Type      string
	Info      string
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

	str1 := config.LogDBUsername + ":" + config.LogDBPasswd + "@/" + config.LogDBName
	logdb, err := sql.Open("mysql", str1)
	if err != nil {
		logrus.Fatalln(err)
	}
	return &SqlCtr{
		db:    db,
		logDB: logdb,
	}
}

// //////////////////////////////////////////////////////////////////////////////////////////////////
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

func (s *SqlCtr) InvoiceinfoToMap(ret []string) (map[string]map[string]string, error) {
	invoiceInfoStruct := handleInvoiceInfo(ret)
	ans, err := json.Marshal(invoiceInfoStruct)
	if err != nil {
		return nil, err
	}
	input := string(ans)
	var data []map[string]string
	if err := json.Unmarshal([]byte(input), &data); err != nil {
		return nil, err
	}

	result := make(map[string]map[string]string)

	for _, item := range data {
		key1 := item["certificateId"]
		key2 := item["customerId"]
		key3 := item["corpName"]
		key4 := item["certificateType"]
		key5 := item["interCustomerId"]
		delete(item, "certificateId")
		delete(item, "customerId")
		delete(item, "corpName")
		delete(item, "certificateType")
		delete(item, "interCustomerId")
		result[key1+"|"+key2+"|"+key3+"|"+key4+"|"+key5] = item
	}
	return result, nil
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

///////////////////////////////////////////////////////////////////////////////////////////////////////////////

// 插入日志数据
func (s *SqlCtr) InsertLogs(Timestamp string, Type string, Info string) error {
	_, err := s.logDB.Exec("insert into u_t_log(timestamp, type, info) values (?,?,?)", Timestamp, Type, Info)
	if err != nil {
		// fmt.Printf("insert failed, err: %v\n", err)
		return err
	}
	return nil
}

// 查询日志数据
func (s *SqlCtr) QueryLogsByOrder(order string) interface{} {
	in_stmt, err := s.logDB.Prepare(order)
	if err != nil {
		logrus.Panicln(err)
	}
	rows, err := in_stmt.Query()
	if err != nil {
		return err
	}
	switch order {
	case "select * from u_t_log":
		ret := make([]*LogData, 0)
		count := 0
		i := 0
		for rows.Next() {
			record := &LogData{}
			err = rows.Scan(&record.Id, &record.Timestamp, &record.Type, &record.Info)
			//fmt.Println(err)
			if err != nil {
				count++
				continue
			}
			ret = append(ret, record)
			// fmt.Println(ret[i])
			i = i + 1
		}
		logrus.Infof("select %d information from u_t_log and error counts is %x", len(ret), count)
		return ret
	}
	return nil
}

// 根据时间戳，获取当前时间前一小时以内的数据
func (s *SqlCtr) ExtractLogByHour() (interface{}, int) {
	currentTime := time.Now()         //当前时间
	m, _ := time.ParseDuration("-1h") //当前时间往前推1小时
	checktime := currentTime.Add(m)
	checktime_str := checktime.String()
	timeTemplate1 := "2006-01-02 15:04:05"                                                //改变时间戳模板， 方便与下面的recordtime比较
	checktime1, _ := time.ParseInLocation(timeTemplate1, checktime_str[0:19], time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	// fmt.Println("checktime is: ", checktime1)
	in_stmt, err := s.logDB.Prepare("select * from u_t_log")
	if err != nil {
		logrus.Panicln(err)
	}
	rows, err := in_stmt.Query()
	if err != nil {
		return err, 0
	}
	ret := make([]*LogData, 0)
	count := 0
	i := 0
	fmt.Printf("最近一小时以内的数据: \n")
	for rows.Next() {
		record := &LogData{}
		err = rows.Scan(&record.Id, &record.Timestamp, &record.Type, &record.Info)
		//fmt.Println(err)
		if err != nil {
			count++
			continue
		}
		recordtime, _ := time.ParseInLocation(timeTemplate1, record.Timestamp, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
		// fmt.Println("recordtime is: ", recordtime)
		if recordtime.After(checktime1) {
			ret = append(ret, record)
			fmt.Println(ret[i])
			i = i + 1
		}
	}
	return ret, i
}
