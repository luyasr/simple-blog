package response

import (
	"github.com/gin-gonic/gin"
	"github.com/luyasr/simple-blog/pkg/e"
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func JSON(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{Code: 0, Data: data, Message: "success"})
}

func JSONWithError(c *gin.Context, err error) {
	defer c.Abort()
	c.JSON(http.StatusOK, Response{Code: e.GetCode(err), Data: nil, Message: e.GetMessage(err)})
}
