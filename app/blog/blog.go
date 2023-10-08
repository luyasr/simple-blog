package blog

import (
	"context"
	"github.com/luyasr/simple-blog/config"
	"github.com/luyasr/simple-blog/pkg/ioc"
	"github.com/luyasr/simple-blog/pkg/validate"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

var (
	_ Service = (*ServiceImpl)(nil)
)

func init() {
	ioc.Controller().Registry(&ServiceImpl{})
}

func (s *ServiceImpl) Init() error {
	s.db = config.C.Mysql.GetConn()
	return nil
}

func (s *ServiceImpl) Name() string {
	return Name
}

type ServiceImpl struct {
	db  *gorm.DB
	log zerolog.Logger
}

func (s *ServiceImpl) CreateBlog(ctx context.Context, req *CreateBlogRequest) (*Blog, error) {
	if err := validate.Struct(req); err != nil {
		return nil, err
	}

	ins := NewBlog(req)
	err := s.db.WithContext(ctx).Create(ins).Error
	if err != nil {
		return nil, err
	}
	return ins, nil
}
func (s *ServiceImpl) DeleteBlog(ctx context.Context, req *DeleteBlogRequest) error {
	if err := validate.Struct(req); err != nil {
		return err
	}

	var blog Blog
	blog.Id = req.Id
	err := s.db.WithContext(ctx).Delete(blog).Error
	if err != nil {
		return err
	}
	return nil
}
func (s *ServiceImpl) UpdateBlog(ctx context.Context, req *UpdateBlogRequest) error {
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
