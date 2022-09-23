package routes

import (
	"hello/controller"

	"github.com/gin-gonic/gin"
)

func BookRoutes(router *gin.Engine) {
	router.GET("/books", controller.Books_Get)
	router.GET("/book/:id",controller.Book_Get)
	router.POST("/books", controller.Books_Post)
	router.PATCH("/book/:id",controller.Books_Update)
}