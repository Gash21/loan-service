package dto

import (
	"github.com/Gash21/amartha-test/internal/domain/loan"
	"github.com/Gash21/amartha-test/internal/shared/config"
)

type ListRequest struct {
	Page   int     `query:"page"`
	Limit  int     `query:"limit"`
	Status *string `query:"status"`
}

type ListResponse struct {
	Data []DetailResponse
}

func (r *ListResponse) FromModel(cfg *config.GlobalConfig, loans *[]loan.Loan) *ListResponse {
	for _, loan := range *loans {
		detail := &DetailResponse{}
		detail = detail.FromModel(cfg, &loan)
		r.Data = append(r.Data, *detail)
	}
	return r
}
