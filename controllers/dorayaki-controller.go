package controllers

import (
	"net/http"

	"github.com/fabiansdp/golang_rest_api/config"
	"github.com/fabiansdp/golang_rest_api/dto"
	"github.com/fabiansdp/golang_rest_api/models"
	"github.com/gin-gonic/gin"
)

// GET Request
// Find all dorayakis
func GetDorayakis(c *gin.Context) {
	var dorayakis []models.Dorayaki

	config.DB.Find(&dorayakis)

	c.JSON(http.StatusOK, gin.H{"data": dorayakis})
}

// GET Request
// Find a single dorayaki
func GetDorayaki(c *gin.Context) {
	var dorayaki models.Dorayaki

	if err := config.DB.First(&dorayaki, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dorayaki})
}

// POST Request
// Create New Dorayaki
func CreateDorayaki(c *gin.Context) {
	var input dto.CreateDorayakiInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dorayaki := models.Dorayaki{Rasa: input.Rasa, Deskripsi: input.Deskripsi, Gambar: input.Gambar}
	config.DB.Create(&dorayaki)

	c.JSON(http.StatusOK, gin.H{"data": dorayaki})
}

// PATCH Request
// Update existing dorayaki
func UpdateDorayaki(c *gin.Context) {
	var dorayaki models.Dorayaki

	if err := config.DB.First(&dorayaki, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input dto.UpdateDorayakiInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&dorayaki).Updates(models.Dorayaki{Rasa: input.Rasa, Deskripsi: input.Deskripsi, Gambar: input.Gambar})
	c.JSON(http.StatusOK, gin.H{"data": dorayaki})
}

// DELETE Request
// Delete a dorayaki
func DeleteDorayaki(c *gin.Context) {
	var dorayaki models.Dorayaki

	if err := config.DB.First(&dorayaki, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	config.DB.Delete(&dorayaki)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
