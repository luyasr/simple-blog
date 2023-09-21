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
