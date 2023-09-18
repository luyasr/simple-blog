package e

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
)

const (
	defaultErrorCode = iota + 1
)

const (
	defaultErrorMessage = "failed"
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

func NewMessage(err error, obj any) string {
	var eError *Error
	if errors.As(err, &eError) {
		return err.Error()
	}

	typeOf := reflect.TypeOf(obj)
	if vErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range vErrors {
			if s, ok := typeOf.Elem().FieldByName(e.Field()); ok {
				msg := s.Tag.Get("msg")
				return fmt.Sprintf("%s %s", msg, e.Error())
			}
		}
	}
	return defaultErrorMessage
}

func (e *Error) Error() string {
	return e.Message
}
