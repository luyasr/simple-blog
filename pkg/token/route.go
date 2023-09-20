package token

import "github.com/gin-gonic/gin"

func (h *Handler) InitTokenRoute(route *gin.RouterGroup) {
	t := route.Group("token")
	t.Use()
	{
		t.POST("", h.Login)
	}
}
