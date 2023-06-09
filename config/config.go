package config

import (
	"fmt"
	"sync"
)

type Config struct {
	Mysql  *mysql  `json:"mysql"`
	Server *server `json:"server"`
	Redis  *redis  `json:"redis"`
}

type IConfig interface {
	GetMysqlConf() string
	GetRedisConf() *redis
}

type mysql struct {
	Host    string `json:"host"`
	Port    int    `json:"port"`
	User    string `json:"user"`
	Passwd  string `json:"passwd"`
	DB      string `json:"db"`
	Timeout string `json:"timeout"`
}

type server struct {
	Server    string `json:"server"`
	Env       string `json:"env"`
	LogFile   string `json:"logFile"`
	LogRotate int    `json:"logRotate"`
}

type redis struct {
	Host   string `json:"host"`
	Port   int    `json:"port"`
	Passwd string `json:"passwd"`
	DB     int    `json:"db"`
}

var conf *Config
var once sync.Once

func GetConfig() IConfig {
	return conf
}

func (c *Config) GetMysqlConf() string {
	mysqlconf := conf.Mysql
	if mysqlconf == nil {
		panic("mysql config is null")
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", mysqlconf.User, mysqlconf.Passwd, mysqlconf.Host, mysqlconf.Port, mysqlconf.DB, mysqlconf.Timeout)
}

func (c *Config) GetRedisConf() *redis {
	return conf.Redis
}
