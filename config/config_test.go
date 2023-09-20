package config

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	t.Log(C)
}

func TestMysql_GetConn(t *testing.T) {
	t.Log(C.Mysql.GetConn())
}
