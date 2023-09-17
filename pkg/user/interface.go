package user

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

// Role 用户角色
type Role int

const (
	RoleMember Role = iota + 1
	RoleAdmin
)

// DescribeBy 用户查询条件
type DescribeBy int

const (
	DescribeById = iota + 1
	DescribeByUsername
)

// Service 接口定义
type Service interface {
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	DeleteUser(context.Context, *DeleteUserRequest) error
	UpdateUser(context.Context, *UpdateUserRequest) (*User, error)
	DescribeUser(context.Context, *DescribeUserRequest) (*User, error)
}

// CreateUserRequest 创建用户的请求
type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
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

type UpdateUserRequest struct {
	*User
}

// DescribeUserRequest 查看用户的请求
type DescribeUserRequest struct {
	DescribeBy    DescribeBy `json:"describe_by"`
	DescribeValue string     `json:"describe_value"`
}

func NewDescribeUserRequestById(id string) *DescribeUserRequest {
	return &DescribeUserRequest{
		DescribeValue: id,
	}
}

func NewDescribeUserRequestByUsername(username string) *DescribeUserRequest {
	return &DescribeUserRequest{
		DescribeBy:    DescribeByUsername,
		DescribeValue: username,
	}
}
