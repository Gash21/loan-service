package investor

import (
	"github.com/Gash21/amartha-test/internal/domain/investor"
	"github.com/Gash21/amartha-test/internal/shared/database"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *database.Database) investor.InvestorRepository {
	return &Repository{db: db.Gorm}
}

func (r *Repository) Create(m *investor.Investor) *investor.Investor {
	r.db.Create(&m)
	return m
}

func (r *Repository) FindByName(name string) (investor.Investor, error) {
	var m investor.Investor
	err := r.db.Where("name = ?", name).First(&m)
	return m, err.Error
}

func (r *Repository) FindByID(id int64) (investor.Investor, error) {
	var m investor.Investor
	err := r.db.Where("id = ?", id).First(&m)
	return m, err.Error
}
