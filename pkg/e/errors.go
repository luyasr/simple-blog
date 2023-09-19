package e

import (
	"errors"
	"fmt"
)

const (
	defaultErrorCode = iota + 1
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func New(code int, format string, a ...any) *Error {
	return &Error{
		Code:    code,
		Message: fmt.Sprintf(format, a...),
	}
}

func NewCode(err error) int {
	var eError *Error
	if errors.As(err, &eError) {
		return eError.Code
	}
	return defaultErrorCode
}

func NewMessage(err error, obj any) string {
	var eError *Error
	if errors.As(err, &eError) {
		return eError.Error()
	}

	return err.Error()
}

func (e *Error) Error() string {
	return e.Message
}
