package user

import (
	"github.com/gin-gonic/gin"
	"github.com/luyasr/simple-blog/pkg/ioc"
	"github.com/luyasr/simple-blog/pkg/response"
)

func init() {
	ioc.Handler().Registry(&Handler{})
}

func (h *Handler) Init() error {
	h.service = ioc.Controller().Get(Name).(Service)
	return nil
}

func (h *Handler) Name() string {
	return Name
}

type Handler struct {
	service Service
}

func (h *Handler) Registry(r gin.IRouter) {
	group := r.Group("user")
	group.Use()
	{
		group.POST("", h.CreateUser)
		group.DELETE(":id", h.DeleteUser)
		group.PUT(":id", h.UpdateUser)
		group.GET(":id", h.QueryUser)
	}
}

func (h *Handler) CreateUser(c *gin.Context) {
	req := NewCreateUserRequest()
	err := c.BindJSON(req)
	if err != nil {
		response.JSONWithError(c, err)
		return
	}
	createUser, err := h.service.CreateUser(c.Request.Context(), req)
	if err != nil {
		response.JSONWithError(c, err)
		return
	}
	response.JSON(c, createUser)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	req := NewDeleteUserRequest(id)
	err := h.service.DeleteUser(c.Request.Context(), req)
	if err != nil {
		response.JSONWithError(c, err)
		return
	}
	response.JSON(c, nil)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	req := NewUpdateUser()
	req.Id = c.Param("id")

	err := c.BindJSON(req)
	if err != nil {
		response.JSONWithError(c, err)
		return
	}

	err = h.service.UpdateUser(c.Request.Context(), req)
	if err != nil {
		response.JSONWithError(c, err)
		return
	}
	response.JSON(c, nil)
}

func (h *Handler) QueryUser(c *gin.Context) {
	req := NewQueryUserRequestById(c.Param("id"))
	user, err := h.service.QueryUser(c.Request.Context(), req)
	if err != nil {
		response.JSONWithError(c, err)
		return
	}
	response.JSON(c, user)
}
