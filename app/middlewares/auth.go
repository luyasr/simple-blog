package middlewares

import (
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
	// 获取cookie
	cookie, err := c.Cookie(token.CookieName)
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			c.JSON(http.StatusOK, token.CookieNotFound)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, response.NewResponseWithError(err))
		c.Abort()
		return
	}

	validateToken := token.NewValidateToken(cookie)
	tk, err := a.token.Validate(c.Request.Context(), validateToken)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseWithError(err))
		c.Abort()
		return
	}

	// 把鉴权后的token放入请求上下文中
	if c.Keys == nil {
		c.Keys = map[string]any{}
	}
	c.Keys[token.GinContextKey] = tk
	c.Next()
}
