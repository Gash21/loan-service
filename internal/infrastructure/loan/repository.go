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

func (r *Repository) FindPaginated(page, limit int, status *string) ([]loan.Loan, int64) {
	var loans []loan.Loan
	var total int64

	tx := r.db

	if status != nil {
		tx = tx.Where("status = ?", *status)
	}

	tx.Count(&total)
	tx.Limit(limit).Offset((page - 1) * limit)
	tx.Preload("LoanInvestors").Find(&loans)

	if tx.Error != nil {
		return nil, total
	}
	return loans, total
}

func (r *Repository) FindByID(id *int64) (*loan.Loan, error) {
	var data loan.Loan
	tableName := loan.Loan{}.TableName()

	err := r.db.Table(tableName).Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *Repository) FindByIDWithLoanInvestor(id *int64) (*loan.Loan, error) {
	var data loan.Loan
	tableName := loan.Loan{}.TableName()

	err := r.db.Table(tableName).Where("id = ?", id).Preload("LoanInvestors").First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *Repository) FindByIDAndStatus(id *int64, status string) (*loan.Loan, error) {
	var data loan.Loan
	tableName := loan.Loan{}.TableName()

	err := r.db.Table(tableName).Where("id = ?", id).Where("status = ?", status).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *Repository) FindInvestableLoanByID(id *int64) (*loan.Loan, error) {
	var data loan.Loan
	tableName := loan.Loan{}.TableName()

	err := r.db.Table(tableName).Where("id = ?", id).Preload("LoanInvestors").
		Where("status = ?", "approved").First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *Repository) Create(l *loan.Loan) (*loan.Loan, error) {
	tx := r.db.Create(&l)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return l, nil
}

func (r *Repository) Update(l *loan.Loan) (*loan.Loan, error) {
	tx := r.db.Save(&l)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return l, nil
}
