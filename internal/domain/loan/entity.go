package loan

import (
	"time"

	"github.com/Gash21/amartha-test/internal/domain/borrower"
	"github.com/Gash21/amartha-test/internal/domain/employee"
	"github.com/Gash21/amartha-test/internal/domain/loan_investor"
)

type Loan struct {
	ID               *int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement:true;not null"`
	BorrowerID       *int64     `json:"borrower_id" gorm:"column:borrower_id"`
	TotalInvested    float64    `json:"total_invested" gorm:"column:total_invested"`
	PrincipalAmount  float64    `json:"principal_amount" gorm:"column:principal_amount"`
	Status           string     `json:"status" gorm:"column:status"`
	Rate             float64    `json:"rate" gorm:"column:rate"`
	TotalAmount      float64    `json:"total_amount" gorm:"column:total_amount"`
	ApprovedAt       *time.Time `json:"approved_at" gorm:"column:approved_at"`
	ApprovedBy       *int64     `json:"approved_by" gorm:"column:approved_by"`
	ApprovalDocument *string    `json:"approval_document" gorm:"column:approval_document"`
	BorrowerLetter   *string    `json:"borrower_letter" gorm:"column:borrower_letter;type:text"`
	CreatedAt        time.Time  `gorm:"column:created_at;autoCreateTime;->;<-:create;" json:"created_at"`
	UpdatedAt        time.Time  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	Approver      *employee.Employee            `gorm:"foreignKey:ApprovedBy;references:ID" json:"approver,omitempty"`
	Borrower      *borrower.Borrower            `gorm:"foreignKey:BorrowerID;references:ID" json:"borrower,omitempty"`
	LoanInvestors *[]loan_investor.LoanInvestor `gorm:"foreignKey:LoanID;references:ID" json:"investors,omitempty"`
}

func (Loan) TableName() string {
	return "loans"
}
