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

func (r *Repository) FindPaginated(page, limit int) ([]loan.Loan, int64) {
	var loans []loan.Loan
	var total int64
	tableName := loan.Loan{}.TableName()

	tx := r.db.Table(tableName)
	tx.Count(&total)
	tx.Limit(limit).Offset((page - 1) * limit).Find(&loans)

	if tx.Error != nil {
		return nil, total
	}
	return loans, total
}

func (r *Repository) Create(l *loan.Loan) *loan.Loan {
	r.db.Create(&l)
	return l
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
