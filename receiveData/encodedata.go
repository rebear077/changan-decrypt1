package receive

// 发票信息推送接口
type EncodeInvoiceInformation struct {
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
type EncodeTransactionHistory struct {
	Customergrade       string
	Certificatetype     string
	Intercustomerid     string
	Corpname            string
	Financeid           string
	Certificateid       string
	Customerid          string
	UsedinfosList       string
	SettleinfosList     string
	OrderinfosList      string
	ReceivableinfosList string
}

// 推送入池数据接口
type EncodeEnterpoolData struct {
	Datetimepoint         string
	Ccy                   string
	Customerid            string
	Intercustomerid       string
	Receivablebalance     string
	PlaninfosList         string
	ProviderusedinfosList string
}

// 提交融资意向接口
type EncodeFinancingIntention struct {
	Custcdlinkposition string
	Custcdlinkname     string
	Certificateid      string
	Corpname           string
	Remark             string
	Bankcontact        string
	Banklinkname       string
	Custcdcontact      string
	Customerid         string
	Financeid          string
	Cooperationyears   string
	Certificatetype    string
	Intercustomerid    string
}

// 推送回款账户接口
type EncodeCollectionAccount struct {
	Backaccount     string
	Certificateid   string
	Customerid      string
	Corpname        string
	Lockremark      string
	Certificatetype string
	Intercustomerid string
}
