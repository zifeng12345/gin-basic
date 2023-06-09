package config

import (
	"fmt"

	uctl "nwd/controller/users"
	wctl "nwd/controller/waiting"
	urep "nwd/repository/users"
	wrep "nwd/repository/waiting"
	usrv "nwd/service/users"
	wsrv "nwd/service/waiting"
	"nwd/shared/database"
	"nwd/shared/log"
	rds "nwd/shared/redis"

	"github.com/BurntSushi/toml"
)

func Init() {
	once.Do(func() {
		conf = new(Config)
		if _, err := toml.DecodeFile("../config/config.toml", conf); err != nil {
			fmt.Printf("decode config file fail, err:%s\n", err.Error())
			panic("load config file failed")
		}

		dsn := conf.GetMysqlConf()
		database.Init(dsn)

		waitingInit()
		userInit()
		logInit()
		rds.Init(conf.Redis.Host, conf.Redis.Passwd, conf.Redis.Port, conf.Redis.DB)
	})
}

func waitingInit() {
	wrep := wrep.GetWaitingRepo()
	wsrv.Init(wrep)
	wsrv := wsrv.GetWaitingSrv()
	wctl.Init(wsrv)
}

func userInit() {
	urep := urep.GetWaitingRepo()
	usrv.Init(urep)
	usrv := usrv.GetUsers()
	uctl.Init(usrv)
}

func logInit() {
	log.GetLog().WithFile("api", conf.Server.LogFile, conf.Server.LogRotate)
	log.GetLog().Info("", "Service start! NR ")
}
