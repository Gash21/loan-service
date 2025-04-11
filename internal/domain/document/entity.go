package document

import (
	"time"
)

type Document struct {
	ID        *int64    `json:"id" gorm:"column:id;primaryKey;autoIncrement:true;not null"`
	Path      string    `json:"path" gorm:"column:path;type:varchar(255)"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;->;<-:create;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (Document) TableName() string {
	return "documents"
}
