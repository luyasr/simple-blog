package token

import "github.com/luyasr/simple-blog/pkg/e"

type QueryBy int

const (
	QueryByUserId QueryBy = iota
	QueryByAccessToken
	QueryByLogoutRequest
)

const (
	CookieName    = "access_token"
	GinContextKey = "token"
)

var (
	CookieNotFound = e.NewAuthFailed("cookie %s not found", CookieName)
	NotFound       = e.NewNotFound("token not found")
	AuthFailed     = e.NewAuthFailed("incorrect account number or password")
	InvalidToken   = e.NewAuthFailed("invalid token")
	ExpiresToken   = e.NewAuthFailed("the token has expired please log in again")
)
