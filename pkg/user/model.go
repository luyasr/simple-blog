package user

import (
	"encoding/json"
	"github.com/luyasr/simple-blog/common"
	"github.com/luyasr/simple-blog/pkg/logger"
	"strconv"
)

type User struct {
	*common.Meta
	*CreateUserRequest
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) String() string {
	j, err := json.Marshal(u)
	if err != nil {
		logger.L.Err(err)
	}

	return string(j)
}

func IdInt64(id string) int64 {
	idInt64, _ := strconv.ParseInt(id, 10, 64)
	return idInt64
}

func NewDefaultUser() *User {
	return &User{
		&common.Meta{},
		&CreateUserRequest{},
	}
}

func NewUser(req *CreateUserRequest) *User {
	req.PasswordHash()

	return &User{
		Meta:              common.NewMeta(),
		CreateUserRequest: req,
	}
}
