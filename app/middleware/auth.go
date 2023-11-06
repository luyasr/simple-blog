package middleware

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/luyasr/simple-blog/app/token"
	"github.com/luyasr/simple-blog/app/user"
	"github.com/luyasr/simple-blog/pkg/e"
	"github.com/luyasr/simple-blog/pkg/ioc"
	"github.com/luyasr/simple-blog/pkg/response"
	"net/http"
)

type Auth struct {
	token token.Service
	Role  user.Role
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
	if accessToken == "" {
		accessToken = c.Request.Header.Get(token.CookieName)
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

func (a *Auth) rolePermissions(c *gin.Context) {
	tk, err := a.token.GetTokenByContext(c.Request.Context())
	if err != nil {
		response.JSONWithError(c, err)
		return
	}

	if tk.Role == user.RoleAdmin {
		return
	}

	if tk.Role != a.Role {
		response.JSONWithError(c, e.NewAccessDenied("role %v not allow", a.Role))
		return
	}
}

func RolePermissions(role user.Role) gin.HandlerFunc {
	a := NewAuth()
	a.Role = role

	return a.rolePermissions
}
