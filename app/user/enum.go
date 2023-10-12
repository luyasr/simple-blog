package user

import "github.com/luyasr/simple-blog/pkg/e"

// Role 用户角色
type Role int

const (
	RoleMember Role = iota
	RoleAdmin
)

// QueryBy 用户查询条件
type QueryBy int

const (
	QueryById QueryBy = iota
	QueryByUsername
)

var (
	NotFound     = e.NewNotFound("user not found")
	AlreadyExist = e.NewAlreadyExist("user already exists")
	UpdateFailed = e.NewUpdateFailed("user update failed")
)
