package controllers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/fabiansdp/golang_rest_api/config"
	"github.com/fabiansdp/golang_rest_api/helper"
	"github.com/fabiansdp/golang_rest_api/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GET Request
// Find all dorayakis
func GetDorayakis(c *gin.Context) {
	var dorayakis []models.Dorayaki

	config.DB.Find(&dorayakis)

	res := helper.BuildResponse(true, "OK", dorayakis)

	c.JSON(http.StatusOK, res)
}

// GET Request
// Find a single dorayaki
func GetDorayaki(c *gin.Context) {
	var dorayaki models.Dorayaki

	if err := config.DB.First(&dorayaki, c.Param("id")).Error; err != nil {
		res := helper.BuildErrorResponse("Record not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := helper.BuildResponse(true, "OK", dorayaki)

	c.JSON(http.StatusOK, res)
}

// GET a dorayaki image
func GetDorayakiImg(c *gin.Context) {
	filename := c.Param("filename")

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "inline; filename="+filename)
	c.File("images/" + filename)
}

// POST Request
// Create New Dorayaki
func CreateDorayaki(c *gin.Context) {
	rasa := c.PostForm("rasa")
	deskripsi := c.PostForm("deskripsi")
	file, errFile := c.FormFile("file")

	if errFile != nil {
		res := helper.BuildErrorResponse("File not found", errFile.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}

	// Retrieve file information
	extension := filepath.Ext(file.Filename)

	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	newFileName := uuid.New().String() + extension

	// Path to images folder
	path := "images/" + newFileName

	if err := c.SaveUploadedFile(file, path); err != nil {
		res := helper.BuildErrorResponse("Unable to save file", errFile.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}

	dorayaki := models.Dorayaki{Rasa: rasa, Deskripsi: deskripsi, Gambar: newFileName}

	config.DB.Create(&dorayaki)

	res := helper.BuildResponse(true, "OK", dorayaki)

	c.JSON(http.StatusOK, res)
}

// PATCH Request
// Update existing dorayaki
func UpdateDorayaki(c *gin.Context) {
	var dorayaki models.Dorayaki

	rasa := c.PostForm("rasa")
	deskripsi := c.PostForm("deskripsi")
	file, errFile := c.FormFile("file")

	if err := config.DB.First(&dorayaki, c.Param("id")).Error; err != nil {
		res := helper.BuildErrorResponse("Record not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if errFile != nil {
		config.DB.Model(&dorayaki).Updates(models.Dorayaki{Rasa: rasa, Deskripsi: deskripsi})
	} else {
		// Remove file
		err := os.Remove("images/" + dorayaki.Gambar)

		if err != nil {
			res := helper.BuildErrorResponse("Cannot delete file", err.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, res)
			return
		}

		// Retrieve file information
		extension := filepath.Ext(file.Filename)

		// Generate random file name for the new uploaded file so it doesn't override the old file with same name
		newFileName := uuid.New().String() + extension

		// Path to images folder
		path := "images/" + newFileName

		if err := c.SaveUploadedFile(file, path); err != nil {
			res := helper.BuildErrorResponse("Unable to save file", errFile.Error(), helper.EmptyObj{})
			c.JSON(http.StatusBadRequest, res)
			return
		}

		config.DB.Model(&dorayaki).Updates(models.Dorayaki{Rasa: rasa, Deskripsi: deskripsi, Gambar: newFileName})
	}

	res := helper.BuildResponse(true, "OK", dorayaki)
	c.JSON(http.StatusOK, res)
}

// DELETE Request
// Delete a dorayaki
func DeleteDorayaki(c *gin.Context) {
	var dorayaki models.Dorayaki
	var shop_dorayaki models.ShopDorayaki

	if err := config.DB.First(&dorayaki, c.Param("id")).Error; err != nil {
		res := helper.BuildErrorResponse("Record not found", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}

	// Remove file
	err := os.Remove("images/" + dorayaki.Gambar)

	if err != nil {
		res := helper.BuildErrorResponse("Cannot delete file", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
		return
	}

	config.DB.Where("dorayaki_id = ?", c.Param("id")).Delete(&shop_dorayaki)
	config.DB.Delete(&dorayaki)

	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	c.JSON(http.StatusOK, res)
}
