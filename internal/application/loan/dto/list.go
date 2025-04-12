package dto

import "github.com/Gash21/amartha-test/internal/domain/loan"

type ListRequest struct {
	Page   int     `query:"page"`
	Limit  int     `query:"limit"`
	Status *string `query:"status"`
}

type ListResponse struct {
	Data []DetailResponse
}

func (r *ListResponse) FromModel(loans []loan.Loan) *ListResponse {
	for _, loan := range loans {
		r.Data = append(r.Data, DetailResponse{
			ID:              loan.ID,
			Status:          loan.Status,
			PrincipalAmount: loan.PrincipalAmount,
			Rate:            loan.Rate,
			ROI:             loan.TotalAmount - loan.PrincipalAmount,
		})
	}
	return r
}
