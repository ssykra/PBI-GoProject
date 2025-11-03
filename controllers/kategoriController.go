package controllers

import (
	"authentication/models"
	"authentication/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type KategoriController struct {
	KategoriService *services.KategoriService
}

func NewKategoriController(service *services.KategoriService) *KategoriController {
	return &KategoriController{KategoriService: service}
}

func (c *KategoriController) GetAllKategori(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "10")
	search := ctx.DefaultQuery("search", "")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	kategoris, total, err := c.KategoriService.GetAllKategori(page, limit, search)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":       kategoris,
		"total_data": total,
		"page":       page,
		"limit":      limit,
	})
}

// GET /kategori/:id
func (c *KategoriController) GetKategoriByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	kategori, err := c.KategoriService.GetKategoriByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Kategori tidak ditemukan"})
		return
	}

	ctx.JSON(http.StatusOK, kategori)
}

// POST /kategori (hanya Admin)
func (c *KategoriController) CreateKategori(ctx *gin.Context) {
	var kategori models.Kategori
	if err := ctx.ShouldBindJSON(&kategori); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.KategoriService.CreateKategori(&kategori); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Kategori berhasil dibuat",
		"data":    kategori,
	})
}

// PUT /kategori/:id (hanya Admin)
func (c *KategoriController) UpdateKategori(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var kategori models.Kategori
	if err := ctx.ShouldBindJSON(&kategori); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	kategori.ID = uint(id)

	if err := c.KategoriService.UpdateKategori(&kategori); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Kategori berhasil diperbarui",
		"data":    kategori,
	})
}

// DELETE /kategori/:id (hanya Admin)
func (c *KategoriController) DeleteKategori(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	if err := c.KategoriService.DeleteKategori(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Kategori berhasil dihapus"})
}
