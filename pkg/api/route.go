package api

import "github.com/gin-gonic/gin"

func InitRoute(route *gin.RouterGroup) {
	route.GET("ping", Ping)
	user := route.Group("user")
	{
		user.GET("", UserList)
		user.POST("id", CreateUser)
	}
}
