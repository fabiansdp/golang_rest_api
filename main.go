package main

import (
	"github.com/fabiansdp/golang_rest_api/config"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello",
		})
	})

	config.ConnectDatabase()

	router.Run()
}
