package dto

type InvestRequest struct {
	ID         *int64  `params:"id"`
	InvestorID *int64  `json:"investor_id" validate:"required"`
	Amount     float64 `json:"amount" validate:"required"`
}
