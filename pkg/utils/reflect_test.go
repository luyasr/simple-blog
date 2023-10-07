package utils

import (
	user2 "github.com/luyasr/simple-blog/app/user"
	"testing"
)

func TestUpdateNonZeroFields(t *testing.T) {
	u := user2.NewUser(&user2.CreateUserRequest{
		Username: "admin",
		Password: "12345",
		Role:     user2.RoleMember,
	})
	fields, err := UpdateNonZeroFields(u)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(fields)
}
