package main

import (
	"os"

	"github.com/fabiansdp/golang_rest_api/config"
	"github.com/fabiansdp/golang_rest_api/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	errEnv := godotenv.Load()

	if errEnv != nil {
		panic("Failed to load environment variables")
	}

	imagePath := "./images"
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		err := os.Mkdir(imagePath, os.ModePerm)
		if err != nil {
			panic("Falied to create images folder")
		}
	}

	PORT := os.Getenv("PORT")

	router := gin.Default()
	router.Use(cors.Default())

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
		shopRoutes.PATCH("/:id", controllers.UpdateShop)
		shopRoutes.DELETE("/:id", controllers.DeleteShop)
	}

	// Shop Inventory Routes
	// Jujur ini ga di-group karena error cors kalau di-group tapi
	// berjalan biasa kalau sendirian
	router.GET("/api/file/:filename", controllers.GetDorayakiImg)
	router.POST("api/inventory", controllers.AddInventory)
	router.PUT("api/inventory/:id", controllers.UpdateInventory)
	router.PATCH("api/inventory/:id", controllers.MoveInventory)
	router.DELETE("api/inventory/:id", controllers.DeleteInventory)

	router.Run(":" + PORT)
}
