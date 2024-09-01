package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jamesmpuya/tapindigital/tapinapi-gin/config"
	"github.com/jamesmpuya/tapindigital/tapinapi-gin/routes"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	router.Run(":8080")
}
