package token

import (
	"context"
	"github.com/luyasr/simple-blog/config"
	"github.com/luyasr/simple-blog/pkg/e"
	"github.com/luyasr/simple-blog/pkg/user"
	"github.com/luyasr/simple-blog/pkg/utils"
	"github.com/luyasr/simple-blog/pkg/validate"
	"gorm.io/gorm"
	"time"
)

var (
	_ Service = (*ServiceImpl)(nil)
)

type ServiceImpl struct {
	db   *gorm.DB
	user *user.ServiceImpl
}

func NewServiceImpl(userService *user.ServiceImpl) *ServiceImpl {
	return &ServiceImpl{
		db:   config.C.Mysql.GetConn(),
		user: userService,
	}
}

func (s *ServiceImpl) Login(ctx context.Context, req *LoginRequest) (*Token, error) {
	// 校验登录请求
	if err := validate.Struct(req); err != nil {
		return nil, err
	}

	// 查询用户信息
	byUsername := user.NewDescribeUserRequestByUsername(req.Username)
	u, err := s.user.DescribeUser(ctx, byUsername)
	if err != nil {
		return nil, e.NewAuthFailed("用户名或密码错误")
	}

	// 校验用户密码是否正确
	if err = utils.PasswordCompare(u.Password, req.Password); err != nil {
		return nil, e.NewAuthFailed("用户名或密码错误")
	}

	// 颁发token
	token := NewToken()
	token.UserID = u.ID
	token.Username = u.Username

	// 限制多地登录, 登录前删除现有token
	if err = s.db.WithContext(ctx).Where("username = ?", token.Username).Delete(token).Error; err != nil {
		return nil, err
	}

	// token入库
	if err := s.db.WithContext(ctx).Create(token).Error; err != nil {
		return nil, err
	}

	return token, nil
}

func (s *ServiceImpl) Logout(ctx context.Context, req *LogoutRequest) error {
	return nil
}

func (s *ServiceImpl) Validate(ctx context.Context, req *ValidateToken) error {
	token := NewDefaultToken()
	// 校验token请求
	if err := validate.Struct(req); err != nil {
		return err
	}

	// 查询token
	if err := s.db.WithContext(ctx).Where("access_token = ?", req.AccessToken).First(token).Error; err != nil {
		return e.NewAuthFailed("无效的token")
	}

	// 校验token是否过期
	accessTokenExpiredTime := token.AccessTokenExpiredTime()
	refreshTokenExpiredTime := token.RefreshTokenExpiredTime()
	if time.Since(accessTokenExpiredTime) > 0 {
		if time.Since(refreshTokenExpiredTime) < 0 {
			token.Refresh()
			// 刷新token, 重新入库
			if err := s.db.WithContext(ctx).Model(token).Where("user_id = ?", token.UserID).Update("access_token", token.AccessToken).Error; err != nil {
				return err
			}
		} else {
			return e.NewAuthFailed("token过期,请重新登录")
		}
	}

	return nil
}
