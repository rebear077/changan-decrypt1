package sql

import "strings"

type InvoiceInfo struct {
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

func handleInvoiceInfo(data []string) []InvoiceInformation {
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
