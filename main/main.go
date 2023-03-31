package main

import (
	"nwd/config"
	"nwd/routers"
)

func main() {
	config.Init()

	routers.Routers()
}
