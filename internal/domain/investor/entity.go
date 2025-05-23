package investor

import (
	"time"

	"github.com/Gash21/amartha-test/internal/domain/loan_investor"
)

type Investor struct {
	ID        *int64    `json:"id" gorm:"column:id;primaryKey;autoIncrement:true;not null"`
	Name      string    `json:"name" gorm:"column:name;type:varchar(255)"`
	Email     string    `json:"email" gorm:"column:email;unique;type:varchar(255)"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;->;<-:create;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	LoanInvestors []*loan_investor.LoanInvestor `gorm:"foreignKey:InvestorID;references:ID" json:"loan_investors"`
}

func (Investor) TableName() string {
	return "investors"
}
