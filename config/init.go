package config

import (
	"fmt"
	//"nwd/controller/waiting"

	"github.com/BurntSushi/toml"
)

func Init() {
	once.Do(func() {
		conf = new(Config)
		if _, err := toml.DecodeFile("../config/config.toml", conf); err != nil {
			fmt.Printf("decode config file fail, err:%s\n", err.Error())
		}
	})
}
