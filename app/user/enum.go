package user

// Role 用户角色
type Role int

const (
	RoleUnKnown Role = iota
	RoleMember
	RoleAdmin
)

// QueryBy 用户查询条件
type QueryBy int

const (
	QueryUnKnown QueryBy = iota
	QueryById
	QueryByUsername
)
