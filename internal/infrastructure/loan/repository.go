package loan

import (
	"github.com/Gash21/amartha-test/internal/application/loan"
	"github.com/Gash21/amartha-test/internal/shared/database"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *database.Database) loan.LoanRepository {
	return &Repository{db: db.Gorm}
}

func (r *Repository) ProposeLoan() loan.Loan {
	return loan.Loan{}
}

func (r *Repository) ApprovedLoan() loan.Loan {
	return loan.Loan{}
}

func (r *Repository) InvestedLoan() loan.Loan {
	return loan.Loan{}
}

func (r *Repository) DisbursedLoan() loan.Loan {
	return loan.Loan{}
}
