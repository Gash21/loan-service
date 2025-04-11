package loan_investor

import "time"

type LoanInvestor struct {
	LoanID            *int64  `json:"loan_id" gorm:"column:loan_id;type:bigint;primarykey"`
	InvestorID        *int64  `json:"investor_id" gorm:"column:investor_id;type:bigint;primarykey"`
	Amount            float64 `json:"amount" gorm:"column:amount"`
	Rate              float64 `json:"rate" gorm:"column:rate"`
	ROI               float64 `json:"roi" gorm:"column:roi"`
	AgreementLetterID *int64  `json:"agreement_letter_id" gorm:"column:agreement_letter_id;type:bigint"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;->;<-:create;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (LoanInvestor) TableName() string {
	return "loan_investors"
}
