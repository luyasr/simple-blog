package token

import "context"

type Service interface {
	Login(context.Context, *LoginRequest) (*Token, error)
	Logout(context.Context, *LogoutRequest) error
	Validate(context.Context, *ValidateToken) error
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=3,max=20" label:"用户名"`
	Password string `json:"password" validate:"required,min=6,max=20" label:"密码"`
}

type LogoutRequest struct {
}

type ValidateToken struct {
	AccessToken string `json:"access_token" validate:"required"`
}

func NewLoginRequest() *LoginRequest {
	return &LoginRequest{}
}

func NewValidateToken(token string) *ValidateToken {
	return &ValidateToken{
		AccessToken: token,
	}
}
