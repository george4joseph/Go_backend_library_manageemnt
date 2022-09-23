package main

import (
	"hello/config"
	"hello/routes"

	"github.com/gin-gonic/gin"
)

// main function
func main() {

	router := gin.Default()
	config.Connect()
	routes.MemberRoute(router)
	routes.BookRoutes(router)
	router.Run()
}
