package querytable

import (
	"fmt"
	"net/http"
	"reflect"

	decrypt "github.com/FISCO-BCOS/go-sdk/decodeMysql"
	sql "github.com/FISCO-BCOS/go-sdk/sqlController"
)

func DecryptInvoiceInformation(writer http.ResponseWriter, request *http.Request) {
	decrypte := decrypt.NewSqlCtr()
	Sql := sql.NewSqlCtr()
	slice := Sql.InvoiceInformationIndex(request)
	immutable := reflect.ValueOf(slice).Elem()
	id := immutable.FieldByName("id").String()
	ret := decrypte.QueryInvoiceInformation("u_t_invoice_information", id)
	jsonData := decrypte.ConvertoStruct("InvoiceInformation", ret)
	fmt.Fprint(writer, jsonData)
}

func DecryptHistoricaltransactionUsedinfos(writer http.ResponseWriter, request *http.Request) {
	decrypte := decrypt.NewSqlCtr()
	ret, _ := decrypte.QueryTablesByOrder2("select * from u_t_history_used_information")
	jsonData := decrypte.ConvertoStruct("HistoricaltransactionUsedinfos", ret)
	fmt.Fprint(writer, jsonData)
}

func DecryptHistoricaltransactionSettleinfos(writer http.ResponseWriter, request *http.Request) {
	decrypte := decrypt.NewSqlCtr()
	ret, _ := decrypte.QueryTablesByOrder2("select * from u_t_history_settle_information")
	jsonData := decrypte.ConvertoStruct("HistoricaltransactionSettleinfos", ret)
	fmt.Fprint(writer, jsonData)
}

func DecryptHistoricaltransactionOrderinfos(writer http.ResponseWriter, request *http.Request) {
	decrypte := decrypt.NewSqlCtr()
	ret, _ := decrypte.QueryTablesByOrder2("select * from u_t_history_order_information")
	jsonData := decrypte.ConvertoStruct("HistoricaltransactionOrderinfos", ret)
	fmt.Fprint(writer, jsonData)
}

func DecryptHistoricaltransactionReceivableinfos(writer http.ResponseWriter, request *http.Request) {
	decrypte := decrypt.NewSqlCtr()
	ret, _ := decrypte.QueryTablesByOrder2("select * from u_t_history_receivable_information")
	jsonData := decrypte.ConvertoStruct("HistoricaltransactionReceivableinfos", ret)
	fmt.Fprint(writer, jsonData)
}

func DecryptEnterpoolDataPlaninfos(writer http.ResponseWriter, request *http.Request) {
	decrypte := decrypt.NewSqlCtr()
	ret, _ := decrypte.QueryTablesByOrder2("select * from u_t_pool_plan_information")
	jsonData := decrypte.ConvertoStruct("EnterpoolDataPlaninfos", ret)
	fmt.Fprint(writer, jsonData)
}

func DecryptEnterpoolDataUsedinfos(writer http.ResponseWriter, request *http.Request) {
	decrypte := decrypt.NewSqlCtr()
	ret, _ := decrypte.QueryTablesByOrder2("select * from u_t_pool_used_information")
	jsonData := decrypte.ConvertoStruct("EnterpoolDataUsedinfos", ret)
	fmt.Fprint(writer, jsonData)
}

func DecryptFinancingIntention(writer http.ResponseWriter, request *http.Request) {
	decrypte := decrypt.NewSqlCtr()
	Sql := sql.NewSqlCtr()
	slice := Sql.FinancingIntentionIndex(request)
	immutable := reflect.ValueOf(slice).Elem()
	id := immutable.FieldByName("id").String()
	ret := decrypte.QueryFinancingIntention("u_t_supplier_financing_application", id)
	jsonData := decrypte.ConvertoStruct("FinancingIntention", ret)
	fmt.Fprint(writer, jsonData)
}

func DecryptCollectionAccount(writer http.ResponseWriter, request *http.Request) {
	decrypte := decrypt.NewSqlCtr()
	Sql := sql.NewSqlCtr()
	slice := Sql.CollectionAccountIndex(request)
	immutable := reflect.ValueOf(slice).Elem()
	id := immutable.FieldByName("id").String()
	ret := decrypte.QueryCollectionAccount("u_t_push_payment_accounts", id)
	jsonData := decrypte.ConvertoStruct("CollectionAccount", ret)
	fmt.Fprint(writer, jsonData)
}

// func DecryptInvoiceInformation(decrypte *decrypt.SqlCtr) {
// 	ret, _ := decrypte.QueryTablesByOrder("select * from u_t_invoice_information")
// 	// fmt.Println(len(ret))
// 	decrypte.ConvertoStruct("InvoiceInformation", ret)
// }

// func DecryptHistoricaltransactionUsedinfos(decrypte *decrypt.SqlCtr) {
// 	ret, _ := decrypte.QueryTablesByOrder2("select * from u_t_history_used_information")
// 	decrypte.ConvertoStruct("HistoricaltransactionUsedinfos", ret)
// }

// func DecryptHistoricaltransactionSettleinfos(decrypte *decrypt.SqlCtr) {
// 	ret, _ := decrypte.QueryTablesByOrder2("select * from u_t_history_settle_information")
// 	decrypte.ConvertoStruct("HistoricaltransactionSettleinfos", ret)
// }

// func DecryptHistoricaltransactionOrderinfos(decrypte *decrypt.SqlCtr) {
// 	ret, _ := decrypte.QueryTablesByOrder2("select * from u_t_history_order_information")
// 	decrypte.ConvertoStruct("HistoricaltransactionOrderinfos", ret)
// }

// func DecryptHistoricaltransactionReceivableinfos(decrypte *decrypt.SqlCtr) {
// 	ret, _ := decrypte.QueryTablesByOrder2("select * from u_t_history_receivable_information")
// 	decrypte.ConvertoStruct("HistoricaltransactionReceivableinfos", ret)
// }

// func DecryptEnterpoolDataPlaninfos(decrypte *decrypt.SqlCtr) {
// 	ret, _ := decrypte.QueryTablesByOrder2("select * from u_t_pool_plan_information")
// 	decrypte.ConvertoStruct("EnterpoolDataPlaninfos", ret)
// }

// func DecryptEnterpoolDataUsedinfos(decrypte *decrypt.SqlCtr) {
// 	ret, _ := decrypte.QueryTablesByOrder2("select * from u_t_pool_used_information")
// 	decrypte.ConvertoStruct("EnterpoolDataUsedinfos", ret)
// }

// func DecryptFinancingIntention(decrypte *decrypt.SqlCtr) {
// 	ret, _ := decrypte.QueryTablesByOrder("select * from u_t_supplier_financing_application")
// 	decrypte.ConvertoStruct("FinancingIntention", ret)
// }

// func DecryptCollectionAccount(decrypte *decrypt.SqlCtr) {
// 	ret, _ := decrypte.QueryTablesByOrder("select * from u_t_push_payment_accounts")
// 	decrypte.ConvertoStruct("CollectionAccount", ret)
// }
