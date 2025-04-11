package document

import "github.com/Gash21/amartha-test/internal/domain/loan"

type Document struct {
	ID     int64  `json:"id" gorm:"column:id;type:bigint;primarykey"`
	Path   string `json:"path" gorm:"column:path;type:varchar(255)"`
	LoanID int64  `json:"loan_id" gorm:"column:loan_id;type:bigint"`

	Loan *loan.Loan `gorm:"foreignKey:LoanID;references:ID" json:"loan"`
}

func (Document) TableName() string {
	return "documents"
}
