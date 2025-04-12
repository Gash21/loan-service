package loan

import (
	"context"
	"net/http"

	"github.com/Gash21/amartha-test/internal/application/loan/dto"
	"github.com/Gash21/amartha-test/internal/shared/helper"
	"github.com/Gash21/amartha-test/internal/shared/logger"
	"go.uber.org/zap"
)

func (u *Usecase) Detail(ctx context.Context, req *dto.DetailRequest) helper.JSONResult {
	l := logger.WithId(u.Logger, ContextName, "Detail")
	loan, err := u.LoanRepository.FindByID(&req.ID)
	if err != nil {
		l.Error("failed to find loan", zap.Error(err))
		return helper.ResponseFailed(http.StatusNotFound, "Loan not found", nil)
	}

	res := &dto.DetailResponse{}
	res = res.FromLoan(loan)

	return helper.ResponseSuccess(http.StatusOK, res)
}
