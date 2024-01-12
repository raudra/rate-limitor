package main

import (
	"github.com/gin-gonic/gin"
	"github.com/raudra/rate-limitor/rule-manager/config"
	"github.com/raudra/rate-limitor/rule-manager/src/controllers"
)

func init() {
	config.Init()
}

func Start() {
	router := gin.Default()
	router.Use(gin.Recovery())

	v1 := router.Group("/api/v1")
	{
		v1.GET("/rules", controllers.GetRules)
	}

	router.Run()
}
