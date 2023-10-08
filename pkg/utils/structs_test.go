package utils

import (
	"fmt"
	"testing"
)

type Foo struct {
	A string `json:"a"`
	B int64  `json:"b"`
}

func TestStruct2Map(t *testing.T) {
	a := Foo{A: "test", B: 2}
	m, err := Struct2Map(a)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(m)
}
