package borrower

type BorrowerRepository interface {
	Create(*Borrower) *Borrower
	FindByName(string) (Borrower, error)
	FindByID(int64) (Borrower, error)
}
