package decrypt

// import (
// 	"database/sql"
// 	"fmt"
// 	"strings"

// 	server "github.com/FISCO-BCOS/go-sdk/backend"
// 	"github.com/FISCO-BCOS/go-sdk/conf"
// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/sirupsen/logrus"
// )

// type SqlCtr struct {
// 	db *sql.DB
// }

// type EncodedInvoiceInformation struct {
// 	SqlId  string
// 	Num    string
// 	Status string
// 	Id     string
// 	Data   string
// 	Key    string
// 	Hash   string
// }

// type EncodedHistoricaltransactionInformation struct {
// 	SqlId  string
// 	Num    string
// 	Status string
// 	Id     string
// 	Data   string
// 	Key    string
// 	Hash   string
// }

// type EncodedPushingToPool struct {
// 	SqlId  string
// 	Num    string
// 	Status string
// 	Id     string
// 	Data   string
// 	Key    string
// 	Hash   string
// }

// type EncodedSupplierFinancingApplication struct {
// 	SqlId  string
// 	Num    string
// 	Status string
// 	Id     string
// 	Data   string
// 	Key    string
// 	Hash   string
// }

// type EncodedPushPaymentAccounts struct {
// 	SqlId  string
// 	Num    string
// 	Status string
// 	Id     string
// 	Data   string
// 	Key    string
// 	Hash   string
// }

// func NewSqlCtr() *SqlCtr {
// 	configs, err := conf.ParseConfigFile("./configs/config.toml")
// 	if err != nil {
// 		logrus.Fatalln(err)
// 	}
// 	config := &configs[0]
// 	// db, err := sql.Open("mysql", "root:123456@/db_node0")
// 	str := config.MslUsername + ":" + config.MslPasswd + "@/" + config.MslName
// 	db, err := sql.Open("mysql", str)
// 	if err != nil {
// 		logrus.Fatalln(err)
// 	}
// 	return &SqlCtr{
// 		db: db,
// 	}
// }

// func (s *SqlCtr) QueryTablesByOrder(order string) ([]string, error) {
// 	in_stmt, err := s.db.Prepare(order)
// 	if err != nil {
// 		return nil, err
// 	}
// 	rows, err := in_stmt.Query()
// 	if err != nil {
// 		return nil, err
// 	}
// 	if strings.Contains(order, "u_t_invoice_information") {
// 		ret := make([]string, 0)
// 		//ret := make([]*EncodedInvoiceInformation, 0)
// 		count := 0
// 		i := 0
// 		for rows.Next() {
// 			record := &EncodedInvoiceInformation{}
// 			err = rows.Scan(&record.SqlId, &record.Num, &record.Status, &record.Id, &record.Data, &record.Key, &record.Hash)
// 			if err != nil {
// 				count++
// 				continue
// 			}
// 			serve := server.NewServer()
// 			//解密对称密钥
// 			symkey, err := serve.DecryptSymkey([]byte(record.Key))
// 			if err != nil {
// 				logrus.Infof("利用私钥解密对称密钥失败")
// 				//fmt.Println(err)
// 			}
// 			//解密数据
// 			data, err := serve.DecryptData(record.Data, symkey)
// 			if err != nil {
// 				logrus.Infof("利用对称密钥解密数据失败")
// 			}
// 			fmt.Println(string(data))
// 			if serve.ValidateHash([]byte(record.Hash), data) {
// 				ret = append(ret, string(data))
// 			} else {
// 				logrus.Infof("哈希值与数据对应错误")
// 			}
// 			//fmt.Println(ret[i])
// 			i = i + 1
// 		}
// 		return ret, nil
// 	} else if strings.Contains(order, "u_t_historical_transaction_information") {
// 		var ret []string
// 		count := 0
// 		i := 0
// 		for rows.Next() {
// 			record := &EncodedHistoricaltransactionInformation{}
// 			err = rows.Scan(&record.SqlId, &record.Num, &record.Status, &record.Id, &record.Data, &record.Key, &record.Hash)
// 			if err != nil {
// 				count++
// 				continue
// 			}
// 			serve := server.NewServer()
// 			//解密对称密钥
// 			symkey, err := serve.DecryptSymkey([]byte(record.Key))
// 			if err != nil {
// 				logrus.Infof("利用私钥解密对称密钥失败")
// 			}
// 			//解密数据
// 			data, err := serve.DecryptData(record.Data, symkey)
// 			if err != nil {
// 				logrus.Infof("利用对称密钥解密数据失败")
// 			}
// 			fmt.Println(string(data))
// 			if serve.ValidateHash([]byte(record.Hash), data) {
// 				ret = append(ret, string(data))
// 			} else {
// 				logrus.Infof("哈希值与数据对应错误")
// 			}
// 			i = i + 1
// 		}
// 		return ret, nil
// 	} else if strings.Contains(order, "u_t_pushing_to_pool") {
// 		var ret []string
// 		count := 0
// 		i := 0
// 		for rows.Next() {
// 			record := &EncodedPushingToPool{}
// 			err = rows.Scan(&record.SqlId, &record.Num, &record.Status, &record.Id, &record.Data, &record.Key, &record.Hash)
// 			if err != nil {
// 				count++
// 				continue
// 			}
// 			serve := server.NewServer()
// 			//解密对称密钥
// 			symkey, err := serve.DecryptSymkey([]byte(record.Key))
// 			if err != nil {
// 				logrus.Infof("利用私钥解密对称密钥失败")
// 				//fmt.Println(err)
// 			}
// 			//解密数据
// 			data, err := serve.DecryptData(record.Data, symkey)
// 			if err != nil {
// 				logrus.Infof("利用对称密钥解密数据失败")
// 			}
// 			fmt.Println(string(data))
// 			if serve.ValidateHash([]byte(record.Hash), data) {
// 				ret = append(ret, string(data))
// 			} else {
// 				logrus.Infof("哈希值与数据对应错误")
// 			}
// 			i = i + 1
// 		}
// 		return ret, nil
// 	} else if strings.Contains(order, "u_t_supplier_financing_application") {
// 		var ret []string
// 		count := 0
// 		i := 0
// 		for rows.Next() {
// 			record := &EncodedSupplierFinancingApplication{}
// 			err = rows.Scan(&record.SqlId, &record.Num, &record.Status, &record.Id, &record.Data, &record.Key, &record.Hash)
// 			if err != nil {
// 				count++
// 				continue
// 			}
// 			serve := server.NewServer()
// 			//解密对称密钥
// 			symkey, err := serve.DecryptSymkey([]byte(record.Key))
// 			if err != nil {
// 				logrus.Infof("利用私钥解密对称密钥失败")
// 				//fmt.Println(err)
// 			}
// 			//解密数据
// 			data, err := serve.DecryptData(record.Data, symkey)
// 			if err != nil {
// 				logrus.Infof("利用对称密钥解密数据失败")
// 			}
// 			fmt.Println(string(data))
// 			if serve.ValidateHash([]byte(record.Hash), data) {
// 				ret = append(ret, string(data))
// 			} else {
// 				logrus.Infof("哈希值与数据对应错误")
// 			}
// 			//fmt.Println(ret[i])
// 			i = i + 1
// 		}
// 		return ret, nil
// 	} else if strings.Contains(order, "u_t_push_payment_accounts") {
// 		var ret []string
// 		count := 0
// 		i := 0
// 		for rows.Next() {
// 			record := &EncodedPushPaymentAccounts{}
// 			err = rows.Scan(&record.SqlId, &record.Num, &record.Status, &record.Id, &record.Data, &record.Key, &record.Hash)
// 			if err != nil {
// 				count++
// 				continue
// 			}
// 			serve := server.NewServer()
// 			//解密对称密钥
// 			symkey, err := serve.DecryptSymkey([]byte(record.Key))
// 			if err != nil {
// 				logrus.Infof("利用私钥解密对称密钥失败")
// 				//fmt.Println(err)
// 			}
// 			//解密数据
// 			data, err := serve.DecryptData(record.Data, symkey)
// 			if err != nil {
// 				logrus.Infof("利用对称密钥解密数据失败")
// 			}
// 			fmt.Println(string(data))
// 			if serve.ValidateHash([]byte(record.Hash), data) {
// 				ret = append(ret, string(data))
// 			} else {
// 				logrus.Infof("哈希值与数据对应错误")
// 			}
// 			//fmt.Println(ret[i])
// 			i = i + 1
// 		}
// 		return ret, nil
// 	}
// 	// switch order {
// 	// case "select * from u_t_invoice_information":
// 	// case "select * from u_t_historical_transaction_information":
// 	// case "select * from u_t_pushing_to_pool":
// 	// case "select * from u_t_supplier_financing_application":
// 	// case "select * from u_t_push_payment_accounts":
// 	// }
// 	return nil, nil
// }
