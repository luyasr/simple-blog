package utils

import (
	"dario.cat/mergo"
	"fmt"
	"testing"
)

type Foo struct {
	A string `json:"a"`
	B int64  `json:"b"`
}

func TestStructToMap(t *testing.T) {
	a := Foo{A: "test", B: 2}
	m, err := StructToMap(a)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(m)
}

func TestMerge(t *testing.T) {
	a := Foo{A: "1", B: 1}
	b := Foo{B: 2}
	err := Merge(&a, b, mergo.WithOverride, mergo.WithoutDereference)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(a)
}
