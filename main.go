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

	dorayakiRoutes := router.Group("api/dorayakis")
	{
		dorayakiRoutes.GET("/", controllers.GetDorayakis)
		dorayakiRoutes.GET("/:id", controllers.GetDorayaki)
		dorayakiRoutes.POST("/", controllers.CreateDorayaki)
		dorayakiRoutes.PATCH("/:id", controllers.UpdateDorayaki)
		dorayakiRoutes.DELETE("/:id", controllers.DeleteDorayaki)
	}

	shopRoutes := router.Group("api/shops")
	{
		shopRoutes.GET("/", controllers.GetShops)
		shopRoutes.GET("/:id", controllers.GetShop)
		shopRoutes.POST("/", controllers.CreateShop)
		shopRoutes.POST("/inventory", controllers.AddDorayaki)
		shopRoutes.PATCH("/:id", controllers.UpdateShop)
		shopRoutes.DELETE("/:id", controllers.DeleteShop)
	}

	router.Run(":8080")
}
