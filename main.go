package main

import (
	"log"
	"os"
	app "smtp-artha/src"
	"smtp-artha/src/configs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if os.Getenv("DOCKER_ENVIRONMENT") == "true" {
		log.Println("Running inside a Docker container")
	} else {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func main() {
	gin.SetMode(configs.AppConfig().GinMode)
	server := gin.Default()

	app.Middlewares(server)
	app.Routes(server)
	app.ListenServer(server, configs.AppConfig().Port)
}
