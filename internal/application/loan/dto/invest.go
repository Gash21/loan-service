package dto

import (
	"fmt"

	"github.com/Gash21/amartha-test/internal/shared/config"
)

type InvestRequest struct {
	ID         *int64  `params:"id"`
	InvestorID *int64  `json:"investor_id" validate:"required"`
	Amount     float64 `json:"amount" validate:"required"`
}

type InvestResponse struct {
	AgreementLetter string `json:"agreement_letter"`
}

func (r *InvestResponse) ToResponse(cfg config.GlobalConfig, path string) *InvestResponse {
	r.AgreementLetter = fmt.Sprintf("%s/%s", cfg.PublicURL, path)
	return r
}
