package server

// import (
// 	"strings"

// 	sql "github.com/FISCO-BCOS/go-sdk/sqlController"
// )

// func decodeToProviderissueFinancingData(raw []byte) *sql.ProviderissueFinancingData {
// 	targetStruct := new(sql.ProviderissueFinancingData)
// 	res := strings.Split(string(raw), ",")
// 	targetStruct.Job = res[0]
// 	targetStruct.Contact = res[1]
// 	targetStruct.ProviderID = res[2]
// 	targetStruct.ProviderName = res[3]
// 	targetStruct.Remarks = res[4]
// 	targetStruct.BankContactMethod = res[5]
// 	targetStruct.BankContact = res[6]
// 	targetStruct.ContactMethod = res[7]
// 	targetStruct.ProviderNumber = res[8]
// 	targetStruct.FinancingNumber = res[9]
// 	targetStruct.CooperationYears = res[10]
// 	targetStruct.ProviderIDCardType = res[11]
// 	targetStruct.CoreCompanyID = res[12]
// 	return targetStruct
// }
// func decodeToReceiptInformationData(raw []byte) *sql.ReceiptInformationData {
// 	targetStruct := new(sql.ReceiptInformationData)
// 	res := strings.Split(string(raw), ",")
// 	targetStruct.ProviderID = res[0]
// 	targetStruct.ProviderNumber = res[1]
// 	targetStruct.ProviderName = res[2]
// 	targetStruct.ProviderIDCardType = res[3]
// 	targetStruct.CoreCompanyID = res[4]
// 	targetStruct.RecepitWithoutTax = res[5]
// 	targetStruct.ReceiptMoneyType = res[6]
// 	targetStruct.SellerName = res[7]
// 	targetStruct.ReceiptType = res[8]
// 	targetStruct.BuyerName = res[9]
// 	targetStruct.BuyerTIN = res[10]
// 	targetStruct.ReceiptDate = res[11]
// 	targetStruct.SellerTIN = res[12]
// 	targetStruct.ReceiptCode = res[13]
// 	targetStruct.ReceiptNUmber = res[14]
// 	targetStruct.ReceiptMoney = res[15]
// 	return targetStruct
// }
// func decodeToTransactionHistory(raw []byte) *sql.TransactionHistory {
// 	targetStruct := new(sql.TransactionHistory)
// 	res := strings.Split(string(raw), ",")
// 	targetStruct.ProviderRank = res[0]
// 	targetStruct.ProviderIDCardType = res[1]
// 	targetStruct.CoreCompanyID = res[2]
// 	targetStruct.ProviderName = res[3]
// 	targetStruct.FinancingNumber = res[4]
// 	targetStruct.ProviderID = res[5]
// 	targetStruct.ProviderNumber = res[6]
// 	targetStruct.InTransactionDate = res[7]
// 	targetStruct.InMoney = res[8]
// 	targetStruct.InMoneyType = res[9]
// 	targetStruct.CloseTransactionDate = res[10]
// 	targetStruct.CloseMoney = res[11]
// 	targetStruct.CloseMoneyType = res[12]
// 	targetStruct.OrderTransactionDate = res[13]
// 	targetStruct.OrderMoney = res[14]
// 	targetStruct.OrderMoneyType = res[15]
// 	targetStruct.AccountTransactionDate = res[16]
// 	targetStruct.AccountMoney = res[17]
// 	targetStruct.AccountMoneyType = res[18]
// 	return targetStruct
// }
// func decodeToWeektransactionRecord(raw []byte) *sql.WeektransactionRecord {
// 	targetStruct := new(sql.WeektransactionRecord)
// 	res := strings.Split(string(raw), ",")
// 	targetStruct.DataTime = res[0]
// 	targetStruct.MoneyType = res[1]
// 	targetStruct.ProviderNumber = res[2]
// 	targetStruct.CoreCompanyID = res[3]
// 	targetStruct.AccountRemaining = res[4]
// 	targetStruct.PlanTransactionDate = res[5]
// 	targetStruct.PlanMoney = res[6]
// 	targetStruct.PlanMoneyType = res[7]
// 	targetStruct.InTransactionDate = res[8]
// 	targetStruct.InMoney = res[9]
// 	targetStruct.InMoneyType = res[10]
// 	return targetStruct
// }
// func decodeToFinancingResult(raw []byte) *sql.FinancingResult {
// 	targetStruct := new(sql.FinancingResult)
// 	res := strings.Split(string(raw), ",")
// 	targetStruct.FinancingNumber = res[0]
// 	targetStruct.FinancingResult = res[1]
// 	targetStruct.Quota = res[2]
// 	targetStruct.StartDate = res[3]
// 	targetStruct.EndDate = res[4]
// 	targetStruct.ToAccount = res[5]
// 	targetStruct.Remarks = res[6]
// 	return targetStruct
// }
// func decodeToLoanResult(raw []byte) *sql.LoanResult {
// 	targetStruct := new(sql.LoanResult)
// 	res := strings.Split(string(raw), ",")
// 	targetStruct.ProviderNumber = res[0]
// 	targetStruct.ProviderID = res[1]
// 	targetStruct.ProviderName = res[2]
// 	targetStruct.BusinessNumber = res[3]
// 	targetStruct.BusinessMoney = res[4]
// 	targetStruct.BusinessType = res[5]
// 	targetStruct.StartDate = res[6]
// 	targetStruct.EndDate = res[7]
// 	targetStruct.ReceiptNumber = res[8]
// 	return targetStruct
// }
// func decodeToRepayResult(raw []byte) *sql.RepayResult {
// 	targetStruct := new(sql.RepayResult)
// 	res := strings.Split(string(raw), ",")
// 	targetStruct.ProviderNumber = res[0]
// 	targetStruct.ProviderID = res[1]
// 	targetStruct.ProviderName = res[2]
// 	targetStruct.BusinessNumber = res[3]
// 	targetStruct.BusinessMoney = res[4]
// 	targetStruct.Date = res[5]
// 	targetStruct.LoanNumber = res[6]
// 	targetStruct.BusinessRemaining = res[7]
// 	targetStruct.IsEnd = res[8]
// 	return targetStruct
// }
