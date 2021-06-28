package controllers

import (
	"net/http"

	"github.com/fabiansdp/golang_rest_api/config"
	"github.com/fabiansdp/golang_rest_api/dto"
	"github.com/fabiansdp/golang_rest_api/helper"
	"github.com/fabiansdp/golang_rest_api/models"
	"github.com/gin-gonic/gin"
)

// GET Request
// Find all shops
func GetShops(c *gin.Context) {
	var shops []models.Shop

	config.DB.Find(&shops)

	res := helper.BuildResponse(true, "OK", shops)

	c.JSON(http.StatusOK, res)
}

// GET Request
// Find a single shop and its inventory
func GetShop(c *gin.Context) {
	var shop models.Shop

	err := config.DB.Preload("Dorayakis").First(&shop, c.Param("id")).Error

	if err != nil {
		res := helper.BuildErrorResponse("Record not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := helper.BuildResponse(true, "OK", shop)

	c.JSON(http.StatusOK, res)
}

// POST Request
// Create a shop
func CreateShop(c *gin.Context) {
	var input dto.CreateShopInput

	err := c.ShouldBind(&input)
	if err != nil {
		res := helper.BuildErrorResponse("Create shop failed", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}

	shop := models.Shop{
		Nama:      input.Nama,
		Jalan:     input.Jalan,
		Kecamatan: input.Kecamatan,
		Provinsi:  input.Provinsi,
	}

	config.DB.Create(&shop)

	res := helper.BuildResponse(true, "OK", shop)

	c.JSON(http.StatusOK, res)
}

// POST Request
// Add dorayaki to shop
func AddDorayaki(c *gin.Context) {
	var input dto.AddDorayakiInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// PATCH Request
// Update existing shop
func UpdateShop(c *gin.Context) {
	var shop models.Shop

	if err := config.DB.First(&shop, c.Param("id")).Error; err != nil {
		res := helper.BuildErrorResponse("Record not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}

	var input dto.UpdateShopInput
	if err := c.ShouldBindJSON(&input); err != nil {
		res := helper.BuildErrorResponse("Not JSON binded", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}

	config.DB.Model(&shop).Updates(models.Shop{
		Nama:      input.Nama,
		Jalan:     input.Jalan,
		Kecamatan: input.Kecamatan,
		Provinsi:  input.Provinsi,
	})

	res := helper.BuildResponse(true, "OK", shop)
	c.JSON(http.StatusOK, res)
}

// DELETE Request
// Delete a shop
func DeleteShop(c *gin.Context) {
	var shop models.Shop

	if err := config.DB.First(&shop, c.Param("id")).Error; err != nil {
		res := helper.BuildErrorResponse("Record not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}

	config.DB.Delete(&shop)

	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	c.JSON(http.StatusOK, res)
}
