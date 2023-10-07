package user

import (
	"github.com/gin-gonic/gin"
	"github.com/luyasr/simple-blog/pkg/ioc"
	"github.com/luyasr/simple-blog/pkg/response"
	"github.com/luyasr/simple-blog/pkg/utils"
	"net/http"
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

func (h *Handler) QueryUser(c *gin.Context) {
	id := c.Param("id")
	req := NewQueryUserRequestById(id)
	describeUser, err := h.service.QueryUser(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseWithError(err))
		return
	}
	c.JSON(http.StatusOK, response.NewResponse(describeUser))
}
