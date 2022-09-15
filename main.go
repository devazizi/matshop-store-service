package main

import (
	"store/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	loadEnvVars()

	router := gin.Default()

	routes.RouteV1(router)

	router.Run()
}

func loadEnvVars() {
	godotenv.Load(".env")
}
