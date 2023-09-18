package user_test

import (
	"context"
	"testing"
)

var (
	usrSvc *UserServiceImpl
	ctx    = context.Background()
)

func init() {
	usrSvc = NewUserServiceImpl()
}

func TestUserServiceImpl_CreateUser(t *testing.T) {
	req := NewCreateUserRequest()
	req.Username = "admin"
	req.Password = "12345"
	u, err := usrSvc.CreateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func TestUserServiceImpl_DeleteUser(t *testing.T) {
	err := usrSvc.DeleteUser(ctx, &DeleteUserRequest{Id: 7})
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserServiceImpl_UpdateUser(t *testing.T) {
	req := NewUpdateUserRequest("7")
	err := usrSvc.UpdateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserServiceImpl_DescribeUser(t *testing.T) {
	req := NewDescribeUserRequestByUsername("admin")
	ins, err := usrSvc.DescribeUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
