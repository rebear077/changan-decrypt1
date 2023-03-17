package receive

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/sirupsen/logrus"
)

type FrontEnd struct {
	InvoicePool             []*InvoiceInformation
	TransactionHistoryPool  []*TransactionHistory
	EnterpoolDataPool       []*EnterpoolData
	FinancingIntentionPool  []*FinancingIntention
	CollectionAccountPool   []*CollectionAccount
	invoicemutex            sync.RWMutex
	transactionHistorymutex sync.RWMutex
	enterpoolDatamutex      sync.RWMutex
	financingIntentionmutex sync.RWMutex
	collectionAccountmutex  sync.RWMutex
}

func NewFrontEnd() *FrontEnd {
	return &FrontEnd{
		InvoicePool:            make([]*InvoiceInformation, 0),
		TransactionHistoryPool: make([]*TransactionHistory, 0),
		EnterpoolDataPool:      make([]*EnterpoolData, 0),
		FinancingIntentionPool: make([]*FinancingIntention, 0),
		CollectionAccountPool:  make([]*CollectionAccount, 0),
	}
}
func (f *FrontEnd) HandleInvoiceInformation(writer http.ResponseWriter, request *http.Request) {
	jsonData := []byte(`{
		"msg":"",
		"result": "{}",
		"code":"SUC000000"
	}`)
	len := request.ContentLength
	body := make([]byte, len)
	// request.Body.Read(body)
	body, _ = ioutil.ReadAll(request.Body)
	var message InvoiceInformation
	err := json.Unmarshal(body, &message)
	Check(err)
	f.invoicemutex.Lock()
	f.InvoicePool = append(f.InvoicePool, &message)
	f.invoicemutex.Unlock()
	fmt.Fprintln(writer, string(jsonData))
}

// 推送历史交易信息接口
func (f *FrontEnd) HandleTransactionHistory(writer http.ResponseWriter, request *http.Request) {
	jsonData := []byte(`{
		"msg":"",
		"result": "{}",
		"code":"SUC000000"
	}`)
	len := request.ContentLength
	body := make([]byte, len)
	// request.Body.Read(body)
	body, _ = ioutil.ReadAll(request.Body)
	var message TransactionHistory
	err := json.Unmarshal(body, &message)
	Check(err)
	f.transactionHistorymutex.Lock()
	f.TransactionHistoryPool = append(f.TransactionHistoryPool, &message)
	f.transactionHistorymutex.Unlock()
	fmt.Fprintln(writer, string(jsonData))
}

// 推送入池数据接口
func (f *FrontEnd) HandleEnterpoolData(writer http.ResponseWriter, request *http.Request) {
	jsonData := []byte(`{
		"msg":"",
		"result": "{}",
		"code":"SUC000000"
	}`)
	len := request.ContentLength
	body := make([]byte, len)
	// request.Body.Read(body)
	body, _ = ioutil.ReadAll(request.Body)
	var message EnterpoolData
	err := json.Unmarshal(body, &message)
	Check(err)
	f.enterpoolDatamutex.Lock()
	f.EnterpoolDataPool = append(f.EnterpoolDataPool, &message)
	f.enterpoolDatamutex.Unlock()
	fmt.Fprintln(writer, string(jsonData))
}

// 提交融资意向接口
func (f *FrontEnd) HandleFinancingIntention(writer http.ResponseWriter, request *http.Request) {
	jsonData := []byte(`{
		"msg":"",
		"result": "{}",
		"code":"SUC000000"
	}`)
	len := request.ContentLength
	body := make([]byte, len)
	// request.Body.Read(body)
	body, _ = ioutil.ReadAll(request.Body)
	var message FinancingIntention
	err := json.Unmarshal(body, &message)
	Check(err)
	f.financingIntentionmutex.Lock()
	f.FinancingIntentionPool = append(f.FinancingIntentionPool, &message)
	f.financingIntentionmutex.Unlock()
	fmt.Fprintln(writer, string(jsonData))
}

// 推送回款账户接口
func (f *FrontEnd) HandleCollectionAccount(writer http.ResponseWriter, request *http.Request) {
	jsonData := []byte(`{
		"msg":"",
		"result": "{}",
		"code":"SUC000000"
	}`)
	len := request.ContentLength
	body := make([]byte, len)
	// request.Body.Read(body)
	body, _ = ioutil.ReadAll(request.Body)
	var message CollectionAccount
	err := json.Unmarshal(body, &message)
	Check(err)
	f.collectionAccountmutex.Lock()
	f.CollectionAccountPool = append(f.CollectionAccountPool, &message)
	f.collectionAccountmutex.Unlock()
	fmt.Fprintln(writer, string(jsonData))
}

// func main() {
// 	http.HandleFunc("/asl/universal/push-invoice-info", handleInvoiceInformation)
// 	http.HandleFunc("/asl/universal/push-history", handleTransactionHistory)
// 	http.HandleFunc("/asl/universal/caqc/push-inpool", handleEnterpoolData)
// 	http.HandleFunc("/asl/universal/commmit-intention", handleFinancingIntention)
// 	http.HandleFunc("/asl/universal/back-account-lock", handleCollectionAccount)
// 	err := http.ListenAndServeTLS(":8443", "./server.pem", "./server.key", nil)
// 	//err := http.ListenAndServe(":8443", nil)
// 	if err != nil {
// 		log.Fatalf("启动 HTTPS 服务器失败: %v", err)
// 	}
// }

func Check(err error) {
	if err != nil {
		logrus.Fatalln(err)
	}
}
