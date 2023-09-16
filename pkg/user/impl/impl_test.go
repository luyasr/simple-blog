package impl_test

import (
	"context"
	"github.com/luyasr/simple-blog/pkg/user"
	"github.com/luyasr/simple-blog/pkg/user/impl"
	"testing"
)

var (
	usrSvc *impl.UserServiceImpl
	ctx    = context.Background()
)

func init() {
	usrSvc = impl.NewUserServiceImpl()
}

func TestUserServiceImpl_CreateUser(t *testing.T) {
	req := user.NewCreateUserRequest()
	req.Username = "admin"
	req.Password = "12345"
	u, err := usrSvc.CreateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}
