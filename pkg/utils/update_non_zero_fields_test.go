package utils_test

import (
	"github.com/luyasr/simple-blog/pkg/user"
	"github.com/luyasr/simple-blog/pkg/utils"
	"testing"
)

func TestUpdateNonZeroFields(t *testing.T) {
	u := user.NewUser(&user.CreateUserRequest{
		Username: "admin",
		Password: "12345",
		Role:     user.RoleMember,
	})
	fields, err := utils.UpdateNonZeroFields(u)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(fields)
}
