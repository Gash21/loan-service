package loan

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Gash21/amartha-test/internal/application/loan/dto"
	"github.com/Gash21/amartha-test/internal/domain/loan_investor"
	"github.com/Gash21/amartha-test/internal/shared/helper"
	"github.com/Gash21/amartha-test/internal/shared/logger"
	"github.com/Gash21/amartha-test/internal/shared/pdf"
	"go.uber.org/zap"
)

func (u *Usecase) Invest(ctx context.Context, req *dto.InvestRequest) helper.JSONResult {
	l := logger.WithId(u.Logger, ContextName, "Invest")
	loanData, err := u.LoanRepository.FindInvestableLoanByID(req.ID)
	if err != nil {
		l.Error("failed to find loan", zap.Error(err))
		return helper.ResponseFailed(http.StatusNotFound, "Loan not found", nil)
	}

	investorData, err := u.InvestorRepository.FindByID(req.InvestorID)
	if err != nil {
		l.Error("failed to find investor", zap.Error(err))
		return helper.ResponseFailed(http.StatusNotFound, "Investor not found", nil)
	}

	if loanData.TotalInvested+req.Amount > loanData.PrincipalAmount {
		l.Error("investment amount exceeds loan amount")
		return helper.ResponseFailed(http.StatusBadRequest, "Investment amount exceeds loan amount", nil)
	}

	loanData.TotalInvested += req.Amount
	if loanData.PrincipalAmount == loanData.TotalInvested {
		loanData.Status = "ready_to_disbursed"
	}

	timestamp := time.Now().Format("20060102150405")
	text := fmt.Sprintf("Dear Mr/Mrs. %s, Your money successfully invested on loan %d,\nplease read this agreement for details.", investorData.Name, *loanData.ID)
	savePath := fmt.Sprintf("./uploads/%s_%d_%d_%s_%s.pdf", timestamp, *loanData.ID, *req.InvestorID, "invest", "agreement_letter")
	err = pdf.GeneratePDF(text, savePath)
	if err != nil {
		l.Error("failed to generate agreement letter", zap.Error(err))
		return helper.ResponseFailed(http.StatusInternalServerError, "failed to generate agreement letter", nil)
	}
	loanData.LoanInvestors = append(loanData.LoanInvestors, &loan_investor.LoanInvestor{
		LoanID:          loanData.ID,
		InvestorID:      investorData.ID,
		Amount:          req.Amount,
		Rate:            loanData.Rate,
		ROI:             req.Amount + (req.Amount * loanData.Rate / 100),
		AgreementLetter: helper.ToPointer(strings.Replace(savePath, "./", "", 1)),
	})

	_, err = u.LoanRepository.Update(loanData)
	if err != nil {
		l.Error("failed to update loan", zap.Error(err))
		return helper.ResponseFailed(http.StatusInternalServerError, "failed to update loan", nil)
	}
	return helper.ResponseSuccess(http.StatusOK, loanData)
}
