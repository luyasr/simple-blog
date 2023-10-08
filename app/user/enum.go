package user

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
