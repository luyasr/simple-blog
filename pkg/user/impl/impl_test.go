package impl_test

import (
	"context"
	"github.com/luyasr/simple-blog/common"
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

func TestUserServiceImpl_DeleteUser(t *testing.T) {
	err := usrSvc.DeleteUser(ctx, &user.DeleteUserRequest{Id: 7})
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserServiceImpl_UpdateUser(t *testing.T) {
	req := &user.UpdateUserRequest{
		User: &user.User{
			Meta: &common.Meta{
				Id: 7,
			},
			CreateUserRequest: &user.CreateUserRequest{
				Username: "admin",
				Role:     user.RoleAdmin,
			},
		},
	}
	updateUser, err := usrSvc.UpdateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(updateUser)
}

func TestUserServiceImpl_DescribeUser(t *testing.T) {
	req := user.NewDescribeUserRequestByUsername("admin")
	ins, err := usrSvc.DescribeUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
