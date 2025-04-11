package borrower

import (
	"time"
)

type Borrower struct {
	ID        *int64    `json:"id" gorm:"column:id;primaryKey;autoIncrement:true;not null"`
	Name      string    `json:"name" gorm:"column:name;type:varchar(255)"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;->;<-:create;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (Borrower) TableName() string {
	return "borrowers"
}
