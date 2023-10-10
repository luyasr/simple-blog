package token

import (
	"github.com/gin-gonic/gin"
	"github.com/luyasr/simple-blog/pkg/ioc"
	"github.com/luyasr/simple-blog/pkg/response"
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
	ins, err := h.service.Login(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseWithError(err))
		return
	}

	// 通过SetCookie直接写入到浏览器客户端
	c.SetCookie(tokenCookieName, ins.AccessToken, 0, "/", "localhost", false, true)

	c.JSON(http.StatusOK, response.NewResponse(ins))
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
