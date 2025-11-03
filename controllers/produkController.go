package controllers

import (
	"authentication/models"
	"authentication/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProdukController struct {
	Service *services.ProdukService
}

func NewProdukController(service *services.ProdukService) *ProdukController {
	return &ProdukController{Service: service}
}

func (c *ProdukController) CreateProduk(ctx *gin.Context) {
	var input models.Produk
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := ctx.GetUint("user_id") // dari middleware auth

	if err := c.Service.CreateProduk(&input, userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, input)
}

func (c *ProdukController) UpdateProduk(ctx *gin.Context) {
	produkID, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	var input map[string]interface{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := ctx.GetUint("user_id")

	if err := c.Service.UpdateProduk(uint(produkID), userID, input); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "produk berhasil diupdate"})
}

func (c *ProdukController) GetProduk(ctx *gin.Context) {
	produkID, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	userID := ctx.GetUint("user_id")

	produk, err := c.Service.GetProduk(uint(produkID), userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "produk tidak ditemukan"})
		return
	}
	ctx.JSON(http.StatusOK, produk)
}

func (c *ProdukController) GetAllProduk(ctx *gin.Context) {
	userID := ctx.GetUint("user_id")
	produks, _ := c.Service.GetAllProduk(userID)
	ctx.JSON(http.StatusOK, produks)
}
