package impl

import (
	"context"
	"errors"
	"github.com/luyasr/simple-blog/config"
	"github.com/luyasr/simple-blog/pkg/e"
	"github.com/luyasr/simple-blog/pkg/user"
	"github.com/luyasr/simple-blog/pkg/util"
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
	err := i.db.WithContext(ctx).First(&ins, "username = ?", ins.Username).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, e.NewUserExists("用户 %s 已存在", ins.Username)
	}

	if err := i.db.WithContext(ctx).Create(ins).Error; err != nil {
		return nil, err
	}
	return ins, nil
}

func (i *UserServiceImpl) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) error {
	ins := user.NewUser(user.NewCreateUserRequest())
	ins.Id = req.Id
	if affected := i.db.WithContext(ctx).Delete(ins).RowsAffected; affected == 0 {
		return e.NewUserNotExists("用户ID %d 不存在", ins.Id)
	}
	return nil
}

func (i *UserServiceImpl) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.User, error) {
	ins := user.NewUser(user.NewCreateUserRequest())
	ins.Id = req.Id
	fields, err := util.UpdateNonZeroFields(req)
	if err != nil {
		return nil, err
	}
	if affected := i.db.WithContext(ctx).Model(ins).Updates(fields).RowsAffected; affected == 0 {
		return nil, e.NewUpdateFailed("用户 %d 更新失败", ins.Id)
	}
	return ins, nil
}

func (i *UserServiceImpl) DescribeUser(ctx context.Context, req *user.DescribeUserRequest) (*user.User, error) {
	ins := user.NewUser(user.NewCreateUserRequest())
	query := i.db.WithContext(ctx)

	switch req.DescribeBy {
	case user.DescribeById:
		query = query.Where("id = ?", req.DescribeValue)
	case user.DescribeByUsername:
		query = query.Where("username = ?", req.DescribeValue)
	}

	if err := query.First(ins).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.NewUserNotExists("用户 %v 不存在", req.DescribeValue)
		}
		return nil, err
	}

	return ins, nil
}
