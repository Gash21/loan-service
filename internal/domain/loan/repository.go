package loan

type LoanRepository interface {
	ProposeLoan() Loan
	ApproveLoan() Loan
	InvesteLoan() Loan
	DisburseLoan() Loan
}
