package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/luyasr/simple-blog/pkg/logger"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"sync"
)

var C = new(Config)

type Config struct {
	Server Server `json:"server"`
	Mysql  Mysql  `json:"mysql"`
}

type Server struct {
	Port int `json:"port"`
}

type Mysql struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DB       string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
	Conn     *gorm.DB
	Lock     sync.Mutex
}

func init() {
	NewConfig()
}

func NewConfig() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err := viper.Unmarshal(C); err != nil {
		panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		logger.L.Info().Msg(fmt.Sprintf("Config file changed: %s", e.Name))
		if err := viper.Unmarshal(C); err != nil {
			panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
		}
	})
	viper.WatchConfig()
}

func (m *Mysql) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.DB)
}

func (m *Mysql) GetConn() *gorm.DB {
	if m.Conn == nil {
		m.Lock.Lock()
		defer m.Lock.Unlock()

		conn, err := gorm.Open(mysql.Open(m.DSN()), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		m.Conn = conn
	}

	return m.Conn
}
