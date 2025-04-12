package loan

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Gash21/amartha-test/internal/application/loan/dto"
	"github.com/Gash21/amartha-test/internal/shared/helper"
	"github.com/Gash21/amartha-test/internal/shared/logger"
	"go.uber.org/zap"
)

func (u *Usecase) Approve(ctx context.Context, req *dto.ApproveRequest) helper.JSONResult {
	l := logger.WithId(u.Logger, ContextName, "Approve")
	loan, err := u.LoanRepository.FindByIDAndStatus(&req.ID, "proposed")
	if err != nil {
		l.Error("failed to find loan", zap.Error(err))
		return helper.ResponseFailed(http.StatusNotFound, "Loan not found", nil)
	}

	employee, err := u.EmployeeRepository.FindByID(&req.EmployeeID)
	if err != nil {
		l.Error("failed to find employee", zap.Error(err))
		return helper.ResponseFailed(http.StatusNotFound, "Employee not found", nil)
	}

	mimeType := req.Evidence.Header.Get("Content-Type")
	if !dto.ApproveAllowedMIMEs[mimeType] {
		l.Error("unsupported file type", zap.String("mimeType", mimeType))
		return helper.ResponseFailed(http.StatusBadRequest, "Unsupported file type", nil)
	}

	file, err := req.Evidence.Open()
	if err != nil {
		l.Error("failed to open uploaded file", zap.Error(err))
		return helper.ResponseFailed(http.StatusInternalServerError, "failed to open uploaded file", nil)
	}
	defer file.Close()

	savePath := fmt.Sprintf("./uploads/%s", fmt.Sprintf("%d_%s_%s", req.ID, "approve", req.Evidence.Filename))
	dst, err := os.Create(savePath)
	if err != nil {
		l.Error("failed to create destination file", zap.Error(err))
		return helper.ResponseFailed(http.StatusInternalServerError, "failed to create destination file", nil)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		l.Error("failed to save file", zap.Error(err))
		return helper.ResponseFailed(http.StatusInternalServerError, "failed to save file", nil)
	}

	loan.ApprovedAt = &req.DateOfApproval
	loan.ApprovedBy = employee.ID
	loan.Status = "approved"
	loan.ApprovalDocument = helper.ToPointer(strings.Replace(savePath, "./", "", 1))

	if _, err := u.LoanRepository.Update(loan); err != nil {
		l.Error("failed to update loan", zap.Error(err))
		return helper.ResponseFailed(http.StatusInternalServerError, "failed to update loan", nil)
	}

	return helper.ResponseSuccess(http.StatusOK, loan)
}
