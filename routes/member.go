package routes

import (
	"hello/controller"

	"github.com/gin-gonic/gin"
)

func MemberRoute(router *gin.Engine) {
	router.GET("/members", controller.Members_Get)
	router.POST("/members", controller.Members_Post)
	router.PATCH("/member/:id", controller.Member_Patch)
	router.DELETE("/member/:id", controller.Member_Del)
	router.GET("/member/:id", controller.Member_Get)
}

