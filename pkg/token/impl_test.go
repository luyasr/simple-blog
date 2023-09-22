package token

import (
	"context"
	"github.com/luyasr/simple-blog/pkg/user"
	"testing"
)

var (
	tokenSvc = NewServiceImpl(user.NewServiceImpl())
	ctx      = context.Background()
)

func TestHandler_Login(t *testing.T) {
	req := NewLoginRequest()
	req.Username = "admin5"
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
