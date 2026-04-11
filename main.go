package main

import (
	"exchangeapp/config"
	"exchangeapp/router"
)

func main() {
	config.InitConfig()
	r := router.SetupRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}
