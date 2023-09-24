package token

import (
	"context"
	"github.com/luyasr/simple-blog/pkg/ioc"
	"testing"
)

var (
	tokenSvc Service
	ctx      = context.Background()
)

func init() {
	_ = ioc.Controller().Init()
	tokenSvc = ioc.Controller().Get(Name).(Service)
}

func TestHandler_Login(t *testing.T) {
	req := NewLoginRequest()
	req.Username = "admin2"
	req.Password = "123456"
	token, _ := tokenSvc.Login(ctx, req)
	t.Log(token)
}

func TestHandler_Logout(t *testing.T) {
	req := NewLogoutRequest()
	req.AccessToken = "ck6n4g9qd2tum4iatjkg"
	req.RefreshToken = "ck6n4g9qd2tum4iatjl0"
	err := tokenSvc.Logout(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
}
