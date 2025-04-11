package borrower

type BorrowerRepository interface {
	CreateBorrower(Borrower) Borrower
	GetBorrowerByName(string) (Borrower, error)
	GetBorrowerById(string) (Borrower, error)
}
