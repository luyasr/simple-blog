package user

import (
	"github.com/gin-gonic/gin"
	"github.com/luyasr/simple-blog/pkg/response"
	"github.com/luyasr/simple-blog/pkg/utils"
	"net/http"
)

var userSvc = NewUserServiceImpl()

func CreateUser(c *gin.Context) {
	req := NewCreateUserRequest()
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseWithError(err, req))
		return
	}
	createUser, err := userSvc.CreateUser(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseWithError(err, nil))
		return
	}
	c.JSON(200, response.NewResponse(createUser))
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	req := NewDeleteUserRequest(id)
	err := userSvc.DeleteUser(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseWithError(err, nil))
		return
	}
	c.JSON(http.StatusOK, response.NewResponse(req))
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	req := NewDefaultUser()
	req.ID = utils.StringToInt64(id)

	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseWithError(err, req))
		return
	}
	err = userSvc.UpdateUser(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseWithError(err, req))
		return
	}
	c.JSON(http.StatusOK, response.NewResponse(nil))
}

func DescribeUser(c *gin.Context) {
	username := c.Param("username")
	req := NewDescribeUserRequestByUsername(username)
	describeUser, err := userSvc.DescribeUser(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseWithError(err, nil))
		return
	}
	c.JSON(http.StatusOK, response.NewResponse(describeUser))
}
