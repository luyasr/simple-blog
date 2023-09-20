package user

import (
	"context"
	"github.com/luyasr/simple-blog/pkg/utils"
	"testing"
)

var (
	userSvc *ServiceImpl
	ctx     = context.Background()
)

func init() {
	userSvc = NewServiceImpl()
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
	err := userSvc.DeleteUser(ctx, &DeleteUserRequest{Id: 7})
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
