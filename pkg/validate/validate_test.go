package validate

import "testing"

type test struct {
	Username string `json:"username" validate:"required" label:"用户名"`
}

func TestStruct(t *testing.T) {
	err := Struct(test{Username: ""})
	if err != nil {
		t.Fatal(err)
	}
}
