package dto

import "github.com/Gash21/amartha-test/internal/domain/loan"

type DetailRequest struct {
	ID int64 `json:"id" validate:"required"`
}

type DetailResponse struct {
	ID              *int64  `json:"id"`
	Status          string  `json:"status"`
	PrincipalAmount float64 `json:"principal_amount"`
	Rate            float64 `json:"rate"`
	ROI             float64 `json:"roi"`
}

func (r *DetailResponse) FromLoan(l *loan.Loan) *DetailResponse {
	return &DetailResponse{
		ID:              l.ID,
		Status:          l.Status,
		PrincipalAmount: l.PrincipalAmount,
		Rate:            l.Rate,
		ROI:             l.TotalAmount - l.PrincipalAmount,
	}
}
