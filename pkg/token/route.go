package token

import "github.com/gin-gonic/gin"

func (h *Handler) InitTokenRoute(route *gin.RouterGroup) {
	tokenGroup := route.Group("token")
	tokenGroup.Use()
	{
		tokenGroup.POST("", h.Login)
	}
}
