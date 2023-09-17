package api

import (
	"github.com/gin-gonic/gin"
	"github.com/luyasr/simple-blog/pkg/user"
	"github.com/luyasr/simple-blog/pkg/user/impl"
	"net/http"
)

var userSvc = impl.NewUserServiceImpl()

func Ping(c *gin.Context) {
	c.JSON(200, Response{Code: 0, Data: nil, Message: "pong"})
}

func CreateUser(c *gin.Context) {
	req := user.NewCreateUserRequest()
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusOK, NewResponseWithError(err))
		return
	}
	createUser, err := userSvc.CreateUser(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusOK, NewResponseWithError(err))
		return
	}
	c.JSON(200, NewResponse(createUser))
}

func DeleteUser(c *gin.Context) {

}

func UpdateUser(c *gin.Context) {

}

func DescribeUser(c *gin.Context) {

}
