package loan

import (
	"github.com/Gash21/amartha-test/internal/domain/loan"
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

func (r *Repository) ApproveLoan() loan.Loan {
	return loan.Loan{}
}

func (r *Repository) InvesteLoan() loan.Loan {
	return loan.Loan{}
}

func (r *Repository) DisburseLoan() loan.Loan {
	return loan.Loan{}
}
