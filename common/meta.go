package common

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type Meta struct {
	Id       int64                 `json:"id"`
	CreateAt int64                 `json:"create_at"`
	UpdateAt int64                 `json:"update_at"`
	DeleteAt soft_delete.DeletedAt `json:"delete_at"`
}

func NewMeta() *Meta {
	return &Meta{
		CreateAt: time.Now().Unix(),
	}
}
