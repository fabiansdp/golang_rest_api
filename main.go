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
	router.DELETE("/dorayakis/:id", controllers.DeleteDorayaki)

	router.GET("/shops", controllers.GetShops)
	router.GET("/shops/:id", controllers.GetShop)
	router.POST("/shops", controllers.CreateShop)
	router.PATCH("/shops/:id", controllers.UpdateShop)
	router.DELETE("/shops/:id", controllers.DeleteShop)

	router.Run(":8080")
}
