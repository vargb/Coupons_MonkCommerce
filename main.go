package main

import (
	"monkCommerce/config"
	"monkCommerce/services"
)

func main() {
	conf, err := config.InitConfig()
	if err != nil {
		return
	}
	r, err := services.Initialize(conf)
	if err != nil {
		config.GetLogger().Error("error in initializing server...")
		return
	}
	r.Run(conf.Server.Port)
}
