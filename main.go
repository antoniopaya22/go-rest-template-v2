package main

import (
	"ai4am.com/go-restapi/docs"
	"ai4am.com/go-restapi/internal"
	"ai4am.com/go-restapi/internal/config"
	"fmt"
	"os"
)

// Gets variables from environment variables
//
// returns:
//		configFilePath, logFilePath
func getEnvVariables() (string, string) {
	configFilePath := "data/config.yml"
	logFilePath := "logs/restapi.log"
	if os.Getenv("configFilePath") != "" {
		configFilePath = os.Getenv("configFilePath")
	}
	if os.Getenv("logFilePath") != "" {
		logFilePath = os.Getenv("logFilePath")
	}
	return configFilePath, logFilePath
}

// @Golang API REST
// @title Golang REST API Template
// @version 2.0
// @description API REST in Golang with Gin Framework

// @contact.name Antonio Paya Gonzalez

// @host localhost:3000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configFilePath, logFilePath := getEnvVariables()
	configuration := config.GetConfiguration(configFilePath)
	app := internal.App{Config: configuration, App: nil}
	fmt.Println(docs.SwaggerInfo.Title)
	app.Start(logFilePath)
	fmt.Println("Go API REST Running on port " + app.Config.Server.Port)
	fmt.Println("==================>")
	_ = app.App.Run(":" +  app.Config.Server.Port)
}
