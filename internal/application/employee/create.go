package employee

import (
	"context"
	"net/http"

	"github.com/Gash21/amartha-test/internal/application/employee/dto"
	"github.com/Gash21/amartha-test/internal/domain/employee"
	"github.com/Gash21/amartha-test/internal/shared/helper"
)

func (u *Usecase) Create(ctx context.Context, req dto.CreateRequest) helper.JSONResult {
	employee := u.EmployeeRepository.Create(&employee.Employee{
		Name: req.Name,
	})

	return helper.ResponseSuccess(http.StatusCreated, employee)
}
