package dto

import (
	"mime/multipart"
	"time"
)

type DisburseRequest struct {
	ID                 int64                 `params:"id"`
	Evidence           *multipart.FileHeader `form:"evidence" validate:"required"`
	EmployeeID         int64                 `form:"employee_id" validate:"required"`
	DateString         string                `form:"date_of_disbursement" validate:"required"`
	DateOfDisbursement time.Time             `form:"-"`
}

var DisburseAllowedMIMEs = map[string]bool{
	"image/jpeg":      true,
	"image/png":       true,
	"application/pdf": true,
}
