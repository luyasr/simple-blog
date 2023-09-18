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
		c.JSON(http.StatusOK, NewResponseWithError(err, req))
		return
	}
	createUser, err := userSvc.CreateUser(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusOK, NewResponseWithError(err, nil))
		return
	}
	c.JSON(200, NewResponse(createUser))
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	req := user.NewDeleteUserRequest(id)
	err := userSvc.DeleteUser(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusOK, NewResponseWithError(err, nil))
		return
	}
	c.JSON(http.StatusOK, NewResponse(req))
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	req := user.NewUpdateUserRequest(id)
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusOK, NewResponseWithError(err, req))
		return
	}
	err = userSvc.UpdateUser(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusOK, NewResponseWithError(err, req))
		return
	}
	c.JSON(http.StatusOK, NewResponse(nil))
}

func DescribeUser(c *gin.Context) {
	username := c.Param("username")
	req := user.NewDescribeUserRequestByUsername(username)
	describeUser, err := userSvc.DescribeUser(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusOK, NewResponseWithError(err, nil))
		return
	}
	c.JSON(http.StatusOK, NewResponse(describeUser))
}
