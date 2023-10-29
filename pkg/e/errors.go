package e

import (
	"errors"
	"fmt"
	"net/http"
)

type Error struct {
	HttpCode int    `json:"http_code"`
	BizCode  int    `json:"biz_code"`
	Message  string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

func New(bizCode int, format string, a ...any) *Error {
	var httpCode int

	if bizCode < 600 {
		httpCode = bizCode
	} else {
		httpCode = http.StatusBadRequest
	}

	return &Error{
		HttpCode: httpCode,
		BizCode:  bizCode,
		Message:  fmt.Sprintf(format, a...),
	}
}

func GetErrorInfo(err error) (int, int, string) {
	var eError *Error
	if errors.As(err, &eError) {
		return eError.HttpCode, eError.BizCode, eError.Error()
	}
	return http.StatusBadRequest, http.StatusBadRequest, err.Error()
}
