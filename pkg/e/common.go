package e

func NewNotFound(format string, a ...any) *Error {
	return New(404, format, a...)
}

func NewUserExists(format string, a ...any) *Error {
	return New(10001, format, a...)
}

func NewUserNotExists(format string, a ...any) *Error {
	return New(10002, format, a...)
}

func NewUpdateFailed(format string, a ...any) *Error {
	return New(10003, format, a...)
}

func NewAuthFailed(format string, a ...any) *Error {
	return New(401, format, a...)
}
