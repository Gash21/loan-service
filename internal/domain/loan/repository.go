package loan

type LoanRepository interface {
	FindPaginated(page, limit int) ([]Loan, int64)
	FindByID(id *int64) (*Loan, error)
	FindByIDAndStatus(id *int64, status string) (*Loan, error)
	Create(*Loan) (*Loan, error)
	Update(*Loan) (*Loan, error)
}
