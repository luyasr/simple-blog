package token

import (
	"context"
	"errors"
	"github.com/luyasr/simple-blog/app/user"
	"github.com/luyasr/simple-blog/config"
	"github.com/luyasr/simple-blog/pkg/e"
	"github.com/luyasr/simple-blog/pkg/ioc"
	"github.com/luyasr/simple-blog/pkg/utils"
	"github.com/luyasr/simple-blog/pkg/validate"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"time"
)

var (
	_ Service = (*ServiceImpl)(nil)
)

func init() {
	ioc.Controller().Registry(&ServiceImpl{})
}

func (s *ServiceImpl) Init() error {
	s.db = config.C.Mysql.GetConn()
	s.user = ioc.Controller().Get(user.Name).(user.Service)
	return nil
}

func (s *ServiceImpl) Name() string {
	return Name
}

type ServiceImpl struct {
	db   *gorm.DB
	user user.Service
	log  zerolog.Logger
}

func (s *ServiceImpl) Login(ctx context.Context, req *LoginRequest) (*Token, error) {
	// 校验登录请求
	if err := validate.Struct(req); err != nil {
		return nil, err
	}

	// 查询用户信息
	byUsername := user.NewQueryUserRequestByUsername(req.Username)
	u, err := s.user.QueryUser(ctx, byUsername)
	if err != nil {
		return nil, e.NewAuthFailed("用户名或密码错误")
	}

	// 校验用户密码是否正确
	if err = utils.PasswordCompare(u.Password, req.Password); err != nil {
		return nil, e.NewAuthFailed("用户名或密码错误")
	}

	// 颁发token
	token := NewToken()
	token.UserID = u.Id
	token.Username = u.Username

	// 用户存在token登录更新, 不存在token登录创建
	if err = s.db.WithContext(ctx).Where("user_id = ?", token.UserID).First(&Token{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err = s.db.WithContext(ctx).Create(token).Error; err != nil {
				return nil, err
			}
		}
		return nil, err
	} else {
		if err = s.db.WithContext(ctx).Model(token).Where("user_id", token.UserID).Updates(token).Error; err != nil {
			return nil, err
		}
	}
	return token, err
}

func (s *ServiceImpl) Logout(ctx context.Context, req *LogoutRequest) error {
	// 校验退出请求
	if err := validate.Struct(req); err != nil {
		return err
	}

	var token *Token
	err := s.db.WithContext(ctx).Where("access_token = ? AND refresh_token = ?",
		req.AccessToken,
		req.RefreshToken,
	).First(&token).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return e.NewNotFound("access_token %s没找到", req.AccessToken)
		}
		return err
	}

	// 删除token
	err = s.db.WithContext(ctx).Where("access_token = ? AND refresh_token = ?",
		token.AccessToken,
		token.RefreshToken,
	).Delete(&token).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceImpl) Validate(ctx context.Context, req *ValidateToken) error {
	// 校验token请求
	if err := validate.Struct(req); err != nil {
		return err
	}

	var token *Token
	// 查询token
	if err := s.db.WithContext(ctx).Where("access_token = ?", req.AccessToken).First(&token).Error; err != nil {
		return e.NewAuthFailed("无效的token")
	}

	// 校验token是否过期
	accessTokenExpiredTime := token.AccessTokenExpiredTime()
	refreshTokenExpiredTime := token.RefreshTokenExpiredTime()
	if time.Since(accessTokenExpiredTime) > 0 {
		if time.Since(refreshTokenExpiredTime) < 0 {
			// 后端刷新token
			token.Refresh()

			// token更新
			err := s.db.WithContext(ctx).Model(token).Update("access_token", token.AccessToken).Error
			if err != nil {
				return err
			}
		} else {
			return e.NewAuthFailed("token过期,请重新登录")
		}
	}

	return nil
}
