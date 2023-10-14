package common

type Meta struct {
	// 自增主键
	Id int64 `json:"id"`
	// 创建时间
	CreateAt int64 `json:"create_at" gorm:"autoCreateTime"`
	// 更新时间
	UpdateAt int64 `json:"update_at" gorm:"autoUpdateTime"`
}
