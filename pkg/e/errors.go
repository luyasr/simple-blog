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
	var e *Error
	if errors.As(err, &e) {
		return e.Code
	}
	return defaultErrorCode
}

func (e *Error) Error() string {
	return e.Message
}
