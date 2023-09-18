package user

import (
	"github.com/gin-gonic/gin"
)

func InitUserRoute(route *gin.RouterGroup) {
	userGroup := route.Group("user")
	{
		userGroup.POST("", CreateUser)
		userGroup.DELETE(":id", DeleteUser)
		userGroup.PUT(":id", UpdateUser)
		userGroup.GET(":username", DescribeUser)
	}
}
