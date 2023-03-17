package sql

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/FISCO-BCOS/go-sdk/conf"
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

type InvoiceInformation struct {
	id int
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

func (s *SqlCtr) InvoiceInformationSearchKey(request *http.Request) InvoiceInformation {
	query := request.URL.Query()
	id := -1
	if len(query["id"]) > 0 {
		id, _ = strconv.Atoi(query["id"][0])
	}
	var index InvoiceInformation
	return index{
		id: id,
	}
}

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

// import (
// 	"database/sql"
// 	"fmt"
// 	"io"
// 	"os"

// 	"github.com/FISCO-BCOS/go-sdk/conf"
// 	_ "github.com/go-sql-driver/mysql"
// 	_ "github.com/godror/godror"
// 	"github.com/sirupsen/logrus"
// )

// type SqlCtr struct {
// 	// db *sql.DB
// }

// type supplier struct {
// 	id        int64
// 	num       string
// 	idnum     string
// 	name      string
// 	taskNum   string
// 	taskPrice int64
// 	moneytype string
// 	start     string
// 	end       string
// 	piaohao   string
// }

// func (s *SqlCtr) QueryFromOracle() {
// 	//这里可以采用上面const里面的数据，我为了方便直接放在里面了。192.168.1.11:1521/peixun这个是数据库所在服务器的ip:端口/数据库上面的实例名
// 	db, err := sql.Open("godror", `user="scott" password="123456" connectString="localhost:1521/orcl"`)
// 	if err != nil {
// 		fmt.Println("123")
// 		panic(err)
// 	}
// 	defer db.Close()
// 	err = db.Ping()
// 	if err != nil {
// 		fmt.Println("1234")
// 		fmt.Println(err.Error())
// 		panic(err)
// 	}
// 	// str := "select * from table_feedback_on_lending_results"
// 	str := "select rownum,table_feedback_on_lending_results.* from table_feedback_on_lending_results"
// 	res, err := db.Query(str)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	for res.Next() {
// 		var s supplier
// 		res.Scan(&s.id, &s.num, &s.idnum, &s.name, &s.taskNum, &s.taskPrice, &s.moneytype, &s.start, &s.end, &s.piaohao)
// 		fmt.Println(s)
// 	}
// }

// :password@/dbname
// func NewSqlCtr() *SqlCtr {
// 	return &SqlCtr{}
// configs, err := conf.ParseConfigFile("configs/config.toml")
// if err != nil {
// 	logrus.Fatalln(err)
// }
// config := &configs[0]
// str := config.MslUsername + ":" + config.MslPasswd + "@/" + config.MslName
// db, err := sql.Open("mysql", str)
// if err != nil {
// 	logrus.Fatalln(err)
// }
// return &SqlCtr{
// 	db: db,
// }
// }

// func (s *SqlCtr) QueryPublicKey(role string) {
// 	in_stmt, err := s.db.Prepare("select * from t_public_key")
// 	if err != nil {
// 		logrus.Panicln(err)
// 	}
// 	rows, err := in_stmt.Query()
// 	if err != nil {
// 		return
// 	}
// 	ret := make([]*PublicKey, 0)
// 	count := 0
// 	for rows.Next() {
// 		record := &PublicKey{}

// 		err = rows.Scan(&record.SId, &record.Num, &record.Status, &record.Id, &record.Role, &record.Key)
// 		if err != nil {
// 			count++
// 			continue
// 		}
// 		ret = append(ret, record)
// 	}
// 	for _, r := range ret {
// 		if r.Role == role {
// 			file, _ := os.Create("../configs/" + role + "Public.pem")
// 			_, err := io.WriteString(file, r.Key)
// 			if err != nil {
// 				logrus.Errorln("write public.pem error")
// 			}
// 		}
// 	}
// }
// func (s *SqlCtr) QueryAllRecordsByOrder(order string) interface{} {
// 	in_stmt, err := s.db.Prepare(order)
// 	if err != nil {
// 		logrus.Panicln(err)
// 	}

// 	rows, err := in_stmt.Query()
// 	if err != nil {
// 		return err
// 	}
// 	switch order {
// 	case "select * from u_t_supplier_financing_application":
// 		ret := make([]*EncryptedProviderissueFinancingData, 0)
// 		count := 0
// 		for rows.Next() {
// 			record := &EncryptedProviderissueFinancingData{}

// 			err = rows.Scan(&record.SId, &record.Num, &record.Status, &record.Id, &record.Data, &record.Key, &record.Hash)
// 			if err != nil {
// 				count++
// 				continue
// 			}
// 			ret = append(ret, record)
// 		}
// 		logrus.Infof("select %d information from u_t_supplier_financing_application and error counts is %d", len(ret), count)
// 		return ret
// 	case "select * from u_t_invoice_information":
// 		ret := make([]*EncryptedReceiptInformationData, 0)
// 		count := 0
// 		for rows.Next() {
// 			record := &EncryptedReceiptInformationData{}

// 			err = rows.Scan(&record.SId, &record.Num, &record.Status, &record.Id, &record.Data, &record.Key, &record.Hash)
// 			if err != nil {
// 				count++
// 				continue
// 			}
// 			ret = append(ret, record)
// 		}
// 		logrus.Infof("select %d information from u_t_invoice_information and error counts is %d", len(ret), count)
// 		return ret
// 	case "select * from u_t_historical_transaction_information":
// 		ret := make([]*EncryptedTransactionHistory, 0)
// 		count := 0
// 		for rows.Next() {
// 			record := &EncryptedTransactionHistory{}

// 			err = rows.Scan(&record.SId, &record.Num, &record.Status, &record.Id, &record.Data, &record.Key, &record.Hash)
// 			if err != nil {
// 				count++
// 				continue
// 			}
// 			ret = append(ret, record)
// 		}
// 		logrus.Infof("select %d information from u_t_historical_transaction_information and error counts is %d", len(ret), count)
// 		return ret
// 	case "select * from u_t_weekly_trading_record":
// 		ret := make([]*EncryptedWeektransactionRecord, 0)
// 		count := 0
// 		for rows.Next() {
// 			record := &EncryptedWeektransactionRecord{}

// 			err = rows.Scan(&record.SId, &record.Num, &record.Status, &record.Id, &record.Data, &record.Key, &record.Hash)
// 			if err != nil {
// 				count++
// 				continue
// 			}
// 			ret = append(ret, record)
// 		}
// 		logrus.Infof("select %d information from u_t_weekly_trading_record and error counts is %d", len(ret), count)
// 		return ret
// 	case "select * from u_t_financing_result_feedback":
// 		ret := make([]*EncryptedFinancingResult, 0)
// 		count := 0
// 		for rows.Next() {
// 			record := &EncryptedFinancingResult{}

// 			err = rows.Scan(&record.SId, &record.Num, &record.Status, &record.Id, &record.Data, &record.Key, &record.Hash)
// 			if err != nil {
// 				count++
// 				continue
// 			}
// 			ret = append(ret, record)
// 		}
// 		logrus.Infof("select %d information from u_t_financing_result_feedback and error counts is %d", len(ret), count)
// 		return ret
// 	case "select * from u_t_feedback_on_lending_results":
// 		ret := make([]*EncryptedLoanResult, 0)
// 		count := 0
// 		for rows.Next() {
// 			record := &EncryptedLoanResult{}

// 			err = rows.Scan(&record.SId, &record.Num, &record.Status, &record.Id, &record.Data, &record.Key, &record.Hash)
// 			if err != nil {
// 				count++
// 				continue
// 			}
// 			ret = append(ret, record)
// 		}
// 		logrus.Infof("select %d information from u_t_feedback_on_lending_results and error counts is %d", len(ret), count)
// 		return ret
// 	case "select * from u_t_repayment_information":
// 		ret := make([]*EncryptedRepayResult, 0)
// 		count := 0
// 		for rows.Next() {
// 			record := &EncryptedRepayResult{}

// 			err = rows.Scan(&record.SId, &record.Num, &record.Status, &record.Id, &record.Data, &record.Key, &record.Hash)
// 			if err != nil {
// 				count++
// 				continue
// 			}
// 			ret = append(ret, record)
// 		}
// 		logrus.Infof("select %d information from u_t_repayment_information and error counts is %d", len(ret), count)
// 		return ret
// 	}
// 	return nil

// }
