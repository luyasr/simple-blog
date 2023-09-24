package token

import (
	"github.com/gin-gonic/gin"
	"github.com/luyasr/simple-blog/pkg/ioc"
	"github.com/luyasr/simple-blog/pkg/response"
	"net/http"
)

func init() {
	ioc.ApiHandler().Registry(&Handler{})
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
	group := r.Group("token")
	group.Use()
	{
		group.POST("", h.Login)
		group.DELETE("", h.Logout)
	}
}

func (h *Handler) Login(c *gin.Context) {
	req := NewLoginRequest()
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseWithError(err))
		return
	}
	token, err := h.service.Login(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseWithError(err))
		return
	}
	c.JSON(http.StatusOK, response.NewResponse(token))
}

func (h *Handler) Logout(c *gin.Context) {
	req := NewLogoutRequest()
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseWithError(err))
		return
	}
	err = h.service.Logout(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseWithError(err))
		return
	}
	c.JSON(http.StatusOK, response.NewResponse(nil))
}
