package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/withlin/canal-go/client"
	pbe "github.com/withlin/canal-go/protocol/entry"
)

func main() {
	// 192.168.199.17 替换成你的canal server的地址
	// example 替换成-e canal.destinations=example 你自己定义的名字
	connector := client.NewSimpleCanalConnector("127.0.0.1", 11111, "", "", "example", 60000, 60*60*1000)
	err := connector.Connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// https://github.com/alibaba/canal/wiki/AdminGuide
	//mysql 数据解析关注的表，Perl正则表达式.
	//
	//多个正则之间以逗号(,)分隔，转义符需要双斜杠(\\)
	//
	//常见例子：
	//
	//  1.  所有表：.*   or  .*\\..*
	//	2.  canal schema下所有表： canal\\..*
	//	3.  canal下的以canal打头的表：canal\\.canal.*
	//	4.  canal schema下的一张表：canal\\.test1
	//  5.  多个规则组合使用：canal\\..*,mysql.test1,mysql.test2 (逗号分隔)

	err = connector.Subscribe("db_node1\\.u_t_history_settle_information")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for {

		message, err := connector.Get(100, nil, nil)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		batchId := message.Id
		if batchId == -1 || len(message.Entries) <= 0 {
			time.Sleep(300 * time.Millisecond)
			fmt.Println("===没有数据了===")
			continue
		}

		printEntry(message.Entries)

	}

}

func printEntry(entrys []pbe.Entry) {
	fmt.Println(123)
	for _, entry := range entrys {
		fmt.Println(entry.GetEntryType())
		if entry.GetEntryType() == pbe.EntryType_TRANSACTIONBEGIN || entry.GetEntryType() == pbe.EntryType_TRANSACTIONEND {
			continue
		}
		rowChange := new(pbe.RowChange)

		err := proto.Unmarshal(entry.GetStoreValue(), rowChange)
		checkError(err)
		if rowChange != nil {
			eventType := rowChange.GetEventType()
			header := entry.GetHeader()
			fmt.Println(fmt.Sprintf("================> binlog[%s : %d],name[%s,%s], eventType: %s", header.GetLogfileName(), header.GetLogfileOffset(), header.GetSchemaName(), header.GetTableName(), header.GetEventType()))

			for _, rowData := range rowChange.GetRowDatas() {
				if eventType == pbe.EventType_DELETE {
					printColumn(rowData.GetBeforeColumns())
				} else if eventType == pbe.EventType_INSERT {
					printColumn(rowData.GetAfterColumns())
				} else {
					fmt.Println("-------> before")
					printColumn(rowData.GetBeforeColumns())
					fmt.Println("-------> after")
					printColumn(rowData.GetAfterColumns())
				}
			}
		}
	}
}

func printColumn(columns []*pbe.Column) {
	for _, col := range columns {
		fmt.Println(fmt.Sprintf("%s:%s", col.GetName(), col.GetValue()))
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

// func main() {
// 	decrypte := decrypt.NewSqlCtr()
// 	// querytable.DecryptEnterpoolDataPlaninfos(decrypte)
// 	// querytable.DecryptEnterpoolDataUsedinfos(decrypte)
// 	// querytable.DecryptHistoricaltransactionUsedinfos(decrypte)
// 	// querytable.DecryptHistoricaltransactionSettleinfos(decrypte)
// 	// querytable.DecryptHistoricaltransactionOrderinfos(decrypte)
// 	// querytable.DecryptHistoricaltransactionReceivableinfos(decrypte)
// 	querytable.DecryptInvoiceInformation(decrypte)
// 	// querytable.DecryptFinancingIntention(decrypte)
// 	// querytable.DecryptCollectionAccount(decrypte)
// }

<<<<<<< HEAD
func main() {
	// decrypte := decrypt.NewSqlCtr()
	// for {
	http.HandleFunc("/asl/universal/decryptInvoiceInformation", querytable.DecryptInvoiceInformation)
	http.HandleFunc("/asl/universal/decryptHistoricaltransactionUsedinfos", querytable.DecryptHistoricaltransactionUsedinfos)
	http.HandleFunc("/asl/universal/decryptHistoricaltransactionSettleinfos", querytable.DecryptHistoricaltransactionSettleinfos)
	http.HandleFunc("/asl/universal/decryptHistoricaltransactionOrderinfos", querytable.DecryptHistoricaltransactionOrderinfos)
	http.HandleFunc("/asl/universal/decryptHistoricaltransactionReceivableinfos", querytable.DecryptHistoricaltransactionReceivableinfos)
	http.HandleFunc("/asl/universal/decryptEnterpoolDataPlaninfos", querytable.DecryptEnterpoolDataPlaninfos)
	http.HandleFunc("/asl/universal/decryptEnterpoolDataUsedinfos", querytable.DecryptEnterpoolDataUsedinfos)
	http.HandleFunc("/asl/universal/decryptFinancingIntention", querytable.DecryptFinancingIntention)
	http.HandleFunc("/asl/universal/decryptCollectionAccount", querytable.DecryptCollectionAccount)

	// err := http.ListenAndServeTLS(":8440", "connApi/confs/server.pem", "connApi/confs/server.key", nil)
	err := http.ListenAndServe(":8440", nil)
	if err != nil {
		log.Fatalf("启动 HTTPS 服务器失败: %v", err)
	}
	// }
}
=======
// func main() {
// 	// decrypte := decrypt.NewSqlCtr()
// 	// for {
// 	http.HandleFunc("/asl/universal/decryptInvoiceInformation", querytable.DecryptInvoiceInformation)
// 	http.HandleFunc("/asl/universal/decryptHistoricaltransactionUsedinfos", querytable.DecryptHistoricaltransactionUsedinfos)
// 	http.HandleFunc("/asl/universal/decryptHistoricaltransactionSettleinfos", querytable.DecryptHistoricaltransactionSettleinfos)
// 	http.HandleFunc("/asl/universal/decryptHistoricaltransactionOrderinfos", querytable.DecryptHistoricaltransactionOrderinfos)
// 	http.HandleFunc("/asl/universal/decryptHistoricaltransactionReceivableinfos", querytable.DecryptHistoricaltransactionReceivableinfos)
// 	http.HandleFunc("/asl/universal/decryptEnterpoolDataPlaninfos", querytable.DecryptEnterpoolDataPlaninfos)
// 	http.HandleFunc("/asl/universal/decryptEnterpoolDataUsedinfos", querytable.DecryptEnterpoolDataUsedinfos)
// 	http.HandleFunc("/asl/universal/decryptFinancingIntention", querytable.DecryptFinancingIntention)
// 	http.HandleFunc("/asl/universal/decryptCollectionAccount", querytable.DecryptCollectionAccount)
// 	// err := http.ListenAndServeTLS(":8440", "connApi/confs/server.pem", "connApi/confs/server.key", nil)
// 	err := http.ListenAndServe(":8440", nil)
// 	if err != nil {
// 		log.Fatalf("启动 HTTPS 服务器失败: %v", err)
// 	}
// 	// }
// }
>>>>>>> 6c314fd21a311aab36902699feafcf5d131c648a
