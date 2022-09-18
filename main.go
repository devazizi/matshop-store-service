package main

import (
	"os"
	"store/adapters"
	"store/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	loadEnvVars()

	router := gin.Default()
	dbEngine := adapters.InitDB(os.Getenv("DATABASE_DSN"))
	routes.RouteV1(router, dbEngine)

	router.Run()
}

func loadEnvVars() {
	godotenv.Load(".env")
}
