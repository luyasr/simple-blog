package user

import (
	"context"
	"dario.cat/mergo"
	"errors"
	"github.com/luyasr/simple-blog/config"
	"github.com/luyasr/simple-blog/pkg/ioc"
	"github.com/luyasr/simple-blog/pkg/utils"
	"github.com/luyasr/simple-blog/pkg/validate"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

var (
	_ Service = (*ServiceImpl)(nil)
)

// init ioc controller容器注册
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

func (s *ServiceImpl) CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error) {
	// 校验CreateUserRequest字段
	if err := validate.Struct(req); err != nil {
		return nil, err
	}

	user := NewUser(req)
	queryUser, _ := s.QueryUser(ctx, NewQueryUserRequestByUsername(user.Username))
	if queryUser != nil {
		return nil, AlreadyExist
	}
	// 用户默认角色
	user.Role = RoleAuthor

	if err := s.db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *ServiceImpl) DeleteUser(ctx context.Context, req *DeleteUserRequest) error {
	// 校验DeleteUserRequest请求
	if err := validate.Struct(req); err != nil {
		return err
	}

	// 删除前, 先查询是否存在
	user, err := s.QueryUser(ctx, NewQueryUserRequestById(req.Id))
	if err != nil {
		return err
	}

	if err := s.db.WithContext(ctx).Delete(user).Error; err != nil {
		return err
	}
	return nil
}

func (s *ServiceImpl) UpdateUser(ctx context.Context, req *UpdateUserRequest) error {
	// 校验UpdateUserRequest字段
	if err := validate.Struct(req); err != nil {
		return err
	}
	// 查询用户
	user, err := s.QueryUser(ctx, NewQueryUserRequestById(req.Id))
	if err != nil {
		return err
	}

	// PasswordHash 如果用户密码相同, 密码置空, 否则hash
	if err := utils.PasswordCompare(user.Password, req.Password); err == nil {
		req.Password = ""
	} else {
		req.Password = utils.PasswordHash(req.Password)
	}

	// 合并结构体
	err = utils.Merge(user.CreateUserRequest, req, mergo.WithOverride)
	if err != nil {
		return err
	}

	tx := s.db.WithContext(ctx).Model(user).Updates(&user)
	if err := tx.Error; err != nil {
		return err
	}
	if affected := tx.RowsAffected; affected == 0 {
		return UpdateFailed
	}

	return nil
}

func (s *ServiceImpl) QueryUser(ctx context.Context, req *QueryUserRequest) (*User, error) {
	var user *User
	query := s.db.WithContext(ctx)

	switch req.QueryBy {
	case QueryById:
		query = query.Where("id = ?", req.QueryValue)
	case QueryByUsername:
		query = query.Where("username = ?", req.QueryValue)
	}

	if err := query.First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, NotFound
		}
		return nil, err
	}

	return user, nil
}
