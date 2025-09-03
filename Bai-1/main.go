package main

import (
	"awesomeProject6/config"
	"awesomeProject6/controller"
	"awesomeProject6/service"
)

func main() {
	APIService := service.NewService()
	APIcontroller := controller.NewController(APIService)

	APIConfig := config.NewConfig()

	APIConfig.GET("/healthz", APIcontroller.GetTime)
	APIConfig.Run(":8080")
}
