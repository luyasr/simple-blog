package user

import (
	"context"
	"errors"
	"github.com/luyasr/simple-blog/config"
	"github.com/luyasr/simple-blog/pkg/e"
	"github.com/luyasr/simple-blog/pkg/utils"
	"github.com/luyasr/simple-blog/pkg/validate"
	"gorm.io/gorm"
)

var (
	_ Service = (*UserServiceImpl)(nil)
)

type UserServiceImpl struct {
	db *gorm.DB
}

func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{
		db: config.C.Mysql.GetConn(),
	}
}

func (i *UserServiceImpl) CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error) {
	// 先进行字段参数验证
	if err := validate.Struct(req); err != nil {
		return nil, err
	}

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
	ins := NewDefaultUser()
	ins.ID = req.Id
	if affected := i.db.WithContext(ctx).Delete(ins).RowsAffected; affected == 0 {
		return e.NewUserNotExists("用户ID %d 不存在", ins.ID)
	}
	return nil
}

func (i *UserServiceImpl) UpdateUser(ctx context.Context, req *UpdateUserRequest) error {
	// 校验UpdateUserRequest字段
	if err := validate.Struct(req); err != nil {
		return err
	}
	// 创建用户实例更新
	ins := NewDefaultUser()
	// 查询用户
	if err := i.db.WithContext(ctx).Where("id = ?", req.ID).First(ins).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return e.NewUserNotExists("用户ID %d 不存在", req.ID)
		}
		return err
	}
	// 需要更新的字段 更新多列
	fields, err := newUpdateFields(ins, req)
	if err != nil {
		return err
	}

	result := i.db.WithContext(ctx).Model(ins).Updates(fields)
	if err = result.Error; err != nil {
		return err
	}
	affected := result.RowsAffected
	if affected == 0 {
		return e.NewUpdateFailed("用户ID %d 更新失败, 受影响的行记录 %d", ins.ID, affected)
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

func newUpdateFields(u *User, req *UpdateUserRequest) (map[string]any, error) {
	// User struct to map
	um, err := utils.StructToMap(u)
	if err != nil {
		return nil, err
	}

	// PasswordHash 如果用户密码相同, 密码置空, 否则hash
	if err := utils.PasswordCompare(u.Password, req.Password); err == nil {
		req.Password = ""
	} else {
		req.Password = utils.PasswordHash(req.Password)
	}

	// 更新多列 获取非零字段
	fields, err := utils.UpdateNonZeroFields(req)
	if err != nil {
		return nil, err
	}

	// 用户更新请求字段和当前记录字段一致不更新
	for field, _ := range fields {
		if um[field] == fields[field] {
			delete(fields, field)
		}
	}
	if len(fields) == 0 {
		return nil, e.NewUpdateFailed("当前记录未发生改变")
	}

	return fields, nil
}
