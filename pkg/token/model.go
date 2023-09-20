package token

import (
	"encoding/json"
	"github.com/luyasr/simple-blog/pkg/logger"
	"github.com/rs/xid"
)

type Token struct {
	UserID                int64  `json:"user_id" validate:"required"`
	Username              string `json:"username" validate:"required"`
	AccessToken           string `json:"access_token" validate:"required"`
	AccessTokenExpiredAt  int    `json:"access_token_expired_at" validate:"required"`
	RefreshToken          string `json:"refresh_token" validate:"required"`
	RefreshTokenExpiredAt int    `json:"refresh_token_expired_at" validate:"required"`
	CreatedAt             int64  `json:"created_at" gorm:"autoUpdateTime"`
	UpdatedAt             int64  `json:"updated_at" gorm:"autoUpdateTime"`
}

func (t *Token) TableName() string {
	return "tokens"
}

func (t *Token) String() string {
	j, err := json.Marshal(t)
	if err != nil {
		logger.L.Err(err).Send()
	}
	return string(j)
}

func NewToken() *Token {
	return &Token{
		AccessToken:           xid.New().String(),
		AccessTokenExpiredAt:  7200,
		RefreshToken:          xid.New().String(),
		RefreshTokenExpiredAt: 604800,
	}
}
