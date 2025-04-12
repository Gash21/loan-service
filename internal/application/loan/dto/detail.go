package dto

import (
	"fmt"

	"github.com/Gash21/amartha-test/internal/domain/loan"
	"github.com/Gash21/amartha-test/internal/shared/config"
	"github.com/Gash21/amartha-test/internal/shared/helper"
)

type DetailRequest struct {
	ID int64 `json:"id" validate:"required"`
}

type DocumentDetail struct {
	ApprovalDocument         *string   `json:"approval_document"`
	DisbursalDocument        *string   `json:"disbursal_document"`
	InvestorAgreementLetters []*string `json:"investor_agreement_letters"`
}

type DetailResponse struct {
	ID              *int64          `json:"id"`
	Status          string          `json:"status"`
	PrincipalAmount float64         `json:"principal_amount"`
	Rate            float64         `json:"rate"`
	ROI             float64         `json:"roi"`
	Document        *DocumentDetail `json:"documents"`
}

func (r *DetailResponse) FromModel(cfg *config.GlobalConfig, l *loan.Loan) *DetailResponse {
	r.Document = &DocumentDetail{}
	if l.ApprovalDocument != nil {
		r.Document.ApprovalDocument = helper.ToPointer(fmt.Sprintf("%s/%s", cfg.PublicURL, *l.ApprovalDocument))
	}
	if l.DisbursalDocument != nil {
		r.Document.DisbursalDocument = helper.ToPointer(fmt.Sprintf("%s/%s", cfg.PublicURL, *l.DisbursalDocument))
	}
	for _, investor := range l.LoanInvestors {
		r.Document.InvestorAgreementLetters = append(
			r.Document.InvestorAgreementLetters,
			helper.ToPointer(fmt.Sprintf("%s/%s", cfg.PublicURL, *investor.AgreementLetter)))
	}
	return &DetailResponse{
		ID:              l.ID,
		Status:          l.Status,
		PrincipalAmount: l.PrincipalAmount,
		Rate:            l.Rate,
		ROI:             l.TotalAmount - l.PrincipalAmount,
		Document:        r.Document,
	}
}
