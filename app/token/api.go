package token

import (
	"github.com/gin-gonic/gin"
	"github.com/luyasr/simple-blog/pkg/ioc"
	"github.com/luyasr/simple-blog/pkg/response"
	"github.com/luyasr/simple-blog/pkg/utils"
)

func init() {
	ioc.HttpHandler().Registry(&Handler{})
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
		group.DELETE(":id", h.Logout)
	}
}

func (h *Handler) Login(c *gin.Context) {
	req := NewLoginRequest()
	err := c.BindJSON(req)
	if err != nil {
		response.JSONWithError(c, err)
		return
	}
	ins, err := h.service.Login(c.Request.Context(), req)
	if err != nil {
		response.JSONWithError(c, err)
		return
	}

	// 通过SetCookie直接写入到浏览器客户端
	c.SetCookie(CookieName, ins.AccessToken, 0, "/", "", false, true)

	response.JSON(c, ins)
}

func (h *Handler) Logout(c *gin.Context) {
	req := NewLogoutOrRefreshRequest()
	req.UserId = utils.StringToInt64(c.Param("id"))
	err := c.BindJSON(req)
	if err != nil {
		response.JSONWithError(c, err)
		return
	}

	err = h.service.Logout(c.Request.Context(), req)
	if err != nil {
		response.JSONWithError(c, err)
		return
	}
	response.JSON(c, nil)
}
