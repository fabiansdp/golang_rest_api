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
	var dorayakis []models.Dorayaki
	var inventory []dto.Inventory
	var output dto.GetShopOutput

	err := config.DB.First(&shop, c.Param("id")).Error

	if err != nil {
		res := helper.BuildErrorResponse("Record not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}

	// Find associated dorayakis and get quantity
	config.DB.Model(&shop).Association("Dorayakis").Find(&dorayakis)
	config.DB.Model(&dorayakis).
		Select(
			"dorayakis.*",
			"shop_dorayakis.quantity",
		).
		Joins("JOIN shop_dorayakis ON dorayakis.id = shop_dorayakis.dorayaki_id").
		Scan(&inventory)

	output.ShopInfo = shop
	output.Inventory = inventory

	res := helper.BuildResponse(true, "OK", output)

	c.JSON(http.StatusOK, res)
}

// POST Request
// Create a shop
func CreateShop(c *gin.Context) {
	var input dto.CreateShopInput

	err := c.ShouldBindJSON(&input)
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
func AddInventory(c *gin.Context) {
	var input dto.AddInventoryInput
	var shop models.Shop
	var dorayaki models.Dorayaki

	err := c.ShouldBindJSON(&input)

	if err != nil {
		res := helper.BuildErrorResponse("Add inventory failed", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := config.DB.First(&shop, input.ShopID).Error; err != nil {
		res := helper.BuildErrorResponse("Record not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if err := config.DB.First(&dorayaki, input.DorayakiID).Error; err != nil {
		res := helper.BuildErrorResponse("Record not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}

	shop_dorayaki := models.ShopDorayaki{
		DorayakiID: input.DorayakiID,
		ShopID:     input.ShopID,
		Quantity:   input.Quantity,
	}

	config.DB.Create(&shop_dorayaki)

	res := helper.BuildResponse(true, "OK", shop_dorayaki)

	c.JSON(http.StatusOK, res)
}

// PATCH Request
// Update a shop inventory
func UpdateInventory(c *gin.Context) {
	var input dto.UpdateInventoryInput

	if err := c.ShouldBindJSON(&input); err != nil {
		res := helper.BuildErrorResponse("Not JSON binded", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}

	var shop_dorayaki models.ShopDorayaki

	err := config.DB.Where("dorayaki_id = ? AND shop_id = ?", input.DorayakiID, input.ShopID).First(&shop_dorayaki).Error

	if err != nil {
		res := helper.BuildErrorResponse("Record not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if input.Quantity >= 0 {
		err := config.DB.Model(&shop_dorayaki).Update("quantity", input.Quantity).Error

		if err != nil {
			res := helper.BuildErrorResponse("Error in updating inventory", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, res)
			return
		}

		res := helper.BuildResponse(true, "OK", shop_dorayaki)
		c.JSON(http.StatusOK, res)

	} else {
		res := helper.BuildErrorResponse("Quantity should be >= 0", "", helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}
}

// PATCH Request
// Move inventory from one shop to another
func MoveInventory(c *gin.Context) {
	var input dto.MoveInventoryInput

	// If request not binded
	if err := c.ShouldBindJSON(&input); err != nil {
		res := helper.BuildErrorResponse("Not JSON binded", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}

	var donator models.ShopDorayaki
	var recipient models.ShopDorayaki

	errDonator := config.DB.Where("dorayaki_id = ? AND shop_id = ?", input.DorayakiID, c.Param("id")).First(&donator).Error

	// If donator not found
	if errDonator != nil {
		res := helper.BuildErrorResponse("Record not found", errDonator.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}

	// If donated quantity > than available stock
	if input.Quantity > donator.Quantity {
		res := helper.BuildErrorResponse("Not enough dorayaki", "", helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}

	errRecipient := config.DB.Where("dorayaki_id = ? AND shop_id = ?", input.DorayakiID, input.RecipientID).First(&recipient).Error

	if errRecipient != nil {
		recipient := models.ShopDorayaki{
			DorayakiID: input.DorayakiID,
			ShopID:     input.RecipientID,
			Quantity:   input.Quantity,
		}
		config.DB.Create(&recipient)
	} else {
		config.DB.Model(&recipient).Update("quantity", recipient.Quantity+input.Quantity)
	}

	config.DB.Model(&donator).Update("quantity", donator.Quantity-input.Quantity)

	res := helper.BuildResponse(true, "OK", donator)

	c.JSON(http.StatusOK, res)
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

	config.DB.Select("Dorayakis").Delete(&shop)

	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	c.JSON(http.StatusOK, res)
}
