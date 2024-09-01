package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jamesmpuya/tapindigital/tapinapi-gin/controller"
)

func UserRoute(router *gin.Engine) {
	router.GET("/:id", controller.GetUser)
	router.GET("/", controller.GetUsers)
	router.POST("/", controller.CreateUser)
	router.DELETE("/:id", controller.DeleteUser)
	router.PUT("/:id", controller.UpdateUser)
}
