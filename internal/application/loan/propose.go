package loan

import (
	"context"
	"net/http"

	"github.com/Gash21/amartha-test/internal/application/loan/dto"
	"github.com/Gash21/amartha-test/internal/shared/helper"
)

func (u *Usecase) Propose(ctx context.Context, req dto.ProposedRequest) helper.JSONResult {
	return helper.ResponseSuccess(http.StatusOK, "Loan proposed successfully")
}
