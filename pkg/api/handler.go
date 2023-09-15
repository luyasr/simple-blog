package api

import "github.com/gin-gonic/gin"

func Ping(c *gin.Context) {
	c.JSON(200, Response{Code: 10000, Data: nil, Msg: "pong"})
}
