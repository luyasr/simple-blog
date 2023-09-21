package user

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) InitUserRoute(route *gin.RouterGroup) {
	{
		route.POST("", h.CreateUser)
		route.DELETE(":id", h.DeleteUser)
		route.PUT(":id", h.UpdateUser)
		route.GET(":username", h.DescribeUser)
	}
}
