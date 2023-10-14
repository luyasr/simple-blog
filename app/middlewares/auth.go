package middlewares

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/luyasr/simple-blog/app/token"
	"github.com/luyasr/simple-blog/pkg/ioc"
	"github.com/luyasr/simple-blog/pkg/response"
	"net/http"
)

type Auth struct {
	token token.Service
}

func NewAuth() *Auth {
	return &Auth{
		token: ioc.Controller().Get(token.Name).(token.Service),
	}
}

func (a *Auth) Auth(c *gin.Context) {
	// 从cookie中获取access_token
	accessToken, err := c.Cookie(token.CookieName)
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			response.JSONWithError(c, token.CookieNotFound)
			return
		}
		response.JSONWithError(c, err)
		return
	}

	validateToken := token.NewValidateToken(accessToken)
	tk, err := a.token.Validate(c.Request.Context(), validateToken)
	if err != nil {
		response.JSONWithError(c, err)
		return
	}

	// 把鉴权后的token放入请求上下文中
	ctx := context.WithValue(c.Request.Context(), token.GinContextKey, tk)
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}
