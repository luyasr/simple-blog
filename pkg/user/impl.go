package user

import (
	"context"
	"errors"
	"github.com/luyasr/simple-blog/config"
	"github.com/luyasr/simple-blog/pkg/e"
	"github.com/luyasr/simple-blog/pkg/utils"
	"gorm.io/gorm"
)

var _ Service = (*UserServiceImpl)(nil)

type UserServiceImpl struct {
	db *gorm.DB
}

func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{
		db: config.C.Mysql.GetConn(),
	}
}

func (i *UserServiceImpl) CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error) {
	ins := NewUser(req)
	err := i.db.WithContext(ctx).First(&ins, "username = ?", ins.Username).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, e.NewUserExists("用户 %s 已存在", ins.Username)
	}

	if err := i.db.WithContext(ctx).Create(ins).Error; err != nil {
		return nil, err
	}
	return ins, nil
}

func (i *UserServiceImpl) DeleteUser(ctx context.Context, req *DeleteUserRequest) error {
	ins := NewUser(NewCreateUserRequest())
	ins.Id = req.Id
	if affected := i.db.WithContext(ctx).Delete(ins).RowsAffected; affected == 0 {
		return e.NewUserNotExists("用户ID %d 已注销或不存在", ins.Id)
	}
	return nil
}

func (i *UserServiceImpl) UpdateUser(ctx context.Context, req *UpdateUserRequest) error {
	// 创建用户实例更新
	ins := NewUpdateUserRequest(req)
	// 更新多列
	fields, err := utils.UpdateNonZeroFields(req)
	if err != nil {
		return err
	}

	result := i.db.WithContext(ctx).Model(ins).Updates(fields)
	if err = result.Error; err != nil {
		return err
	}
	affected := result.RowsAffected
	if affected == 0 {
		return e.NewUpdateFailed("用户ID %d RowsAffected: 0", ins.Id)
	}

	return nil
}

func (i *UserServiceImpl) DescribeUser(ctx context.Context, req *DescribeUserRequest) (*User, error) {
	ins := NewDefaultUser()
	query := i.db.WithContext(ctx)

	switch req.DescribeBy {
	case DescribeById:
		query = query.Where("id = ?", req.DescribeValue)
	case DescribeByUsername:
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
