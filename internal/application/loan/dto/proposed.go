package dto

type ProposedRequest struct {
	BorrowerID      int64   `json:"borrower_id"`
	BorrowerName    string  `json:"borrower_name" validate:"required"`
	PrincipalAmount float64 `json:"principal_amount" validate:"required"`
	Rate            float64 `json:"rate" validate:"required"`
}

type ProposeResponse struct {
	LoanID int64 `json:"loan_id"`
}
