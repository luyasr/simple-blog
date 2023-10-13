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
	req.AccessToken = "ckkcnuk2a0g395d00610"
	req.RefreshToken = "ckkcnuk2a0g395d0061g"
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
