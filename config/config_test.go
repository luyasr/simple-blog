package config_test

import (
	"github.com/luyasr/simple-blog/config"
	"testing"
)

func TestNewConfig(t *testing.T) {
	t.Log(config.C)
}

func TestMysql_GetConn(t *testing.T) {
	t.Log(config.C.Mysql.GetConn())
}
