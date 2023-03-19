package decrypt

import (
	"database/sql"
	"fmt"

	"github.com/FISCO-BCOS/go-sdk/conf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

type SqlCtr struct {
	db *sql.DB
}

type EncodedInformation struct {
	SqlId  string
	Num    string
	Status string
	Id     string
	Data   string
	Key    string
	Hash   string
}

type EncodedInformation2 struct {
	SqlId  string
	Num    string
	Status string
	Id     string
	Time   string
	Data   string
	Key    string
	Hash   string
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
	return &SqlCtr{
		db: db,
	}
}

func (s *SqlCtr) QueryInvoiceInformation(tablename string, id string) []string {
	var ret []string
	if id == "" {
		ret, _ = s.QueryTablesByOrder("select * from u_t_invoice_information")
	} else {
		ret, _ = s.QueryTablesByOrder("select * from u_t_invoice_information where id = " + id)
	}
	return ret
}

func (s *SqlCtr) QueryFinancingIntention(tablename string, id string) []string {
	var ret []string
	if id == "" {
		ret, _ = s.QueryTablesByOrder("select * from u_t_supplier_financing_application")
	} else {
		ret, _ = s.QueryTablesByOrder("select * from u_t_supplier_financing_application where id = " + id)
	}
	return ret
}

func (s *SqlCtr) QueryCollectionAccount(tablename string, id string) []string {
	var ret []string
	if id == "" {
		ret, _ = s.QueryTablesByOrder("select * from u_t_push_payment_accounts")
	} else {
		ret, _ = s.QueryTablesByOrder("select * from u_t_push_payment_accounts where id = " + id)
	}
	return ret
}

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
		record := &EncodedInformation{}
		err = rows.Scan(&record.SqlId, &record.Num, &record.Status, &record.Id, &record.Data, &record.Key, &record.Hash)
		if err != nil {
			fmt.Println(err)
			count++
			continue
		}
		serve := NewDecrypter()
		//解密对称密钥
		// fmt.Println("key: ", []byte(record.Key))
		// fmt.Println("hash: ", record.Hash)
		symkey, err := serve.DecryptSymkey([]byte(record.Key))
		if err != nil {
			logrus.Infof("利用私钥解密对称密钥失败")
		}
		//解密数据
		data, err := serve.DecryptData(record.Data, symkey)
		if err != nil {
			logrus.Infof("利用对称密钥解密数据失败")
		}
		//fmt.Println(string(data))
		if serve.ValidateHash([]byte(record.Hash), data) {
			ret = append(ret, string(data))
		} else {
			logrus.Infof("哈希值与数据对应错误")
		}
		//fmt.Println(ret[i])
		i = i + 1
	}
	//fmt.Println(ret)
	return ret, nil
}

func (s *SqlCtr) QueryTablesByOrder2(order string) ([]string, error) {
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
		record := &EncodedInformation2{}
		err = rows.Scan(&record.SqlId, &record.Num, &record.Status, &record.Id, &record.Time, &record.Data, &record.Key, &record.Hash)
		if err != nil {
			fmt.Println(err)
			count++
			continue
		}
		serve := NewDecrypter()
		symkey, err := serve.DecryptSymkey([]byte(record.Key))
		if err != nil {
			logrus.Infof("利用私钥解密对称密钥失败")
		}
		data, err := serve.DecryptData(record.Data, symkey)
		if err != nil {
			logrus.Infof("利用对称密钥解密数据失败")
		}
		if serve.ValidateHash([]byte(record.Hash), data) {
			ret = append(ret, string(data))
		} else {
			logrus.Infof("哈希值与数据对应错误")
		}
		i = i + 1
	}
	return ret, nil
}
