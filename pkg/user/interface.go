package user

import (
	"context"
	"golang.org/x/crypto/bcrypt"
)

const (
	RoleMember Role = iota
	RoleAdmin
)

// Role 用户角色
type Role int

// Service 接口定义
type Service interface {
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	DeleteUser(context.Context, *DeleteUserRequest) error
	DescribeUser(context.Context, *DescribeUserRequest) (*User, error)
}

// CreateUserRequest 创建用户的请求
type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
	isHashed bool
}

func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{
		Role: RoleMember,
	}
}

func (req *CreateUserRequest) SetIsHashed() {
	req.isHashed = true
}

func (req *CreateUserRequest) PasswordHash() {
	if req.isHashed {
		return
	}

	bytes, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	req.Password = string(bytes)
	req.isHashed = true
}

// DeleteUserRequest 删除用户的请求
type DeleteUserRequest struct {
	Id int64 `json:"id"`
}

// DescribeUserRequest 查看用户的请求
type DescribeUserRequest struct {
}
