package querytable

import (
	"net/http"

	server "github.com/FISCO-BCOS/go-sdk/backend"
)

type FrontEnd struct {
	server *server.Server
}

func (front *FrontEnd) DecryptInvoiceInformation(writer http.ResponseWriter, request *http.Request) {
	// decrypte := decrypt.NewSqlCtr()
	// Sql := sql.NewSqlCtr()
	// slice := Sql.InvoiceInformationIndex(request)
	// immutable := reflect.ValueOf(slice).Elem()
	// id := immutable.FieldByName("id").String()
	// ret := decrypte.QueryInvoiceInformation(id)

	// jsonData := decrypte.ConvertoStruct("InvoiceInformation", ret)
	// fmt.Fprint(writer, jsonData)
}

// func DecryptHistoricaltransactionUsedinfos(writer http.ResponseWriter, request *http.Request) {
// 	decrypte := decrypt.NewSqlCtr()
// 	ret, _ := decrypte.QueryTablesByOrder2("select * from u_t_history_used_information")
// 	jsonData := decrypte.ConvertoStruct("HistoricaltransactionUsedinfos", ret)
// 	fmt.Fprint(writer, jsonData)
// }

// func DecryptHistoricaltransactionSettleinfos(writer http.ResponseWriter, request *http.Request) {
// 	decrypte := decrypt.NewSqlCtr()
// 	ret, _ := decrypte.QueryTablesByOrder2("select * from u_t_history_settle_information")
// 	jsonData := decrypte.ConvertoStruct("HistoricaltransactionSettleinfos", ret)
// 	fmt.Fprint(writer, jsonData)
// }

// func DecryptHistoricaltransactionOrderinfos(writer http.ResponseWriter, request *http.Request) {
// 	decrypte := decrypt.NewSqlCtr()
// 	ret, _ := decrypte.QueryTablesByOrder2("select * from u_t_history_order_information")
// 	jsonData := decrypte.ConvertoStruct("HistoricaltransactionOrderinfos", ret)
// 	fmt.Fprint(writer, jsonData)
// }

// func DecryptHistoricaltransactionReceivableinfos(writer http.ResponseWriter, request *http.Request) {
// 	decrypte := decrypt.NewSqlCtr()
// 	ret, _ := decrypte.QueryTablesByOrder2("select * from u_t_history_receivable_information")
// 	jsonData := decrypte.ConvertoStruct("HistoricaltransactionReceivableinfos", ret)
// 	fmt.Fprint(writer, jsonData)
// }

// func DecryptEnterpoolDataPlaninfos(writer http.ResponseWriter, request *http.Request) {
// 	decrypte := decrypt.NewSqlCtr()
// 	ret, _ := decrypte.QueryTablesByOrder2("select * from u_t_pool_plan_information")
// 	jsonData := decrypte.ConvertoStruct("EnterpoolDataPlaninfos", ret)
// 	fmt.Fprint(writer, jsonData)
// }

// func DecryptEnterpoolDataUsedinfos(writer http.ResponseWriter, request *http.Request) {
// 	decrypte := decrypt.NewSqlCtr()
// 	ret, _ := decrypte.QueryTablesByOrder2("select * from u_t_pool_used_information")
// 	jsonData := decrypte.ConvertoStruct("EnterpoolDataUsedinfos", ret)
// 	fmt.Fprint(writer, jsonData)
// }

// func DecryptFinancingIntention(writer http.ResponseWriter, request *http.Request) {
// 	decrypte := decrypt.NewSqlCtr()
// 	Sql := sql.NewSqlCtr()
// 	slice := Sql.FinancingIntentionIndex(request)
// 	immutable := reflect.ValueOf(slice).Elem()
// 	id := immutable.FieldByName("id").String()
// 	ret := decrypte.QueryFinancingIntention(id)
// 	jsonData := decrypte.ConvertoStruct("FinancingIntention", ret)
// 	fmt.Fprint(writer, jsonData)
// }

// func DecryptCollectionAccount(writer http.ResponseWriter, request *http.Request) {
// 	decrypte := decrypt.NewSqlCtr()
// 	Sql := sql.NewSqlCtr()
// 	slice := Sql.CollectionAccountIndex(request)
// 	immutable := reflect.ValueOf(slice).Elem()
// 	id := immutable.FieldByName("id").String()
// 	ret := decrypte.QueryCollectionAccount(id)
// 	jsonData := decrypte.ConvertoStruct("CollectionAccount", ret)
// 	fmt.Fprint(writer, jsonData)
// }
