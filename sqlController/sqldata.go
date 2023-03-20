package sql

import (
	"strings"

	types "github.com/FISCO-BCOS/go-sdk/type"
)

func handleInvoiceInfo(data []string) []*types.InvoiceInformation {
	//如果其他输入中存在[]怎么办？
	//最后返回的结果，目前是结构体的切片
	var INV []*types.InvoiceInformation
	for i := 0; i < len(data); i++ {
		str := data[i]
		//fmt.Println(str)
		str_split := strings.Split(str, ",")

		ICfo := types.InvoiceInformation{
			Certificateid:   str_split[0],
			Customerid:      str_split[1],
			Corpname:        str_split[2],
			Certificatetype: str_split[3],
			Intercustomerid: str_split[4],
			Invoicenotaxamt: str_split[5],
			Invoiceccy:      str_split[6],
			Sellername:      str_split[7],
			Invoicetype:     str_split[8],
			Buyername:       str_split[9],
			Buyerusccode:    str_split[10],
			Invoicedate:     str_split[11],
			Sellerusccode:   str_split[12],
			Invoicecode:     str_split[13],
			Invoicenum:      str_split[14],
			Checkcode:       str_split[15],
			Invoiceamt:      str_split[16],
		}
		INV = append(INV, &ICfo)
	}
	// fmt.Println(INV)
	return INV
}
