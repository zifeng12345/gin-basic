package config

import (
	"fmt"
	"sync"
)

type Config struct {
	Mysql *mysql `json:"mysql"`
}

type IConfig interface {
	GetMysqlConf() string
}

type mysql struct {
	Host    string `json:"host"`
	Port    int    `json:"port"`
	User    string `json:"user"`
	Passwd  string `json:"passwd"`
	DB      string `json:"db"`
	Timeout string `json:"timeout"`
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
