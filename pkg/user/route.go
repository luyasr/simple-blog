package user

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUserRoute(route *gin.RouterGroup) {
	{
		route.POST("", h.CreateUser)
	}
}

func (h *Handler) UserRoute(route *gin.RouterGroup) {
	{
		route.DELETE(":id", h.DeleteUser)
		route.PUT(":id", h.UpdateUser)
		route.GET(":id", h.DescribeUser)
	}
}
