package helper

import (
	"math"
	"net/http"
)

type JSONResult struct {
	Code    int             `json:"-"`
	Message string          `json:"message"`
	Meta    *PaginationMeta `json:"meta,omitempty"`
	Data    interface{}     `json:"data"`
}

type PaginationMeta struct {
	Page            int `json:"page"`
	TotalData       int `json:"total_data"`
	TotalPage       int `json:"total_page"`
	TotalDataOnPage int `json:"total_data_on_page"`
}

func ResponseSuccess(code int, data interface{}) JSONResult {
	return JSONResult{
		Code:    code,
		Message: "Success",
		Data:    data,
	}
}

func ResponseFailed(httpCode int, message string, data interface{}) JSONResult {
	return JSONResult{
		Code:    httpCode,
		Message: message,
		Data:    data,
	}
}

func ResponsePagination(page int, limit int, count int, total int, data interface{}) JSONResult {
	return JSONResult{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    data,
		Meta: &PaginationMeta{
			Page:            page,
			TotalData:       total,
			TotalDataOnPage: count,
			TotalPage:       int(math.Ceil(float64(total) / float64(limit))),
		},
	}
}
