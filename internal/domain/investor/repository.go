package investor

type InvestorRepository interface {
	Create(*Investor) *Investor
	FindByID(*int64) (*Investor, error)
	FindByName(string) (Investor, error)
}
