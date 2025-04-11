package loan

type LoanRepository interface {
	ProposeLoan() Loan
	ApprovedLoan() Loan
	InvestedLoan() Loan
	DisbursedLoan() Loan
}
