package blog

import (
	"github.com/gin-gonic/gin"
	"github.com/luyasr/simple-blog/app/middleware"
	"github.com/luyasr/simple-blog/app/user"
	"github.com/luyasr/simple-blog/pkg/ioc"
	"github.com/luyasr/simple-blog/pkg/response"
	"github.com/luyasr/simple-blog/pkg/utils"
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
		group.Use(middleware.NewAuth().Auth, middleware.RolePermissions(user.RoleAuthor))
		group.POST("", h.CreateBlog)
		group.DELETE(":id", h.DeleteBlog)
		group.PUT(":id", h.UpdateBlog)
		group.POST("publish/:id", h.UpdateBlogStatus)

	}
}

func (h *Handler) CreateBlog(c *gin.Context) {
	req := NewCreateBlogRequest()
	err := c.BindJSON(req)
	if err != nil {
		response.JSONWithError(c, err)
		return
	}

	blog, err := h.service.CreateBlog(c.Request.Context(), req)
	if err != nil {
		response.JSONWithError(c, err)
		return
	}
	response.JSON(c, blog)
}

func (h *Handler) DeleteBlog(c *gin.Context) {
	req := NewDeleteBlogRequest()
	req.Id = utils.StringToInt64(c.Param("id"))

	err := h.service.DeleteBlog(c.Request.Context(), req)
	if err != nil {
		response.JSONWithError(c, err)
		return
	}
	response.JSON(c, nil)
}

func (h *Handler) UpdateBlog(c *gin.Context) {
	req := NewUpdateBlogRequest()
	req.Id = utils.StringToInt64(c.Param("id"))
	err := c.BindJSON(req)
	if err != nil {
		response.JSONWithError(c, err)
		return
	}

	err = h.service.UpdateBlog(c.Request.Context(), req)
	if err != nil {
		response.JSONWithError(c, err)
		return
	}
	response.JSON(c, nil)
}

func (h *Handler) QueryBlog(c *gin.Context) {
	req := NewQueryBlogRequest()
	switch c.Query("status") {
	case "draft":
		req.SetStatus(StatusDraft)
	case "published":
		req.SetStatus(StatusPublished)
	}
	req.ParsePageSize(c.Query("page_size"))
	req.ParsePageNumber(c.Query("page_number"))

	blogs, err := h.service.QueryBlog(c.Request.Context(), req)
	if err != nil {
		response.JSONWithError(c, err)
		return
	}
	response.JSON(c, blogs)
}

func (h *Handler) UpdateBlogStatus(c *gin.Context) {
	req := NewUpdateBlogStatusRequest()
	req.Id = utils.StringToInt64(c.Param("id"))
	req.Status = StatusPublished

	err := h.service.UpdateBlogStatus(c.Request.Context(), req)
	if err != nil {
		response.JSONWithError(c, err)
		return
	}
	response.JSON(c, nil)
}

func (h *Handler) AuditBlog(c *gin.Context) {
	req := NewAuditBlogRequest()
	req.Id = utils.StringToInt64(c.Param("id"))
	err := c.BindJSON(req)
	if err != nil {
		response.JSONWithError(c, err)
		return
	}
	err = h.service.AuditBlog(c.Request.Context(), req)
	if err != nil {
		response.JSONWithError(c, err)
		return
	}
	response.JSON(c, nil)
}
