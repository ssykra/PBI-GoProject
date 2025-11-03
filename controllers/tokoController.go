package controllers

import (
	"authentication/config"
	"authentication/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Ambil data toko user login
func GetMyToko(c *gin.Context) {
	userID := c.GetUint("user_id") 
	var toko models.Toko

	if err := config.DB.Where("user_id = ?", userID).First(&toko).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Toko tidak ditemukan"})
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil mengambil data toko",
		"data":    toko,
	})
}

// Update toko milik user login
func UpdateMyToko(c *gin.Context) {
	userID := c.GetUint("user_id")
	var toko models.Toko

	if err := config.DB.Where("user_id = ?", userID).First(&toko).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Toko tidak ditemukan"})
		return
	}


	var input struct {
		NamaToko string `json:"nama_toko"`
		UrlToko  string `json:"url_toko"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	toko.NamaToko = input.NamaToko
	toko.UrlToko = input.UrlToko
	toko.UpdatedAt = time.Now()

	config.DB.Save(&toko)

	c.JSON(http.StatusOK, gin.H{
		"message": "Toko berhasil diperbarui",
		"data":    toko,
	})
}

// Hanya admin yang boleh akses semua toko
func GetAllToko(c *gin.Context) {
	var toko []models.Toko
	config.DB.Find(&toko)
	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil mengambil semua toko",
		"data":    toko,
	})
}
