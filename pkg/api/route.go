package api

import "github.com/gin-gonic/gin"

func InitRoute(route *gin.RouterGroup) {
	route.GET("ping", Ping)
}
