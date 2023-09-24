package user

import (
	"context"
	"github.com/luyasr/simple-blog/pkg/ioc"
	"github.com/luyasr/simple-blog/pkg/utils"
	"testing"
)

var (
	userSvc Service
	ctx     = context.Background()
)

func init() {
	_ = ioc.Controller().Init()
	userSvc = ioc.Controller().Get(Name).(Service)
}

func TestUserServiceImpl_CreateUser(t *testing.T) {
	req := NewCreateUserRequest()
	req.Username = "admin"
	req.Password = "12345"
	u, err := userSvc.CreateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func TestUserServiceImpl_DeleteUser(t *testing.T) {
	err := userSvc.DeleteUser(ctx, &DeleteUserRequest{ID: 7})
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserServiceImpl_UpdateUser(t *testing.T) {
	req := NewUpdateUser()
	req.ID = 22
	err := userSvc.UpdateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserServiceImpl_DescribeUser(t *testing.T) {
	req := NewDescribeUserRequestByUsername("admin")
	ins, err := userSvc.DescribeUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins, utils.PasswordCompare(ins.Password, "123456"))
}
