package loan_investor

import "time"

type LoanInvestor struct {
	ID              *int64  `json:"id" gorm:"column:id;primaryKey;autoIncrement:true;not null"`
	LoanID          *int64  `json:"loan_id" gorm:"column:loan_id;type:bigint"`
	InvestorID      *int64  `json:"investor_id" gorm:"column:investor_id;type:bigint"`
	Amount          float64 `json:"amount" gorm:"column:amount"`
	Rate            float64 `json:"rate" gorm:"column:rate"`
	ROI             float64 `json:"roi" gorm:"column:roi"`
	AgreementLetter *string `json:"agreement_letter" gorm:"column:agreement_letter;type:text"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;->;<-:create;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (LoanInvestor) TableName() string {
	return "loan_investors"
}
