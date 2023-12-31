package e

func NewAuthFailed(format string, a ...any) *Error {
	return New(401, format, a...)
}

func NewAccessDenied(format string, a ...any) *Error {
	return New(403, format, a...)
}

func NewNotFound(format string, a ...any) *Error {
	return New(404, format, a...)
}

func NewAlreadyExist(format string, a ...any) *Error {
	return New(10001, format, a...)
}

func NewUpdateFailed(format string, a ...any) *Error {
	return New(10002, format, a...)
}

func NewDeleteFailed(format string, a ...any) *Error {
	return New(10003, format, a...)
}

func NewAssertionFailed(format string, a ...any) *Error {
	return New(10004, format, a...)
}
