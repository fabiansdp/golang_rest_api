package controllers

import (
	"net/http"

	"github.com/fabiansdp/golang_rest_api/config"
	"github.com/fabiansdp/golang_rest_api/models"
	"github.com/gin-gonic/gin"
)

type CreateShopInput struct {
	Nama      string `json:"nama" binding:"required"`
	Jalan     string `json:"jalan" binding:"required"`
	Kecamatan string `json:"kecamatan" binding:"required"`
	Provinsi  string `json:"provinsi" binding:"required"`
}

type UpdateShopInput struct {
	Nama      string `json:"nama"`
	Jalan     string `json:"jalan"`
	Kecamatan string `json:"kecamatan"`
	Provinsi  string `json:"provinsi"`
}

// GET Request
// Find all shops
func GetShops(c *gin.Context) {
	var shops []models.Shop

	config.DB.Find(&shops)

	c.JSON(http.StatusOK, gin.H{"data": shops})
}

// GET Request
// Find a single shop
func GetShop(c *gin.Context) {
	var shop models.Shop

	if err := config.DB.First(&shop, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": shop})
}

// POST Request
// Create a shop
func CreateShop(c *gin.Context) {
	var input CreateShopInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shop := models.Shop{
		Nama:      input.Nama,
		Jalan:     input.Jalan,
		Kecamatan: input.Kecamatan,
		Provinsi:  input.Provinsi,
	}

	config.DB.Create(&shop)

	c.JSON(http.StatusOK, gin.H{"data": shop})
}

// PATCH Request
// Update existing shop
func UpdateShop(c *gin.Context) {
	var shop models.Shop

	if err := config.DB.First(&shop, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateShopInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&shop).Updates(models.Shop{
		Nama:      input.Nama,
		Jalan:     input.Jalan,
		Kecamatan: input.Kecamatan,
		Provinsi:  input.Provinsi,
	})

	c.JSON(http.StatusOK, gin.H{"data": shop})
}

// DELETE Request
// Delete a shop
func DeleteShop(c *gin.Context) {
	var shop models.Shop

	if err := config.DB.First(&shop, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	config.DB.Delete(&shop)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
