package user

import (
	"context"
	"github.com/luyasr/simple-blog/pkg/utils"
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
	Username string `json:"username" validate:"required,min=3,max=20" label:"用户名"`
	Password string `json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role     Role   `json:"role"`
}

func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{
		Role: RoleMember,
	}
}

// DeleteUserRequest 删除用户的请求
type DeleteUserRequest struct {
	Id int64 `json:"id"`
}

func NewDeleteUserRequest(id string) *DeleteUserRequest {
	return &DeleteUserRequest{
		Id: utils.StringToInt64(id),
	}
}

type UpdateUserRequest struct {
	ID       int64  `json:"id"`
	Username string `json:"username" validate:"omitempty,min=3,max=20" label:"用户名"`
	Password string `json:"password" validate:"omitempty,min=6,max=20" label:"密码"`
	Role     Role   `json:"role"`
}

func NewUpdateUser() *UpdateUserRequest {
	return &UpdateUserRequest{}
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
