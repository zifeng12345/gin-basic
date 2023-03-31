package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"nwd/config"
	"nwd/routers"
	"nwd/shared/database"
)

func main() {
	config.Init()
	defer stop()
	routers.Routers()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, os.Kill)

	sig := <-signalChan
	fmt.Printf("catch exit signal:%v and exiting\n", sig)

	//Wait the serivce release the resourses
	time.Sleep(time.Second * 2)
}

//close mysql/redis/mqtt resources...
func stop() {
	database.Stop()
}
