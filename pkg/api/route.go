package api

import "github.com/gin-gonic/gin"

func InitRoute(route *gin.RouterGroup) {
	route.GET("ping", Ping)
	user := route.Group("user")
	{
		user.POST("", CreateUser)
		user.DELETE(":id", DeleteUser)
		user.PUT(":id", UpdateUser)
		user.GET(":username", DescribeUser)
	}
}
