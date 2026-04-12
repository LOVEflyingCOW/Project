package main

import (
	"exchangeapp/config"
	"exchangeapp/router"
)

func main() {
	config.InitConfig()
	r := router.SetupRouter()
	r.Run(config.AppConfig.App.Port) // listen and serve on 0.0.0.0:3000
}
