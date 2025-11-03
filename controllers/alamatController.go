package controllers

import (
	"authentication/config"
	"authentication/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Tambah alamat baru
func CreateAlamat(c *gin.Context) {
	userID := c.GetUint("user_id")

	var input models.Alamat
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.UserID = userID
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan alamat"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Alamat berhasil ditambahkan", "data": input})
}

// Lihat semua alamat user login
func GetMyAlamat(c *gin.Context) {
	userID := c.GetUint("user_id")
	var alamats []models.Alamat

	if err := config.DB.Where("user_id = ?", userID).Find(&alamats).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil alamat"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Berhasil mengambil alamat", "data": alamats})
}

// Update alamat milik sendiri
func UpdateAlamat(c *gin.Context) {
	userID := c.GetUint("user_id")
	id := c.Param("id")

	var alamat models.Alamat
	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).First(&alamat).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Alamat tidak ditemukan"})
		return
	}

	var input models.Alamat
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	alamat.NamaPenerima = input.NamaPenerima
	alamat.NoHPPenerima = input.NoHPPenerima
	alamat.DetailAlamat = input.DetailAlamat
	alamat.KodePos = input.KodePos
	alamat.UpdatedAt = time.Now()

	config.DB.Save(&alamat)

	c.JSON(http.StatusOK, gin.H{"message": "Alamat berhasil diperbarui", "data": alamat})
}

// Hapus alamat milik sendiri
func DeleteAlamat(c *gin.Context) {
	userID := c.GetUint("user_id")
	id := c.Param("id")

	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Alamat{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus alamat"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Alamat berhasil dihapus"})
}
