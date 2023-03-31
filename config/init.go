package config

import (
	"fmt"

	wctl "nwd/controller/waiting"
	wrep "nwd/repository/waiting"
	wsrv "nwd/service/waiting"
	"nwd/shared/database"

	uctl "nwd/controller/users"
	urep "nwd/repository/users"
	usrv "nwd/service/users"

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
