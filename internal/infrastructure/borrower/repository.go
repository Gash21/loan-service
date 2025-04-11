package borrower

import (
	"github.com/Gash21/amartha-test/internal/domain/borrower"
	"github.com/Gash21/amartha-test/internal/shared/database"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *database.Database) borrower.BorrowerRepository {
	return &Repository{db: db.Gorm}
}

func (r *Repository) CreateBorrower(m borrower.Borrower) borrower.Borrower {
	r.db.Create(&m)
	return m
}

func (r *Repository) GetBorrowerByName(name string) (borrower.Borrower, error) {
	var m borrower.Borrower
	err := r.db.Where("name = ?", name).First(&m)
	return m, err.Error
}

func (r *Repository) GetBorrowerById(id string) (borrower.Borrower, error) {
	var m borrower.Borrower
	err := r.db.Where("id = ?", id).First(&m)
	return m, err.Error
}
