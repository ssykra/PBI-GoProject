package main

import (
	"authentication/config"
	"authentication/helpers"
	"authentication/models"
	"authentication/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect database
	db := config.ConnectDB()

	// Migrasi semua model
	err := db.AutoMigrate(
		&models.User{},
		&models.Toko{},
		&models.Alamat{},
		&models.Kategori{},
		&models.Produk{},
		&models.FotoProduk{},
		&models.LogProduk{},
		&models.Trx{},
		&models.DetailTrx{},
	)
	if err != nil {
		log.Fatalf("Gagal migrasi database: %v", err)
	}

	// Generate JWT key
	key := config.GenerateRandomKey()
	helpers.SetJWTKey(key)

	// Setup Gin
	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r)

	// Print semua route
	for _, ri := range r.Routes() {
		log.Println("ROUTE:", ri.Method, ri.Path)
	}

	// Run server
	port := "8080"
	log.Println("Server running on port:", port)
	r.Run(":" + port)
}
