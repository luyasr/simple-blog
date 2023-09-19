package token

import "context"

type Service interface {
	Login(context.Context, *LoginRequest) error
	Logout(context.Context, *LogoutRequest) error
}

type LoginRequest struct {
}

type LogoutRequest struct {
}
