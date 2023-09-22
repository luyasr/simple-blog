package token

import "github.com/gin-gonic/gin"

func (h *Handler) InitTokenRoute(route *gin.RouterGroup) {
	{
		route.POST("", h.Login)
		route.DELETE("", h.Logout)
	}
}
