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
	req := NewLogoutOrRefreshRequest()
	req.UserId = 17
	req.AccessToken = "ckv3ql9qd2tj1lnocimg"
	req.RefreshToken = "ckv3ql9qd2tj1lnocin0"
	err := tokenSvc.Logout(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestServiceImpl_Refresh(t *testing.T) {
	req := NewLogoutOrRefreshRequest()
	req.AccessToken = "ckkd7as2a0g97ccldsug"
	req.RefreshToken = "ckkd7as2a0g97ccldsv0"
	tk, err := tokenSvc.Refresh(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}
