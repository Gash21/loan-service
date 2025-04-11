package dto

type ProposedRequest struct {
	BorrowerID      int64   `json:"borrower_id" validate:"required"`
	PrincipalAmount float64 `json:"principal_amount" validate:"required"`
	Rate            float64 `json:"rate" validate:"required"`
}

type ProposeResponse struct {
	LoanID int64 `json:"loan_id"`
}
