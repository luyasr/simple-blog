package api

import (
	"github.com/gin-gonic/gin"
	"github.com/luyasr/simple-blog/pkg/response"
)

func Ping(c *gin.Context) {
	c.JSON(200, response.Response{Code: 0, Data: nil, Message: "pong"})
}
