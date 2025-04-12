package dto

import (
	"mime/multipart"
	"time"
)

type ApproveRequest struct {
	ID             int64                 `params:"id"`
	EmployeeID     int64                 `form:"employee_id" validate:"required"`
	Evidence       *multipart.FileHeader `form:"evidence" validate:"required"`
	DateString     string                `form:"date_of_approval" validate:"required"`
	DateOfApproval time.Time             `form:"-"`
}

var ApproveAllowedMIMEs = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
}
