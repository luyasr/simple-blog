package user

import (
	"context"
	"github.com/luyasr/simple-blog/common"

	"golang.org/x/crypto/bcrypt"
)

// Service 接口定义
type Service interface {
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	DeleteUser(context.Context, *DeleteUserRequest) error
	UpdateUser(context.Context, *UpdateUserRequest) error
	DescribeUser(context.Context, *DescribeUserRequest) (*User, error)
}

// CreateUserRequest 创建用户的请求
type CreateUserRequest struct {
	Username string `json:"username" binding:"required" msg:"用户名校验失败"`
	Password string `json:"password" binding:"required" msg:"密码校验失败"`
	Role     Role   `json:"role"`
}

func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{
		Role: RoleMember,
	}
}

func (req *CreateUserRequest) PasswordHash() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	req.Password = string(bytes)
}

// DeleteUserRequest 删除用户的请求
type DeleteUserRequest struct {
	Id int64 `json:"id"`
}

func NewDeleteUserRequest(id string) *DeleteUserRequest {
	return &DeleteUserRequest{
		Id: IdInt64(id),
	}
}

type UpdateUserRequest = User

func NewUpdateUserRequest(id string) *UpdateUserRequest {
	return &UpdateUserRequest{
		&common.Meta{
			Id: IdInt64(id),
		},
		&CreateUserRequest{},
	}
}

// DescribeUserRequest 查看用户的请求
type DescribeUserRequest struct {
	DescribeBy    DescribeBy `json:"describe_by"`
	DescribeValue string     `json:"describe_value"`
}

func NewDescribeUserRequestById(id string) *DescribeUserRequest {
	return &DescribeUserRequest{
		DescribeBy:    DescribeById,
		DescribeValue: id,
	}
}

func NewDescribeUserRequestByUsername(username string) *DescribeUserRequest {
	return &DescribeUserRequest{
		DescribeBy:    DescribeByUsername,
		DescribeValue: username,
	}
}
