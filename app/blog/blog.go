package blog

import (
	"context"
	"dario.cat/mergo"
	"github.com/luyasr/simple-blog/app/token"
	"github.com/luyasr/simple-blog/config"
	"github.com/luyasr/simple-blog/pkg/ioc"
	"github.com/luyasr/simple-blog/pkg/utils"
	"github.com/luyasr/simple-blog/pkg/validate"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"time"
)

var (
	_ Service = (*ServiceImpl)(nil)
)

func init() {
	ioc.Controller().Registry(&ServiceImpl{})
}

func (s *ServiceImpl) Init() error {
	s.db = config.C.Mysql.GetConn()
	s.token = ioc.Controller().Get(token.Name).(token.Service)
	return nil
}

func (s *ServiceImpl) Name() string {
	return Name
}

type ServiceImpl struct {
	db    *gorm.DB
	token token.Service
	log   zerolog.Logger
}

func (s *ServiceImpl) CreateBlog(ctx context.Context, req *CreateBlogRequest) (*Blog, error) {
	if err := validate.Struct(req); err != nil {
		return nil, err
	}

	blog := NewBlog(req)

	// 从上下文中获取用户名称作为博客创建者
	tk, err := s.token.GetTokenByContext(ctx)
	if err != nil {
		return nil, err
	}
	req.CreateBy = tk.Username

	err = s.db.WithContext(ctx).Create(blog).Error
	if err != nil {
		return nil, err
	}
	return blog, nil
}

func (s *ServiceImpl) DeleteBlog(ctx context.Context, req *DeleteBlogRequest) error {
	if err := validate.Struct(req); err != nil {
		return err
	}

	// 删除前先查询是否存在文章
	blog, err := s.QueryBlogById(ctx, NewQueryBlogByIdRequest(req.Id))
	if err != nil {
		return err
	}

	err = s.db.WithContext(ctx).Delete(blog).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *ServiceImpl) UpdateBlog(ctx context.Context, req *UpdateBlogRequest) error {
	if err := validate.Struct(req); err != nil {
		return err
	}

	blog, err := s.QueryBlogById(ctx, NewQueryBlogByIdRequest(req.Id))
	if err != nil {
		return err
	}

	err = utils.Merge(blog.CreateBlogRequest, req, mergo.WithOverride)
	if err != nil {
		return err
	}

	tx := s.db.WithContext(ctx).Model(&Blog{}).Where("id = ?", req.Id).Updates(blog)
	if err = tx.Error; err != nil {
		return err
	}
	if affected := tx.RowsAffected; affected == 0 {
		return UpdateFailed
	}

	return nil
}

func (s *ServiceImpl) UpdateBlogStatus(ctx context.Context, req *UpdateBlogStatusRequest) error {
	if err := validate.Struct(req); err != nil {
		return err
	}

	blog, err := s.QueryBlogById(ctx, NewQueryBlogByIdRequest(req.Id))
	if err != nil {
		return err
	}

	err = s.db.WithContext(ctx).Model(blog).Update("status", req.Status).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceImpl) QueryBlog(ctx context.Context, req *QueryBlogRequest) (*Blogs, error) {
	if err := validate.Struct(req); err != nil {
		return nil, err
	}

	blogs := NewBlogs()
	query := s.db.WithContext(ctx).Model(&Blog{})
	// 根据请求参数组装查询条件
	if req.Status != nil {
		query = query.Where("status = ?", req.Status)
	}

	// 查询总数
	err := query.Count(&blogs.Total).Error
	if err != nil {
		return nil, err
	}

	// 查询分页
	err = query.Offset(req.Offset()).Limit(req.PageSize).Find(&blogs.Items).Error
	if err != nil {
		return nil, err
	}

	return blogs, nil
}

func (s *ServiceImpl) QueryBlogById(ctx context.Context, req *QueryBlogByIdRequest) (*Blog, error) {
	if err := validate.Struct(req); err != nil {
		return nil, err
	}

	var blog *Blog
	// 使用Find不会排序
	tx := s.db.WithContext(ctx).Where("id = ?", req.Id).Find(&blog)
	if affected := tx.RowsAffected; affected == 0 {
		return nil, NotFound
	}
	if err := tx.Error; err != nil {
		return nil, err
	}

	// 只允许用户增删改查自己的内容, 否则返回403权限拒绝
	tk, err := s.token.GetTokenByContext(ctx)
	if err != nil {
		return nil, err
	}
	if tk.Username != blog.CreateBy {
		return nil, PermissionDenied
	}

	return blog, nil
}

func (s *ServiceImpl) AuditBlog(ctx context.Context, req *AuditBlogRequest) error {
	if err := validate.Struct(req); err != nil {
		return err
	}

	tx := s.db.WithContext(ctx).Model(&Blog{}).Where("id = ?", req.Id).
		Updates(Blog{AuditAt: time.Now().Unix(), AuditStatus: req.AuditStatus})
	if err := tx.Error; err != nil {
		return err
	}
	if affected := tx.RowsAffected; affected == 0 {
		return UpdateFailed
	}

	return nil
}
