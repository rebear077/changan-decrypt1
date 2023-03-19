package main

import (
	"log"
	"net/http"

	"github.com/FISCO-BCOS/go-sdk/querytable"
)

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
