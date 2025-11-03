package controllers

import (
	"authentication/config"
	"authentication/helpers"
	"authentication/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validasi input
		if err := validate.Struct(user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Cek email / no. telp sudah terdaftar
		var count int64
		config.DB.Model(&models.User{}).
			Where("email = ? OR no_telp = ?", user.Email, user.NoTelp).
			Count(&count)

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email atau No. Telepon sudah terdaftar"})
			return
		}

		// Hash password
		hashedPwd, err := helpers.HashPassword(user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal enkripsi password"})
			return
		}

		user.Password = hashedPwd
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()

		// Simpan user ke database
		if err := config.DB.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan user"})
			return
		}

		// Generate token JWT (uint ID)
		accessToken, refreshToken, err := helpers.GenerateToken(user.Email, user.ID, user.IsAdmin)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat token"})
			return
		}

		user.Token = accessToken
		user.RefreshToken = refreshToken
		config.DB.Save(&user)

		// Membuat toko otomatis untuk user baru
		toko := models.Toko{
			UserID:    user.ID,
			NamaToko:  fmt.Sprintf("Toko %s", user.Nama),
			UrlToko:   fmt.Sprintf("https://toko.com/%s", user.Nama),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := config.DB.Create(&toko).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat toko untuk user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Registrasi berhasil, toko otomatis dibuat",
			"user":    user,
			"toko":    toko,
		})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.User
		var foundUser models.User

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Cari user berdasarkan email
		if err := config.DB.Where("email = ?", input.Email).First(&foundUser).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Email atau password salah"})
			return
		}

		// Verifikasi password
		passwordIsValid, _ := helpers.VerifyPassword(foundUser.Password, input.Password)
		if !passwordIsValid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Email atau password salah"})
			return
		}

		token, refreshToken, err := helpers.GenerateToken(foundUser.Email, foundUser.ID, foundUser.IsAdmin)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Update token di database
		if err := helpers.UpdateAllToken(token, refreshToken, foundUser.ID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui token"})
			return
		}



		c.JSON(http.StatusOK, gin.H{
			"message":        "Login berhasil",
			"user":           foundUser,
			"token":          token,
			"refresh_token":  refreshToken,
		})
	}
}

// ===========================
// Tambahan fungsi untuk route
// ===========================

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []models.User
		if err := config.DB.Find(&users).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data user"})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("id")
		var user models.User

		// ✅ gunakan kolom id, bukan user_id
		if err := config.DB.Where("id = ?", userID).First(&user).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}
