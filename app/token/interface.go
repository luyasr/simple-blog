package token

import "context"

const Name = "token"

type Service interface {
	Login(context.Context, *LoginRequest) (*Token, error)
	Logout(context.Context, *LogoutOrRefreshRequest) error
	Query(context.Context, *QueryTokenRequest) (*Token, error)
	Refresh(context.Context, *LogoutOrRefreshRequest) (*Token, error)
	Validate(context.Context, *ValidateToken) (*Token, error)
	GetTokenByContext(context.Context) (*Token, error)
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=3,max=20" label:"用户名"`
	Password string `json:"password" validate:"required,min=6,max=20" label:"密码"`
}

type LogoutOrRefreshRequest struct {
	UserId       int64  `json:"user_id" validate:"required" label:"用户ID"`
	AccessToken  string `json:"access_token" validate:"required" label:"登录token"`
	RefreshToken string `json:"refresh_token" validate:"required" label:"刷新token"`
}

type QueryTokenRequest struct {
	QueryBy      QueryBy `json:"query_by"`
	QueryByValue []any   `json:"query_by_value"`
}

func NewQueryTokenByUserIdRequest(id string) *QueryTokenRequest {
	return &QueryTokenRequest{
		QueryBy:      QueryByUserId,
		QueryByValue: []any{id},
	}
}

func NewQueryTokenByAccessTokenRequest(accessToken string) *QueryTokenRequest {
	return &QueryTokenRequest{
		QueryBy:      QueryByAccessToken,
		QueryByValue: []any{accessToken},
	}
}

func NewQueryByUARRequest(userId int64, accessToken string, refreshToken string) *QueryTokenRequest {
	return &QueryTokenRequest{
		QueryBy:      QueryByLogoutRequest,
		QueryByValue: []any{userId, accessToken, refreshToken},
	}
}

type ValidateToken struct {
	AccessToken string `json:"access_token" validate:"required" label:"登录token"`
}

func NewLoginRequest() *LoginRequest {
	return &LoginRequest{}
}

func NewLogoutOrRefreshRequest() *LogoutOrRefreshRequest {
	return &LogoutOrRefreshRequest{}
}

func NewValidateToken(token string) *ValidateToken {
	return &ValidateToken{
		AccessToken: token,
	}
}
