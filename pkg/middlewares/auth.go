package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/luyasr/simple-blog/pkg/ioc"
	"github.com/luyasr/simple-blog/pkg/response"
	"github.com/luyasr/simple-blog/pkg/token"
	"github.com/luyasr/simple-blog/pkg/utils"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := c.GetHeader("token")
		id := c.Param("id")
		validateToken := token.NewValidateToken(utils.StringToInt64(id), t)
		err := ioc.Controller().Get("token").(token.Service).Validate(c.Request.Context(), validateToken)
		if err != nil {
			c.JSON(http.StatusOK, response.NewResponseWithError(err))
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}
