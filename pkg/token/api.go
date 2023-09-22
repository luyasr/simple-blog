package token

import (
	"github.com/gin-gonic/gin"
	"github.com/luyasr/simple-blog/pkg/response"
	"github.com/luyasr/simple-blog/pkg/user"
	"net/http"
)

type Handler struct {
	server Service
}

func NewHandler(serviceImpl *user.ServiceImpl) *Handler {
	return &Handler{
		server: NewServiceImpl(serviceImpl),
	}
}

func (h *Handler) Login(c *gin.Context) {
	req := NewLoginRequest()
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseWithError(err))
		return
	}
	token, err := h.server.Login(c.Request.Context(), req)
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
	err = h.server.Logout(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseWithError(err))
		return
	}
	c.JSON(http.StatusOK, response.NewResponse(nil))
}
