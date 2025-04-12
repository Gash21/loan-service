package loan_investor

type LoanInvestorRepository interface {
	Create(*LoanInvestor) *LoanInvestor
}
