package token

import (
	"context"
	"errors"
	"github.com/luyasr/simple-blog/app/user"
	"github.com/luyasr/simple-blog/config"
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
		return nil, err
	}

	// 校验用户密码是否正确
	if err = utils.PasswordCompare(u.Password, req.Password); err != nil {
		return nil, AuthFailed
	}

	// 颁发token
	token := NewToken()
	token.UserId = u.Id
	token.Username = u.Username

	// 用户存在token登录更新, 不存在token登录创建
	query, err := s.Query(ctx, NewQueryTokenByUserIdRequest(utils.Int64ToString(token.UserId)))
	if err != nil {
		if !errors.Is(err, NotFound) {
			return nil, err
		}
	}
	if query != nil {
		if err = s.db.WithContext(ctx).Model(&token).Where("user_id", token.UserId).Updates(token).Error; err != nil {
			return nil, err
		}
	} else {
		if err = s.db.WithContext(ctx).Create(token).Error; err != nil {
			return nil, err
		}
	}

	return token, err
}

func (s *ServiceImpl) Logout(ctx context.Context, req *LogoutOrRefreshRequest) error {
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
			return NotFound
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

func (s *ServiceImpl) Query(ctx context.Context, req *QueryTokenRequest) (*Token, error) {
	// 校验查询请求
	if err := validate.Struct(req); err != nil {
		return nil, err
	}

	var token *Token
	query := s.db.WithContext(ctx)

	switch req.QueryBy {
	case QueryByUserId:
		query = query.Where("user_id = ?", req.QueryByValue)
	case QueryByAccessToken:
		query = query.Where("access_token = ?", req.QueryByValue)
	}

	if err := query.First(&token).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, NotFound
		}
		return nil, err
	}

	return token, nil
}

func (s *ServiceImpl) Refresh(ctx context.Context, req *LogoutOrRefreshRequest) (*Token, error) {
	// 校验刷新请求
	if err := validate.Struct(req); err != nil {
		return nil, err
	}

	var token *Token
	err := s.db.WithContext(ctx).Where("access_token = ? AND refresh_token = ?", req.AccessToken, req.RefreshToken).
		First(&token).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, NotFound
		}
		return nil, err
	}

	token.Refresh()
	err = s.db.WithContext(ctx).Model(token).
		Where("user_id = ?", token.UserId).
		Update("access_token", token.AccessToken).Error
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *ServiceImpl) Validate(ctx context.Context, req *ValidateToken) (*Token, error) {
	// 校验token请求
	if err := validate.Struct(req); err != nil {
		return nil, err
	}

	var token *Token
	// 查询token
	if err := s.db.WithContext(ctx).Where("access_token = ?", req.AccessToken).First(&token).Error; err != nil {
		return nil, InvalidToken
	}

	// 校验token是否过期
	accessTokenExpiredTime := token.AccessTokenExpiredTime()
	refreshTokenExpiredTime := token.RefreshTokenExpiredTime()
	if time.Since(accessTokenExpiredTime) > 0 && time.Since(refreshTokenExpiredTime) > 0 {
		return nil, ExpiresToken
	}

	return token, nil
}
