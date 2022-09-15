package main

import (
	"store/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	routes.RouteV1(router)
}
