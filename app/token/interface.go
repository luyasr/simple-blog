package token

import "context"

const Name = "token"

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
	AccessToken  string `json:"access_token" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type ValidateToken struct {
	UserID      int64  `json:"user_id"`
	AccessToken string `json:"access_token"`
}

func NewLoginRequest() *LoginRequest {
	return &LoginRequest{}
}

func NewLogoutRequest() *LogoutRequest {
	return &LogoutRequest{}
}

func NewValidateToken(userId int64, token string) *ValidateToken {
	return &ValidateToken{
		UserID:      userId,
		AccessToken: token,
	}
}
