package loan

import (
	"context"

	"github.com/Gash21/amartha-test/internal/application/loan/dto"
	"github.com/Gash21/amartha-test/internal/domain/borrower"
	"github.com/Gash21/amartha-test/internal/domain/employee"
	"github.com/Gash21/amartha-test/internal/domain/investor"
	"github.com/Gash21/amartha-test/internal/domain/loan"
	"github.com/Gash21/amartha-test/internal/shared/config"
	"github.com/Gash21/amartha-test/internal/shared/helper"
	"github.com/Gash21/amartha-test/internal/shared/mailer"
	"go.uber.org/zap"
)

const ContextName = "Application.Loan"

type (
	Usecase struct {
		LoanRepository     loan.LoanRepository
		EmployeeRepository employee.EmployeeRepository
		InvestorRepository investor.InvestorRepository
		BorrowerRepository borrower.BorrowerRepository
		Config             config.GlobalConfig
		Mailer             mailer.IMailerAdapter
		Logger             *zap.Logger
	}
	IUsecase interface {
		Propose(context.Context, *dto.ProposedRequest) helper.JSONResult
		List(context.Context, *dto.ListRequest) helper.JSONResult
		Approve(context.Context, *dto.ApproveRequest) helper.JSONResult
		Invest(context.Context, *dto.InvestRequest) helper.JSONResult
		Disburse(context.Context, *dto.DisburseRequest) helper.JSONResult
		Detail(context.Context, *dto.DetailRequest) helper.JSONResult
	}
)

func NewUsecase(opts Usecase) IUsecase {
	return &Usecase{
		LoanRepository:     opts.LoanRepository,
		InvestorRepository: opts.InvestorRepository,
		BorrowerRepository: opts.BorrowerRepository,
		EmployeeRepository: opts.EmployeeRepository,
		Logger:             opts.Logger,
		Mailer:             opts.Mailer,
		Config:             opts.Config,
	}
}
