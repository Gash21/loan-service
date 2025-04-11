package loan

import (
	"context"
	"net/http"

	"github.com/Gash21/amartha-test/internal/application/loan/dto"
	"github.com/Gash21/amartha-test/internal/domain/loan"
	"github.com/Gash21/amartha-test/internal/shared/helper"
)

func (u *Usecase) Propose(ctx context.Context, req dto.ProposedRequest) helper.JSONResult {
	debtor, err := u.BorrowerRepository.FindByID(req.BorrowerID)
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

	loan := u.LoanRepository.Create(&loanData)
	resp := dto.ProposeResponse{
		LoanID: *loan.ID,
	}

	return helper.ResponseSuccess(http.StatusOK, resp)
}
