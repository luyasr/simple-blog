package api

import (
	"github.com/gin-gonic/gin"
	"github.com/luyasr/simple-blog/pkg/user"
	"github.com/luyasr/simple-blog/pkg/user/impl"
)

var userSvc = impl.NewUserServiceImpl()

func Ping(c *gin.Context) {
	c.JSON(200, Response{Code: 10000, Data: nil, Msg: "pong"})
}

func CreateUser(c *gin.Context) {
	req := user.NewCreateUserRequest()
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(200, Response{Code: 10001, Data: nil, Msg: err.Error()})
		return
	}
	createUser, err := userSvc.CreateUser(c.Request.Context(), req)
	if err != nil {
		c.JSON(200, Response{Code: 10002, Data: nil, Msg: err.Error()})
		return
	}
	c.JSON(200, Response{Code: 10000, Data: createUser, Msg: "success"})
}

func DeleteUser(c *gin.Context) {

}

func UpdateUser(c *gin.Context) {

}

func DescribeUser(c *gin.Context) {

}
