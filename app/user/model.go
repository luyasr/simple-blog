package user

import (
	"encoding/json"
	"github.com/luyasr/simple-blog/pkg/logger"
	"github.com/luyasr/simple-blog/pkg/utils"
	"gorm.io/plugin/soft_delete"
)

type User struct {
	Id       int64                 `json:"id"`
	CreateAt int64                 `json:"create_at" gorm:"autoCreateTime"`
	UpdateAt int64                 `json:"update_at" gorm:"autoUpdateTime"`
	DeleteAt soft_delete.DeletedAt `json:"delete_at"`
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
		CreateUserRequest: &CreateUserRequest{},
	}
}

func NewUser(req *CreateUserRequest) *User {
	req.Password = utils.PasswordHash(req.Password)

	return &User{
		CreateUserRequest: req,
	}
}
