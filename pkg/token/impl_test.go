package token

import (
	"context"
	"fmt"
	"github.com/luyasr/simple-blog/pkg/user"
	"testing"
)

var (
	tokenSvc = NewServiceImpl(user.NewServiceImpl())
	ctx      = context.Background()
)

func TestHandler_Login(t *testing.T) {
	req := NewLoginRequest()
	req.Username = "admin"
	req.Password = "123456"
	byUsername := user.NewDescribeUserRequestByUsername(req.Username)
	fmt.Println(byUsername)
	_, err := tokenSvc.user.DescribeUser(ctx, byUsername)
	if err != nil {
		t.Fatal(err)
	}
	token, _ := tokenSvc.Login(ctx, req)
	t.Log(token)
}
