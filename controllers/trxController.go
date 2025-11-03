package controllers

import (
	"authentication/models"
	"authentication/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TrxController struct {
	Service *services.TrxService
}

func NewTrxController(service *services.TrxService) *TrxController {
	return &TrxController{Service: service}
}

func (c *TrxController) CreateTrx(ctx *gin.Context) {
	var input models.Trx
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := ctx.GetUint("user_id")

	if err := c.Service.CreateTrx(&input, userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, input)
}

func (c *TrxController) GetTrx(ctx *gin.Context) {
	trxID, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	userID := ctx.GetUint("user_id")

	trx, err := c.Service.GetTrx(uint(trxID), userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "transaksi tidak ditemukan"})
		return
	}
	ctx.JSON(http.StatusOK, trx)
}
