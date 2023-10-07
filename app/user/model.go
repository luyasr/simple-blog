package user

import (
	"encoding/json"
	"github.com/luyasr/simple-blog/common"
	"github.com/luyasr/simple-blog/pkg/logger"
	"github.com/luyasr/simple-blog/pkg/utils"
)

type User struct {
	*common.Meta
	*CreateUserRequest
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) String() string {
	bytes, err := json.Marshal(u)
	if err != nil {
		logger.Console.Error().Err(err).Stack().Send()
	}
	return string(bytes)
}

func NewDefaultUser() *User {
	return &User{
		Meta:              &common.Meta{},
		CreateUserRequest: &CreateUserRequest{},
	}
}

func NewUser(req *CreateUserRequest) *User {
	req.Password = utils.PasswordHash(req.Password)

	return &User{
		Meta:              &common.Meta{},
		CreateUserRequest: req,
	}
}
