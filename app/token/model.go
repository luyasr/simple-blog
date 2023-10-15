package token

import (
	"encoding/json"
	"github.com/luyasr/simple-blog/app/common"
	"github.com/luyasr/simple-blog/app/user"
	"github.com/luyasr/simple-blog/pkg/logger"
	"github.com/rs/xid"
	"time"
)

type Token struct {
	*common.Meta
	UserId                int64     `json:"user_id" validate:"required"`
	Username              string    `json:"username" validate:"required"`
	AccessToken           string    `json:"access_token" validate:"required"`
	AccessTokenExpiredAt  int       `json:"access_token_expired_at" validate:"required"`
	RefreshToken          string    `json:"refresh_token" validate:"required"`
	RefreshTokenExpiredAt int       `json:"refresh_token_expired_at" validate:"required"`
	Role                  user.Role `gorm:"-"`
}

func (t *Token) TableName() string {
	return "tokens"
}

func (t *Token) String() string {
	bytes, err := json.Marshal(t)
	if err != nil {
		logger.Console.Error().Err(err).Stack().Send()
	}
	return string(bytes)
}

func (t *Token) Refresh() string {
	return xid.New().String()
}

func (t *Token) AccessTokenExpiredTime() time.Time {
	return time.Unix(t.UpdateAt, 0).Add(time.Duration(t.AccessTokenExpiredAt) * time.Second)
}

func (t *Token) RefreshTokenExpiredTime() time.Time {
	return time.Unix(t.CreateAt, 0).Add(time.Duration(t.RefreshTokenExpiredAt) * time.Second)
}

func NewToken() *Token {
	return &Token{
		AccessToken:           xid.New().String(),
		AccessTokenExpiredAt:  7200,
		RefreshToken:          xid.New().String(),
		RefreshTokenExpiredAt: 604800,
	}
}
