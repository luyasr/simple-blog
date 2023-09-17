package e

import "fmt"

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

func (e *Error) Error() string {
	return e.Message
}
