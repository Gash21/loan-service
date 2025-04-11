package document

import (
	"github.com/Gash21/amartha-test/internal/domain/document"
	"github.com/Gash21/amartha-test/internal/shared/database"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *database.Database) document.DocumentRepository {
	return &Repository{db: db.Gorm}
}

func (r *Repository) CreateDocument(document *document.Document) error {
	return nil
}
