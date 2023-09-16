package impl

import (
	"context"
	"github.com/luyasr/simple-blog/config"
	"github.com/luyasr/simple-blog/pkg/user"
	"gorm.io/gorm"
)

var _ user.Service = (*UserServiceImpl)(nil)

type UserServiceImpl struct {
	db *gorm.DB
}

func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{
		db: config.C.Mysql.GetConn(),
	}
}

func (i *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.User, error) {
	ins := user.NewUser(req)
	if err := i.db.WithContext(ctx).Create(ins).Error; err != nil {
		return nil, err
	}
	return ins, nil
}

func (i *UserServiceImpl) DeleteUser(context.Context, *user.DeleteUserRequest) error {
	return nil
}

func (i *UserServiceImpl) DescribeUser(context.Context, *user.DescribeUserRequest) (*user.User, error) {
	return nil, nil
}
