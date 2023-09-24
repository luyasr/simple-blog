package user

import (
	"context"
	"errors"
	"github.com/luyasr/simple-blog/config"
	"github.com/luyasr/simple-blog/pkg/e"
	"github.com/luyasr/simple-blog/pkg/ioc"
	"github.com/luyasr/simple-blog/pkg/utils"
	"github.com/luyasr/simple-blog/pkg/validate"
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
	db *gorm.DB
}

func (s *ServiceImpl) CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error) {
	// 校验CreateUserRequest字段
	if err := validate.Struct(req); err != nil {
		return nil, err
	}

	ins := NewUser(req)
	err := s.db.WithContext(ctx).First(&ins, "username = ?", ins.Username).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, e.NewFound("用户%s已存在", ins.Username)
	}

	if err := s.db.WithContext(ctx).Create(ins).Error; err != nil {
		return nil, err
	}
	return ins, nil
}

func (s *ServiceImpl) DeleteUser(ctx context.Context, req *DeleteUserRequest) error {
	// 校验DeleteUserRequest请求
	if err := validate.Struct(req); err != nil {
		return err
	}
	// 删除前, 先查询是否存在
	user, err := s.DescribeUser(ctx, NewDescribeUserRequestById(utils.Int64ToString(req.ID)))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return e.NewNotFound("用户%d没找到", user.ID)
		}
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
	// 创建用户实例更新
	ins := NewDefaultUser()
	// 查询用户
	if err := s.db.WithContext(ctx).Where("id = ?", req.ID).First(ins).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return e.NewNotFound("用户%d没找到", req.ID)
		}
		return err
	}

	// PasswordHash 如果用户密码相同, 密码置空, 否则hash
	if err := utils.PasswordCompare(ins.Password, req.Password); err == nil {
		req.Password = ""
	} else {
		req.Password = utils.PasswordHash(req.Password)
	}

	// 需要更新的字段 更新多列
	fields, err := newUpdateFields(ins, req)
	if err != nil {
		return err
	}

	result := s.db.WithContext(ctx).Model(ins).Updates(fields)
	if err = result.Error; err != nil {
		return err
	}
	affected := result.RowsAffected
	if affected == 0 {
		return e.NewUpdateFailed("用户%d更新失败, 受影响的行记录 %d", ins.ID, affected)
	}

	return nil
}

func (s *ServiceImpl) DescribeUser(ctx context.Context, req *DescribeUserRequest) (*User, error) {
	ins := NewDefaultUser()
	query := s.db.WithContext(ctx)

	switch req.DescribeBy {
	case DescribeById:
		query = query.Where("id = ?", req.DescribeValue)
	case DescribeByUsername:
		query = query.Where("username = ?", req.DescribeValue)
	}

	if err := query.First(ins).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.NewNotFound("用户%v没找到", req.DescribeValue)
		}
		return nil, err
	}

	return ins, nil
}

func newUpdateFields(obj any, updateObj any) (map[string]any, error) {
	// 结构体转map
	m, err := utils.StructToMap(obj)
	if err != nil {
		return nil, err
	}

	// 更新多列 获取非零字段
	fields, err := utils.UpdateNonZeroFields(updateObj)
	if err != nil {
		return nil, err
	}

	// 更新请求字段和当前记录字段一致不更新
	for field := range fields {
		if m[field] == fields[field] {
			delete(fields, field)
		}
	}
	if len(fields) == 0 {
		return nil, e.NewUpdateFailed("当前记录未发生改变")
	}

	return fields, nil
}
