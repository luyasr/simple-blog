package blog

import (
	"github.com/gin-gonic/gin"
	"github.com/luyasr/simple-blog/app/middlewares"
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
	group := r.Group("blog")
	group.Use()
	{
		group.GET("", h.QueryBlog)
		group.Use(middlewares.NewAuth().Auth)
		group.POST("", h.CreateBlog)
		group.DELETE(":id", h.DeleteBlog)
		group.PUT(":id", h.UpdateBlog)

	}
}

func (h *Handler) CreateBlog(c *gin.Context) {
	req := NewCreateBlogRequest()
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseWithError(err))
		return
	}

	blog, err := h.service.CreateBlog(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseWithError(err))
		return
	}
	c.JSON(http.StatusOK, response.NewResponse(blog))
}

func (h *Handler) DeleteBlog(c *gin.Context) {
	req := NewDeleteBlogRequest()
	req.Id = utils.StringToInt64(c.Param("id"))

	err := h.service.DeleteBlog(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseWithError(err))
		return
	}
	c.JSON(http.StatusOK, response.NewResponse(nil))
}

func (h *Handler) UpdateBlog(c *gin.Context) {
	req := NewUpdateBlogRequest()
	req.BlogId = utils.StringToInt64(c.Param("id"))
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseWithError(err))
		return
	}
	err = h.service.UpdateBlog(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseWithError(err))
		return
	}
	c.JSON(http.StatusOK, response.NewResponse(nil))
}

func (h *Handler) QueryBlog(c *gin.Context) {
	req := NewQueryBlogRequest()
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseWithError(err))
		return
	}

	blogs, err := h.service.QueryBlog(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseWithError(err))
		return
	}
	c.JSON(http.StatusOK, response.NewResponse(blogs))
}
