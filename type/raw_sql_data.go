package types

type RawSQLData struct {
	SQLId  string
	Num    string
	Status string
	ID     string
	Data   string
	Key    string
	Hash   string
}

// 发票信息推送接口
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
type TransactionHistory struct {
	Customergrade   string            `json:"customerGrade"`
	Certificatetype string            `json:"certificateType"`
	Intercustomerid string            `json:"interCustomerId"`
	Corpname        string            `json:"corpName"`
	Financeid       string            `json:"financeId"`
	Certificateid   string            `json:"certificateId"`
	Customerid      string            `json:"customerId"`
	Usedinfos       []Usedinfos       `json:"usedInfos"`
	Settleinfos     []Settleinfos     `json:"settleInfos"`
	Orderinfos      []Orderinfos      `json:"orderInfos"`
	Receivableinfos []Receivableinfos `json:"receivableInfos"`
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

// 推送入池数据接口
type EnterpoolData struct {
	Datetimepoint     string              `json:"dateTimePoint"`
	Ccy               string              `json:"ccy"`
	Customerid        string              `json:"customerId"`
	Intercustomerid   string              `json:"interCustomerId"`
	Receivablebalance string              `json:"receivableBalance"`
	Planinfos         []Planinfos         `json:"planInfos"`
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

// 提交融资意向接口
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

// 推送回款账户接口
type CollectionAccount struct {
	Backaccount     string `json:"BackAccount"`
	Certificateid   string `json:"CertificateId"`
	Customerid      string `json:"CustomerId"`
	Corpname        string `json:"CorpName"`
	Lockremark      string `json:"LockRemark"`
	Certificatetype string `json:"CertificateType"`
	Intercustomerid string `json:"InterCustomerId"`
}
type RawInvoiceInformation struct {
	Certificateid   string         `json:"certificateId"`
	Customerid      string         `json:"customerId"`
	Corpname        string         `json:"corpName"`
	Certificatetype string         `json:"certificateType"`
	Intercustomerid string         `json:"interCustomerId"`
	Invoiceinfos    []Invoiceinfos `json:"invoiceInfos"`
}

type Invoiceinfos struct {
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
