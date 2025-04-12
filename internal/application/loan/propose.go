package loan

import (
	"context"
	"net/http"

	"github.com/Gash21/amartha-test/internal/application/loan/dto"
	"github.com/Gash21/amartha-test/internal/domain/loan"
	"github.com/Gash21/amartha-test/internal/shared/helper"
	"github.com/Gash21/amartha-test/internal/shared/logger"
	"go.uber.org/zap"
)

func (u *Usecase) Propose(ctx context.Context, req *dto.ProposedRequest) helper.JSONResult {
	l := logger.WithId(u.Logger, ContextName, "Propose")
	debtor, err := u.BorrowerRepository.FindByID(&req.BorrowerID)
	if err != nil {
		return helper.ResponseFailed(http.StatusNotFound, "Borrower not found", nil)
	}
	interest := req.PrincipalAmount * req.Rate / 100

	loanData := loan.Loan{
		BorrowerID:      debtor.ID,
		PrincipalAmount: req.PrincipalAmount,
		Rate:            req.Rate,
		Status:          "proposed",
		TotalAmount:     req.PrincipalAmount + interest,
	}

	_, err = u.LoanRepository.Create(&loanData)
	if err != nil {
		l.Error("failed to create loan", zap.Error(err))
		return helper.ResponseFailed(http.StatusInternalServerError, "failed to create loan", nil)
	}

	return helper.ResponseSuccess(http.StatusOK, dto.ProposeSuccessResponse)
}
