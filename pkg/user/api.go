package user

import (
	"github.com/gin-gonic/gin"
	"github.com/luyasr/simple-blog/pkg/response"
	"github.com/luyasr/simple-blog/pkg/utils"
	"net/http"
)

type Handler struct {
	service Service
}

func NewHandler() *Handler {
	return &Handler{
		service: NewServiceImpl(),
	}
}

func (h *Handler) CreateUser(c *gin.Context) {
	req := NewCreateUserRequest()
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseWithError(err))
		return
	}
	createUser, err := h.service.CreateUser(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseWithError(err))
		return
	}
	c.JSON(200, response.NewResponse(createUser))
}

func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	req := NewDeleteUserRequest(id)
	err := h.service.DeleteUser(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseWithError(err))
		return
	}
	c.JSON(http.StatusOK, response.NewResponse(nil))
}

func (h *Handler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	req := NewUpdateUser()
	req.ID = utils.StringToInt64(id)

	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseWithError(err))
		return
	}
	err = h.service.UpdateUser(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseWithError(err))
		return
	}
	c.JSON(http.StatusOK, response.NewResponse(nil))
}

func (h *Handler) DescribeUser(c *gin.Context) {
	username := c.Param("username")
	req := NewDescribeUserRequestByUsername(username)
	describeUser, err := h.service.DescribeUser(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseWithError(err))
		return
	}
	c.JSON(http.StatusOK, response.NewResponse(describeUser))
}
