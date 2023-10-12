package token

import "github.com/luyasr/simple-blog/pkg/e"

const (
	CookieName    = "access_token"
	GinContextKey = "token"
)

var (
	CookieNotFound = e.NewAuthFailed("cookie %s not found", CookieName)
)
