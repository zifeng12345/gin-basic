package main

import (
	"nwd/config"
	"nwd/routers"
	"nwd/shared/database"

	wctl "nwd/controller/waiting"
	wrep "nwd/repository/waiting"
	wsrv "nwd/service/waiting"
)

func main() {
	config.Init()
	dsn := config.GetConfig().GetMysqlConf()
	database.Init(dsn)
	waitingInit()

	routers.Routers()
}

func waitingInit() {
	wrep := wrep.GetWaitingRepo()
	wsrv.Init(wrep)
	wsrv := wsrv.GetWaitingSrv()
	wctl.Init(wsrv)
}
