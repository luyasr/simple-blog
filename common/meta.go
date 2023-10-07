package common

import (
	"gorm.io/plugin/soft_delete"
)

type Meta struct {
	Id       int64                 `json:"id"`
	CreateAt int64                 `json:"create_at" gorm:"autoCreateTime"`
	UpdateAt int64                 `json:"update_at" gorm:"autoUpdateTime"`
	DeleteAt soft_delete.DeletedAt `json:"delete_at"`
}
