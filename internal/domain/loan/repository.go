package loan

type LoanRepository interface {
	FindPaginated(page, limit int) ([]Loan, int64)
	Create(Loan) Loan
	ApproveLoan() Loan
	InvesteLoan() Loan
	DisburseLoan() Loan
}
