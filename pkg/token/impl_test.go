package token

import (
	"context"
	"github.com/luyasr/simple-blog/pkg/user"
	"testing"
)

var (
	tokenSvc *ServiceImpl
	ctx      = context.Background()
)

func TestHandler_Login(t *testing.T) {
	req := NewLoginRequest()
	req.Username = "admin"
	req.Password = "123456"
	byUsername := user.NewDescribeUserRequestByUsername(req.Username)
	_, err := tokenSvc.user.DescribeUser(ctx, byUsername)
	if err != nil {
		t.Fatal(err)
	}
	token, _ := tokenSvc.Login(ctx, req)
	t.Log(token)
}
