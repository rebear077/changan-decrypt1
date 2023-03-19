package decrypt

import (
	"encoding/json"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

//发票信息推送接口
type InvoiceInformation struct {
	Certificateid   string `json:"certificateId"`
	Customerid      string `json:"customerId"`
	Corpname        string `json:"corpName"`
	Certificatetype string `json:"certificateType"`
	Intercustomerid string `json:"interCustomerId"`
	Invoicenotaxamt string `json:"InvoiceNotaxAmt"`
	Invoiceccy      string `json:"InvoiceCcy"`
	Sellername      string `json:"SellerName"`
	Invoicetype     string `json:"InvoiceType"`
	Buyername       string `json:"BuyerName"`
	Buyerusccode    string `json:"BuyerUsccode"`
	Invoicedate     string `json:"InvoiceDate"`
	Sellerusccode   string `json:"SellerUsccode"`
	Invoicecode     string `json:"InvoiceCode"`
	Invoicenum      string `json:"InvoiceNum"`
	Checkcode       string `json:"CheckCode"`
	Invoiceamt      string `json:"InvoiceAmt"`
}

// 推送历史交易信息接口
// type TransactionHistory struct {
// 	Customergrade   string            `json:"customerGrade"`
// 	Certificatetype string            `json:"certificateType"`
// 	Intercustomerid string            `json:"interCustomerId"`
// 	Corpname        string            `json:"corpName"`
// 	Financeid       string            `json:"financeId"`
// 	Certificateid   string            `json:"certificateId"`
// 	Customerid      string            `json:"customerId"`
// 	Usedinfos       []Usedinfos       `json:"usedInfos"`
// 	Settleinfos     []Settleinfos     `json:"settleInfos"`
// 	Orderinfos      []Orderinfos      `json:"orderInfos"`
// 	Receivableinfos []Receivableinfos `json:"receivableInfos"`
// }

type Usedinfos struct {
	Tradeyearmonth string `json:"TradeYearMonth"`
	Usedamount     string `json:"UsedAmount"`
	Ccy            string `json:"Ccy"`
}
type Settleinfos struct {
	Tradeyearmonth string `json:"TradeYearMonth"`
	Settleamount   string `json:"SettleAmount"`
	Ccy            string `json:"Ccy"`
}
type Orderinfos struct {
	Tradeyearmonth string `json:"TradeYearMonth"`
	Orderamount    string `json:"OrderAmount"`
	Ccy            string `json:"Ccy"`
}
type Receivableinfos struct {
	Tradeyearmonth   string `json:"TradeYearMonth"`
	Receivableamount string `json:"ReceivableAmount"`
	Ccy              string `json:"Ccy"`
}

type TransactionHistoryUsedinfos struct {
	Customergrade   string      `json:"customerGrade"`
	Certificatetype string      `json:"certificateType"`
	Intercustomerid string      `json:"interCustomerId"`
	Corpname        string      `json:"corpName"`
	Financeid       string      `json:"financeId"`
	Certificateid   string      `json:"certificateId"`
	Customerid      string      `json:"customerId"`
	Usedinfos       []Usedinfos `json:"usedInfos"`
}

type TransactionHistorySettleinfos struct {
	Customergrade   string        `json:"customerGrade"`
	Certificatetype string        `json:"certificateType"`
	Intercustomerid string        `json:"interCustomerId"`
	Corpname        string        `json:"corpName"`
	Financeid       string        `json:"financeId"`
	Certificateid   string        `json:"certificateId"`
	Customerid      string        `json:"customerId"`
	Settleinfos     []Settleinfos `json:"settleInfos"`
}

type TransactionHistoryOrderinfos struct {
	Customergrade   string       `json:"customerGrade"`
	Certificatetype string       `json:"certificateType"`
	Intercustomerid string       `json:"interCustomerId"`
	Corpname        string       `json:"corpName"`
	Financeid       string       `json:"financeId"`
	Certificateid   string       `json:"certificateId"`
	Customerid      string       `json:"customerId"`
	Orderinfos      []Orderinfos `json:"orderInfos"`
}

type TransactionHistoryReceivableinfos struct {
	Customergrade   string            `json:"customerGrade"`
	Certificatetype string            `json:"certificateType"`
	Intercustomerid string            `json:"interCustomerId"`
	Corpname        string            `json:"corpName"`
	Financeid       string            `json:"financeId"`
	Certificateid   string            `json:"certificateId"`
	Customerid      string            `json:"customerId"`
	Receivableinfos []Receivableinfos `json:"receivableInfos"`
}

type EnterpoolDataPlaninfos struct {
	Datetimepoint     string      `json:"dateTimePoint"`
	Ccy               string      `json:"ccy"`
	Customerid        string      `json:"customerId"`
	Intercustomerid   string      `json:"interCustomerId"`
	Receivablebalance string      `json:"receivableBalance"`
	Planinfos         []Planinfos `json:"planInfos"`
}

type EnterpoolDataProviderusedinfos struct {
	Datetimepoint     string              `json:"dateTimePoint"`
	Ccy               string              `json:"ccy"`
	Customerid        string              `json:"customerId"`
	Intercustomerid   string              `json:"interCustomerId"`
	Receivablebalance string              `json:"receivableBalance"`
	Providerusedinfos []Providerusedinfos `json:"ProviderUsedInfos"`
}

type Planinfos struct {
	Tradeyearmonth string `json:"TradeYearMonth"`
	Planamount     string `json:"PlanAmount"`
	Currency       string `json:"Currency"`
}
type Providerusedinfos struct {
	Tradeyearmonth string `json:"TradeYearMonth"`
	Usedamount     string `json:"UsedAmount"`
	Currency       string `json:"Currency"`
}

//提交融资意向接口
type FinancingIntention struct {
	Custcdlinkposition string `json:"CustcdLinkPosition"`
	Custcdlinkname     string `json:"CustcdLinkName"`
	Certificateid      string `json:"CertificateId"`
	Corpname           string `json:"CorpName"`
	Remark             string `json:"Remark"`
	Bankcontact        string `json:"BankContact"`
	Banklinkname       string `json:"BankLinkName"`
	Custcdcontact      string `json:"CustcdContact"`
	Customerid         string `json:"CustomerId"`
	Financeid          string `json:"FinanceId"`
	Cooperationyears   string `json:"CooperationYears"`
	Certificatetype    string `json:"CertificateType"`
	Intercustomerid    string `json:"InterCustomerId"`
}

//推送回款账户接口
type CollectionAccount struct {
	Backaccount     string `json:"BackAccount"`
	Certificateid   string `json:"CertificateId"`
	Customerid      string `json:"CustomerId"`
	Corpname        string `json:"CorpName"`
	Lockremark      string `json:"LockRemark"`
	Certificatetype string `json:"CertificateType"`
	Intercustomerid string `json:"InterCustomerId"`
}

func sliceinfohandler(str string) (string, string) {
	flag := 0
	header := ""
	infos := ""
	for index, val := range str {
		if index+1 >= len(str) {
			break
		}
		if flag == 0 {
			if str[index] == ',' && str[index+1] == '[' {
				flag = 1
			} else {
				header = header + string(val)
			}
		} else if flag == 1 {
			//应该是防止有[,]的情况，即子表单中无内容
			if str[index] == '[' && str[index+1] == ',' {
				flag = 2
			} else if str[index] == ']' {
				flag = 2
			} else if str[index] != '[' && str[index] != ']' {
				infos = infos + string(val)
			}
		} else if flag == 2 {
			break
		}
	}
	return header, infos
}

func handleInvoiceInformation(data []string) []InvoiceInformation {
	//如果其他输入中存在[]怎么办？
	//最后返回的结果，目前是结构体的切片
	var INV []InvoiceInformation
	for i := 0; i < len(data); i++ {
		str := data[i]
		//fmt.Println(str)
		str_split := strings.Split(str, ",")
		ICfo := InvoiceInformation{
			str_split[0],
			str_split[1],
			str_split[2],
			str_split[3],
			str_split[4],
			str_split[5],
			str_split[6],
			str_split[7],
			str_split[8],
			str_split[9],
			str_split[10],
			str_split[11],
			str_split[12],
			str_split[13],
			str_split[14],
			str_split[15],
			str_split[16],
		}
		INV = append(INV, ICfo)
	}
	// fmt.Println(INV)
	return INV
}

func handleHistoricaltransactionUsedinfos(data []string) []TransactionHistoryUsedinfos {
	var HUI []TransactionHistoryUsedinfos
	for i := 0; i < len(data); i++ {
		str := data[i]
		header, usedinfos := sliceinfohandler(str)
		header_split := strings.Split(header, ",")
		var UsedInfos []Usedinfos
		usedinfos_split := strings.Split(usedinfos, "|")
		if usedinfos_split[0] != "" {
			for i := 0; i < len(usedinfos_split); i++ {
				us := strings.Split(usedinfos_split[i], ",")
				UIfo := Usedinfos{
					us[0],
					us[1],
					us[2],
				}
				UsedInfos = append(UsedInfos, UIfo)
			}
		}
		trui := TransactionHistoryUsedinfos{
			header_split[0],
			header_split[1],
			header_split[2],
			header_split[3],
			header_split[4],
			header_split[5],
			header_split[6],
			UsedInfos,
		}
		// fmt.Println(trsh)
		HUI = append(HUI, trui)
	}
	// fmt.Println(HUI)
	return HUI
}

func handleHistoricaltransactionSettleinfos(data []string) []TransactionHistorySettleinfos {
	var HSI []TransactionHistorySettleinfos
	for i := 0; i < len(data); i++ {
		str := data[i]
		header, settleinfos := sliceinfohandler(str)
		header_split := strings.Split(header, ",")
		var SettleInfos []Settleinfos
		settleinfos_split := strings.Split(settleinfos, "|")
		if settleinfos_split[0] != "" {
			for i := 0; i < len(settleinfos_split); i++ {
				st := strings.Split(settleinfos_split[i], ",")
				SIfo := Settleinfos{
					st[0],
					st[1],
					st[2],
				}
				SettleInfos = append(SettleInfos, SIfo)
			}
		}
		trsi := TransactionHistorySettleinfos{
			header_split[0],
			header_split[1],
			header_split[2],
			header_split[3],
			header_split[4],
			header_split[5],
			header_split[6],
			SettleInfos,
		}
		HSI = append(HSI, trsi)
	}
	// fmt.Println(HSI)
	return HSI
}

func handleHistoricaltransactionOrderinfos(data []string) []TransactionHistoryOrderinfos {
	var HOI []TransactionHistoryOrderinfos
	for i := 0; i < len(data); i++ {
		str := data[i]
		header, orderinfos := sliceinfohandler(str)
		header_split := strings.Split(header, ",")
		var OrderInfos []Orderinfos
		orderinfos_split := strings.Split(orderinfos, "|")
		if orderinfos_split[0] != "" {
			for i := 0; i < len(orderinfos_split); i++ {
				od := strings.Split(orderinfos_split[i], ",")
				OIfo := Orderinfos{
					od[0],
					od[1],
					od[2],
				}
				OrderInfos = append(OrderInfos, OIfo)
			}
		}
		troi := TransactionHistoryOrderinfos{
			header_split[0],
			header_split[1],
			header_split[2],
			header_split[3],
			header_split[4],
			header_split[5],
			header_split[6],
			OrderInfos,
		}
		HOI = append(HOI, troi)
	}
	// fmt.Println(HOI)
	return HOI
}

func handleHistoricaltransactionReceivableinfos(data []string) []TransactionHistoryReceivableinfos {
	var HRI []TransactionHistoryReceivableinfos
	for i := 0; i < len(data); i++ {
		str := data[i]
		header, receivableinfos := sliceinfohandler(str)
		header_split := strings.Split(header, ",")
		var ReceivableInfos []Receivableinfos
		receivableinfos_split := strings.Split(receivableinfos, "|")
		if receivableinfos_split[0] != "" {
			for i := 0; i < len(receivableinfos_split); i++ {
				rc := strings.Split(receivableinfos_split[i], ",")
				RIfo := Receivableinfos{
					rc[0],
					rc[1],
					rc[2],
				}
				ReceivableInfos = append(ReceivableInfos, RIfo)
			}
		}
		trri := TransactionHistoryReceivableinfos{
			header_split[0],
			header_split[1],
			header_split[2],
			header_split[3],
			header_split[4],
			header_split[5],
			header_split[6],
			ReceivableInfos,
		}
		HRI = append(HRI, trri)
	}
	// fmt.Println(HRI)
	return HRI
}

func handleEnterpoolDataPlaninfos(data []string) []EnterpoolDataPlaninfos {
	//如果其他输入中存在[]怎么办？
	//最后返回的结果，目前是结构体的切片
	var EPD []EnterpoolDataPlaninfos
	for i := 0; i < len(data); i++ {
		str := data[i]

		header, planinfos := sliceinfohandler(str)

		header_split := strings.Split(header, ",")
		var PlanInfos []Planinfos
		planinfos_split := strings.Split(planinfos, "|")
		if planinfos_split[0] != "" {
			for i := 0; i < len(planinfos_split); i++ {
				pl := strings.Split(planinfos_split[i], ",")
				PLfo := Planinfos{
					pl[0],
					pl[1],
					pl[2],
				}
				PlanInfos = append(PlanInfos, PLfo)
			}
		}

		epdt := EnterpoolDataPlaninfos{
			header_split[0],
			header_split[1],
			header_split[2],
			header_split[3],
			header_split[4],
			PlanInfos,
		}
		EPD = append(EPD, epdt)
	}
	// fmt.Println(EPD)
	return EPD
}

func handleEnterpoolDataProviderusedinfos(data []string) []EnterpoolDataProviderusedinfos {
	var EPD []EnterpoolDataProviderusedinfos
	for i := 0; i < len(data); i++ {
		str := data[i]
		header, providerusedinfos := sliceinfohandler(str)
		header_split := strings.Split(header, ",")
		var ProviderusedInfos []Providerusedinfos
		providerusedinfos_split := strings.Split(providerusedinfos, "|")
		if providerusedinfos_split[0] != "" {
			for i := 0; i < len(providerusedinfos_split); i++ {
				pr := strings.Split(providerusedinfos_split[i], ",")
				PRfo := Providerusedinfos{
					pr[0],
					pr[1],
					pr[2],
				}
				ProviderusedInfos = append(ProviderusedInfos, PRfo)
			}
		}

		epdt := EnterpoolDataProviderusedinfos{
			header_split[0],
			header_split[1],
			header_split[2],
			header_split[3],
			header_split[4],
			ProviderusedInfos,
		}
		EPD = append(EPD, epdt)
	}
	// fmt.Println(EPD)
	return EPD
}

func handleFinancingIntention(data []string) []FinancingIntention {
	var FCI []FinancingIntention
	for i := 0; i < len(data); i++ {
		str := data[i]
		//fmt.Println(str)
		flag := 0
		header := ""
		for index, val := range str {
			if index+1 >= len(str) {
				break
			}
			if flag == 0 {
				if str[index] == ',' && str[index+1] == '[' {
					flag = 1
				} else {
					header = header + string(val)
				}
			}
		}
		header_split := strings.Split(header, ",")
		fcin := FinancingIntention{
			header_split[0],
			header_split[1],
			header_split[2],
			header_split[3],
			header_split[4],
			header_split[5],
			header_split[6],
			header_split[7],
			header_split[8],
			header_split[9],
			header_split[10],
			header_split[11],
			header_split[12],
		}
		FCI = append(FCI, fcin)
	}
	// fmt.Println(FCI)
	return FCI
}

func handleCollectionAccount(data []string) []CollectionAccount {
	var COLA []CollectionAccount
	for i := 0; i < len(data); i++ {
		str := data[i]
		//fmt.Println(str)
		flag := 0
		header := ""
		for index, val := range str {
			if index+1 >= len(str) {
				break
			}
			if flag == 0 {
				if str[index] == ',' && str[index+1] == '[' {
					flag = 1
				} else {
					header = header + string(val)
				}
			}
		}
		header_split := strings.Split(header, ",")
		cola := CollectionAccount{
			header_split[0],
			header_split[1],
			header_split[2],
			header_split[3],
			header_split[4],
			header_split[5],
			header_split[6],
		}
		COLA = append(COLA, cola)
	}
	// fmt.Println(COLA)
	return COLA
}

func (s *SqlCtr) ConvertoStruct(method string, data []string) string {
	switch method {
	case "HistoricaltransactionUsedinfos":
		result := handleHistoricaltransactionUsedinfos(data)
		ans, err := json.Marshal(result)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(string(ans))
		return string(ans)

	case "HistoricaltransactionSettleinfos":
		result := handleHistoricaltransactionSettleinfos(data)
		ans, err := json.Marshal(result)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(string(ans))
		return string(ans)

	case "HistoricaltransactionOrderinfos":
		result := handleHistoricaltransactionOrderinfos(data)
		ans, err := json.Marshal(result)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(string(ans))
		return string(ans)

	case "HistoricaltransactionReceivableinfos":
		result := handleHistoricaltransactionReceivableinfos(data)
		ans, err := json.Marshal(result)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(string(ans))
		return string(ans)

	case "InvoiceInformation":
		result := handleInvoiceInformation(data)
		ans, err := json.Marshal(result)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(ans))
		return string(ans)

	case "EnterpoolDataPlaninfos":
		result := handleEnterpoolDataPlaninfos(data)
		ans, err := json.Marshal(result)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(string(ans))
		return string(ans)

	case "EnterpoolDataUsedinfos":
		result := handleEnterpoolDataProviderusedinfos(data)
		ans, err := json.Marshal(result)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(string(ans))
		return string(ans)

	case "FinancingIntention":
		result := handleFinancingIntention(data)
		ans, err := json.Marshal(result)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(string(ans))
		return string(ans)

	case "CollectionAccount":
		result := handleCollectionAccount(data)
		ans, err := json.Marshal(result)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(string(ans))
		return string(ans)
	}
	return ""
}

// type InvoiceInformation struct {
// 	Certificateid   string         `json:"certificateId"`
// 	Customerid      string         `json:"customerId"`
// 	Corpname        string         `json:"corpName"`
// 	Certificatetype string         `json:"certificateType"`
// 	Intercustomerid string         `json:"interCustomerId"`
// 	Invoiceinfos    []Invoiceinfos `json:"invoiceInfos"`
// }

// type Invoiceinfos struct {
// 	Invoicenotaxamt string `json:"InvoiceNotaxAmt"`
// 	Invoiceccy      string `json:"InvoiceCcy"`
// 	Sellername      string `json:"SellerName"`
// 	Invoicetype     string `json:"InvoiceType"`
// 	Buyername       string `json:"BuyerName"`
// 	Buyerusccode    string `json:"BuyerUsccode"`
// 	Invoicedate     string `json:"InvoiceDate"`
// 	Sellerusccode   string `json:"SellerUsccode"`
// 	Invoicecode     string `json:"InvoiceCode"`
// 	Invoicenum      string `json:"InvoiceNum"`
// 	Checkcode       string `json:"CheckCode"`
// 	Invoiceamt      string `json:"InvoiceAmt"`
// }

// func handleHistoricaltransactionInformation(data []string) ([]TransactionHistoryUsedinfos, []TransactionHistorySettleinfos, []TransactionHistoryOrderinfos, []TransactionHistoryReceivableinfos) {
// 	var HUI []TransactionHistoryUsedinfos
// 	var HSI []TransactionHistorySettleinfos
// 	var HOI []TransactionHistoryOrderinfos
// 	var HRI []TransactionHistoryReceivableinfos
// 	for i := 0; i < len(data); i++ {
// 		str := data[i]
// 		flag := 0
// 		header := ""
// 		usedinfos := ""
// 		settleinfos := ""
// 		orderinfos := ""
// 		receivableinfos := ""
// 		for index, val := range str {
// 			if index+1 >= len(str) {
// 				break
// 			}
// 			if flag == 0 {
// 				if str[index] == ',' && str[index+1] == '[' {
// 					flag = 1
// 				} else {
// 					header = header + string(val)
// 				}
// 			} else if flag == 1 {
// 				if str[index] == '[' && str[index+1] == ',' {
// 					flag = 2
// 				} else if str[index] == ']' {
// 					flag = 2
// 				} else if str[index] != '[' && str[index] != ']' {
// 					usedinfos = usedinfos + string(val)
// 				}
// 			} else if flag == 2 {
// 				if str[index] == '[' && str[index+1] == ',' {
// 					flag = 3
// 				} else if str[index] == ']' {
// 					flag = 3
// 				} else if str[index] != '[' && str[index] != ']' {
// 					if len(settleinfos) == 0 && str[index] == ',' {
// 						continue
// 					} else {
// 						settleinfos = settleinfos + string(val)
// 					}
// 				}
// 			} else if flag == 3 {
// 				if str[index] == '[' && str[index+1] == ',' {
// 					flag = 4
// 				} else if str[index] == ']' {
// 					flag = 4
// 				} else if str[index] != '[' && str[index] != ']' {
// 					if len(orderinfos) == 0 && str[index] == ',' {
// 						continue
// 					} else {
// 						orderinfos = orderinfos + string(val)
// 					}
// 				}
// 			} else if flag == 4 {
// 				if str[index] == '[' && str[index+1] == ',' {
// 					flag = 5
// 				} else if str[index] == ']' {
// 					flag = 5
// 				} else if str[index] != '[' && str[index] != ']' {
// 					if len(receivableinfos) == 0 && str[index] == ',' {
// 						continue
// 					} else {
// 						receivableinfos = receivableinfos + string(val)
// 					}
// 				}
// 			}
// 		}
// 		header_split := strings.Split(header, ",")
// 		var UsedInfos []Usedinfos
// 		var SettleInfos []Settleinfos
// 		var OrderInfos []Orderinfos
// 		var ReceivableInfos []Receivableinfos

// 		usedinfos_split := strings.Split(usedinfos, "|")
// 		if usedinfos_split[0] != "" {
// 			for i := 0; i < len(usedinfos_split); i++ {
// 				us := strings.Split(usedinfos_split[i], ",")
// 				UIfo := Usedinfos{
// 					us[0],
// 					us[1],
// 					us[2],
// 				}
// 				UsedInfos = append(UsedInfos, UIfo)
// 			}
// 		}

// 		settleinfos_split := strings.Split(settleinfos, "|")
// 		if settleinfos_split[0] != "" {
// 			for i := 0; i < len(settleinfos_split); i++ {
// 				st := strings.Split(settleinfos_split[i], ",")
// 				SIfo := Settleinfos{
// 					st[0],
// 					st[1],
// 					st[2],
// 				}
// 				SettleInfos = append(SettleInfos, SIfo)
// 			}
// 		}

// 		orderinfos_split := strings.Split(orderinfos, "|")
// 		if orderinfos_split[0] != "" {
// 			for i := 0; i < len(orderinfos_split); i++ {
// 				od := strings.Split(orderinfos_split[i], ",")
// 				OIfo := Orderinfos{
// 					od[0],
// 					od[1],
// 					od[2],
// 				}
// 				OrderInfos = append(OrderInfos, OIfo)
// 			}
// 		}

// 		receivableinfos_split := strings.Split(receivableinfos, "|")
// 		if receivableinfos_split[0] != "" {
// 			for i := 0; i < len(receivableinfos_split); i++ {
// 				rc := strings.Split(receivableinfos_split[i], ",")
// 				RIfo := Receivableinfos{
// 					rc[0],
// 					rc[1],
// 					rc[2],
// 				}
// 				ReceivableInfos = append(ReceivableInfos, RIfo)
// 			}
// 		}

// 		trui := TransactionHistoryUsedinfos{
// 			header_split[0],
// 			header_split[1],
// 			header_split[2],
// 			header_split[3],
// 			header_split[4],
// 			header_split[5],
// 			header_split[6],
// 			UsedInfos,
// 		}
// 		// fmt.Println(trsh)
// 		HUI = append(HUI, trui)
// 		trsi := TransactionHistorySettleinfos{
// 			header_split[0],
// 			header_split[1],
// 			header_split[2],
// 			header_split[3],
// 			header_split[4],
// 			header_split[5],
// 			header_split[6],
// 			SettleInfos,
// 		}
// 		HSI = append(HSI, trsi)
// 		troi := TransactionHistoryOrderinfos{
// 			header_split[0],
// 			header_split[1],
// 			header_split[2],
// 			header_split[3],
// 			header_split[4],
// 			header_split[5],
// 			header_split[6],
// 			OrderInfos,
// 		}
// 		HOI = append(HOI, troi)
// 		trri := TransactionHistoryReceivableinfos{
// 			header_split[0],
// 			header_split[1],
// 			header_split[2],
// 			header_split[3],
// 			header_split[4],
// 			header_split[5],
// 			header_split[6],
// 			ReceivableInfos,
// 		}
// 		HRI = append(HRI, trri)
// 	}
// 	return HUI, HSI, HOI, HRI
// }

// func handleEnterpoolData(data []string) []EnterpoolData {
// 	//如果其他输入中存在[]怎么办？
// 	//最后返回的结果，目前是结构体的切片
// 	var EPD []EnterpoolData
// 	for i := 0; i < len(data); i++ {
// 		str := data[i]
// 		flag := 0
// 		header := ""
// 		planinfos := ""
// 		providerusedinfos := ""
// 		for index, val := range str {
// 			if index+1 >= len(str) {
// 				break
// 			}
// 			if flag == 0 {
// 				if str[index] == ',' && str[index+1] == '[' {
// 					flag = 1
// 				} else {
// 					header = header + string(val)
// 				}
// 			} else if flag == 1 {
// 				if str[index] == '[' && str[index+1] == ',' {
// 					flag = 2
// 				} else if str[index] == ']' {
// 					flag = 2
// 				} else if str[index] != '[' && str[index] != ']' {
// 					planinfos = planinfos + string(val)
// 				}
// 			} else if flag == 2 {
// 				if str[index] == '[' && str[index+1] == ',' {
// 					flag = 3
// 				} else if str[index] == ']' {
// 					flag = 3
// 				} else if str[index] != '[' && str[index] != ']' {
// 					if len(providerusedinfos) == 0 && str[index] == ',' {
// 						continue
// 					} else {
// 						providerusedinfos = providerusedinfos + string(val)
// 					}
// 				}
// 			}
// 		}
// 		header_split := strings.Split(header, ",")
// 		var PlanInfos []Planinfos
// 		planinfos_split := strings.Split(planinfos, "|")
// 		if planinfos_split[0] != "" {
// 			for i := 0; i < len(planinfos_split); i++ {
// 				pl := strings.Split(planinfos_split[i], ",")
// 				PLfo := Planinfos{
// 					pl[0],
// 					pl[1],
// 					pl[2],
// 				}
// 				PlanInfos = append(PlanInfos, PLfo)
// 			}
// 		}
// 		var ProviderusedInfos []Providerusedinfos
// 		providerusedinfos_split := strings.Split(providerusedinfos, "|")
// 		if providerusedinfos_split[0] != "" {
// 			for i := 0; i < len(providerusedinfos_split); i++ {
// 				pr := strings.Split(providerusedinfos_split[i], ",")
// 				PRfo := Providerusedinfos{
// 					pr[0],
// 					pr[1],
// 					pr[2],
// 				}
// 				ProviderusedInfos = append(ProviderusedInfos, PRfo)
// 			}
// 		}

// 		epdt := EnterpoolData{
// 			header_split[0],
// 			header_split[1],
// 			header_split[2],
// 			header_split[3],
// 			header_split[4],
// 			PlanInfos,
// 			ProviderusedInfos,
// 		}
// 		EPD = append(EPD, epdt)
// 	}
// 	fmt.Println(EPD)
// 	return EPD
// }
