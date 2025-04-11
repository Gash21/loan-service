package loan

import (
	"time"

	"github.com/Gash21/amartha-test/internal/domain/borrower"
	"github.com/Gash21/amartha-test/internal/domain/investor"
)

type Loan struct {
	ID              int64     `json:"id" gorm:"column:id,primarykey"`
	BorrowerID      int64     `json:"borrower_id" gorm:"column:borrower_id"`
	PrincipalAmount float64   `json:"principal_amount" gorm:"column:principal_amount"`
	Status          string    `json:"status" gorm:"column:status"`
	Rate            float64   `json:"rate" gorm:"column:rate"`
	ROI             float64   `json:"roi" gorm:"column:roi"`
	CreatedAt       time.Time `gorm:"column:created_at;autoCreateTime;->;<-:create;" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	Borrower  borrower.Borrower   `gorm:"foreignKey:BorrowerID;references:ID" json:"borrower"`
	Investors []investor.Investor `gorm:"many2many:loan_investors;" json:"investors"`
}

func (Loan) TableName() string {
	return "loans"
}
