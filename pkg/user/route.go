package user

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) InitUserRoute(route *gin.RouterGroup) {
	userGroup := route.Group("user")
	{
		userGroup.POST("", h.CreateUser)
		userGroup.DELETE(":id", h.DeleteUser)
		userGroup.PUT(":id", h.UpdateUser)
		userGroup.GET(":username", h.DescribeUser)
	}
}
