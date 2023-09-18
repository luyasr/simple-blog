package user

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
