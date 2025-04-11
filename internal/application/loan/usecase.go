package loan

import (
	"context"

	"github.com/Gash21/amartha-test/internal/application/loan/dto"
	"github.com/Gash21/amartha-test/internal/domain/borrower"
	"github.com/Gash21/amartha-test/internal/domain/document"
	"github.com/Gash21/amartha-test/internal/domain/loan"
	"github.com/Gash21/amartha-test/internal/shared/helper"
	"go.uber.org/zap"
)

type (
	Usecase struct {
		LoanRepository     loan.LoanRepository
		DocumentRepository document.DocumentRepository
		BorrowerRepository borrower.BorrowerRepository
		Logger             *zap.Logger
	}
	IUsecase interface {
		Propose(context.Context, dto.ProposedRequest) helper.JSONResult
		List(context.Context, dto.ListRequest) helper.JSONResult
		Approve() helper.JSONResult
		Invest() helper.JSONResult
		Disburse() helper.JSONResult
	}
)

func NewUsecase(opts Usecase) IUsecase {
	return &Usecase{
		LoanRepository:     opts.LoanRepository,
		BorrowerRepository: opts.BorrowerRepository,
		DocumentRepository: opts.DocumentRepository,
		Logger:             opts.Logger,
	}
}
