package user

import (
	"context"
)

const Name = "user"

// Service 接口定义
type Service interface {
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	DeleteUser(context.Context, *DeleteUserRequest) error
	UpdateUser(context.Context, *UpdateUserRequest) error
	QueryUser(context.Context, *QueryUserRequest) (*User, error)
}

// CreateUserRequest 创建用户的请求
type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=20" label:"用户名"`
	Password string `json:"password" validate:"required,min=6,max=20" label:"密码"`
}

func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{}
}

// DeleteUserRequest 删除用户的请求
type DeleteUserRequest struct {
	Id string `json:"id"`
}

func NewDeleteUserRequest(id string) *DeleteUserRequest {
	return &DeleteUserRequest{
		Id: id,
	}
}

type UpdateUserRequest struct {
	Id       string `json:"id" validate:"required" label:"用户id"`
	Username string `json:"username" validate:"omitempty,min=3,max=20" label:"用户名"`
	Password string `json:"password" validate:"omitempty,min=6,max=20" label:"密码"`
}

func NewUpdateUser() *UpdateUserRequest {
	return &UpdateUserRequest{}
}

// QueryUserRequest 查看用户的请求
type QueryUserRequest struct {
	QueryBy    QueryBy `json:"query_by"`
	QueryValue string  `json:"query_value"`
}

func NewQueryUserRequestById(id string) *QueryUserRequest {
	return &QueryUserRequest{
		QueryBy:    QueryById,
		QueryValue: id,
	}
}

func NewQueryUserRequestByUsername(username string) *QueryUserRequest {
	return &QueryUserRequest{
		QueryBy:    QueryByUsername,
		QueryValue: username,
	}
}
