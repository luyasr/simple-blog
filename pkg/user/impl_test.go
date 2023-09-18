package user_test

import (
	"context"
	"github.com/luyasr/simple-blog/pkg/user"
	"testing"
)

var (
	usrSvc *user.UserServiceImpl
	ctx    = context.Background()
)

func init() {
	usrSvc = user.NewUserServiceImpl()
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

func TestUserServiceImpl_DeleteUser(t *testing.T) {
	err := usrSvc.DeleteUser(ctx, &user.DeleteUserRequest{Id: 7})
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserServiceImpl_UpdateUser(t *testing.T) {
	req := user.NewUpdateUserRequest("7")
	err := usrSvc.UpdateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserServiceImpl_DescribeUser(t *testing.T) {
	req := user.NewDescribeUserRequestByUsername("admin")
	ins, err := usrSvc.DescribeUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
