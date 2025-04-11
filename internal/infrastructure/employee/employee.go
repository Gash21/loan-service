package employee

import (
	"github.com/Gash21/amartha-test/internal/domain/employee"
	"github.com/Gash21/amartha-test/internal/shared/database"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *database.Database) employee.EmployeeRepository {
	return &Repository{db: db.Gorm}
}

func (r *Repository) Create(m *employee.Employee) *employee.Employee {
	r.db.Create(&m)
	return m
}

func (r *Repository) FindByName(name string) (employee.Employee, error) {
	var m employee.Employee
	err := r.db.Where("name = ?", name).First(&m)
	return m, err.Error
}

func (r *Repository) FindByID(id int64) (employee.Employee, error) {
	var m employee.Employee
	err := r.db.Where("id = ?", id).First(&m)
	return m, err.Error
}
