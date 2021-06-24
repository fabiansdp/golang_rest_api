package main

import (
	"github.com/fabiansdp/golang_rest_api/config"
	"github.com/fabiansdp/golang_rest_api/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config.ConnectDatabase()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello",
		})
	})

	router.GET("/dorayakis", controllers.GetDorayakis)
	router.GET("/dorayakis/:id", controllers.GetDorayaki)
	router.POST("/dorayakis", controllers.CreateDorayaki)
	router.PATCH("/dorayakis/:id", controllers.UpdateDorayaki)

	router.Run()
}
