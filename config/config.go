package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"path"
	"runtime"
	"strings"
	"sync"
)

var C = new(Config)

type Config struct {
	Server Server `json:"server"`
	Mysql  Mysql  `json:"mysql"`
	Minio  Minio  `json:"minio"`
}

type Server struct {
	Port  int  `json:"port"`
	Debug bool `json:"debug"`
	Log   Log  `json:"log"`
}

type Log struct {
	Dir string `json:"dir"`
}

type Mysql struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DataBase string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
	Conn     *gorm.DB
	Lock     sync.Mutex
}

type Minio struct {
	Endpoint        string `json:"endpoint"`
	AccessKeyID     string `json:"accessKeyID"`
	SecretAccessKey string `json:"secretAccessKey"`
	UseSSL          bool   `json:"useSSL"`
	BucketName      string `json:"bucketName"`
	Core            *minio.Core
	Lock            sync.Mutex
}

func init() {
	newConfig()
}

func rootPath() string {
	_, filename, _, _ := runtime.Caller(0)
	root := path.Dir(path.Dir(filename))
	return root
}

func newConfig() {
	viper.AddConfigPath(rootPath())
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal errors config file: %s \n", err))
	}

	if err := viper.Unmarshal(C); err != nil {
		panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(C); err != nil {
			panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
		}
	})
	viper.WatchConfig()
}

func (m *Mysql) dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.DataBase)
}

func (m *Mysql) GetConn() *gorm.DB {
	var gormLogMode gormLogger.Interface
	if C.Server.Debug {
		gormLogMode = gormLogger.Default.LogMode(gormLogger.Info)
	} else {
		gormLogMode = gormLogger.Default.LogMode(gormLogger.Silent)
	}

	if m.Conn == nil {
		m.Lock.Lock()
		defer m.Lock.Unlock()

		conn, err := gorm.Open(mysql.Open(m.dsn()), &gorm.Config{
			Logger: gormLogMode,
		})
		if err != nil {
			panic(err)
		}

		m.Conn = conn
	}

	return m.Conn
}

func (m *Minio) options() *minio.Options {
	return &minio.Options{
		Creds:  credentials.NewStaticV4(m.AccessKeyID, m.SecretAccessKey, ""),
		Secure: m.UseSSL,
	}
}

func (m *Minio) NewCore() *minio.Core {
	if m.Core == nil {
		m.Lock.Lock()
		defer m.Lock.Unlock()

		core, err := minio.NewCore(m.Endpoint, m.options())
		if err != nil {
			panic(err)
		}
		m.Core = core
	}

	return m.Core
}
