package user

import (
	"encoding/json"
	"github.com/luyasr/simple-blog/app/common"
	"github.com/luyasr/simple-blog/pkg/logger"
	"github.com/luyasr/simple-blog/pkg/utils"
	"gorm.io/plugin/soft_delete"
)

type User struct {
	*common.Meta
	*CreateUserRequest
	// 头像
	Avatar string `json:"avatar"`
	// 角色
	Role Role `json:"role"`
	// 逻辑删除, 有唯一索引的情况下, 需要和唯一索引建立复合索引
	DeleteAt soft_delete.DeletedAt `json:"delete_at"`
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

func NewUser(req *CreateUserRequest) *User {
	req.Password = utils.PasswordHash(req.Password)

	return &User{
		CreateUserRequest: req,
	}
}
