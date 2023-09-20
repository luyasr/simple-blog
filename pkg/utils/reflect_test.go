package utils

import (
	"github.com/luyasr/simple-blog/pkg/user"
	"testing"
)

func TestUpdateNonZeroFields(t *testing.T) {
	u := user.NewUser(&user.CreateUserRequest{
		Username: "admin",
		Password: "12345",
		Role:     user.RoleMember,
	})
	fields, err := UpdateNonZeroFields(u)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(fields)
}
