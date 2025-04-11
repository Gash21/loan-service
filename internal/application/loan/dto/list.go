package dto

import "github.com/Gash21/amartha-test/internal/domain/loan"

type ListRequest struct {
	Page  int `query:"page"`
	Limit int `query:"limit"`
}

type ListResponse struct {
	Data  []loan.Loan
	Total int64
}
