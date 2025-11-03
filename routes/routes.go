package routes

import (
	"authentication/config"
	"authentication/controllers"
	"authentication/middleware"
	"authentication/repositories"
	"authentication/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Auth
	router.POST("/signup", controllers.Signup())
	router.POST("/login", controllers.Login())

	// Protected Route
	protected := router.Group("/", middleware.Authenticate())
	{
		protected.GET("/users", controllers.GetUsers())
		protected.GET("/user/:id", controllers.GetUser())

		// Route Toko
		protected.GET("/toko/me", controllers.GetMyToko)
		protected.PUT("/toko/me", controllers.UpdateMyToko)

		// Alamat
		protected.POST("/alamat", controllers.CreateAlamat)
		protected.GET("/alamat", controllers.GetMyAlamat)
		protected.PUT("/alamat/:id", controllers.UpdateAlamat)
		protected.DELETE("/alamat/:id", controllers.DeleteAlamat)

		admin := protected.Group("/admin", middleware.AdminOnly())
		{
			admin.GET("/toko", controllers.GetAllToko)
		}

		// Route Produk
		produkRepo := repositories.NewProdukRepository(config.DB)
		produkService := services.NewProdukService(produkRepo)
		produkController := controllers.NewProdukController(produkService)

		produk := protected.Group("/produk")
		{
			produk.POST("/", produkController.CreateProduk)
			produk.PUT("/:id", produkController.UpdateProduk)
			produk.GET("/:id", produkController.GetProduk)
			produk.GET("/", produkController.GetAllProduk)
		}

		// Route Transaksi
		trxService := services.NewTrxService(config.DB)
		trxController := controllers.NewTrxController(trxService)

		trx := protected.Group("/trx")
		{
			trx.POST("/", trxController.CreateTrx)
			trx.GET("/:id", trxController.GetTrx)
		}
	}

	// Route Kategori
	kategoriRepo := repositories.NewKategoriRepository(config.DB)
	kategoriService := services.NewKategoriService(kategoriRepo)
	kategoriController := controllers.NewKategoriController(kategoriService)

	kategori := router.Group("/kategori")
	{
		kategori.GET("/", kategoriController.GetAllKategori)
		kategori.GET("/:id", kategoriController.GetKategoriByID)

		adminKategori := kategori.Group("")
		adminKategori.Use(middleware.Authenticate(), middleware.AdminOnly())
		{
			adminKategori.POST("/", kategoriController.CreateKategori)
			adminKategori.PUT("/:id", kategoriController.UpdateKategori)
			adminKategori.DELETE("/:id", kategoriController.DeleteKategori)
		}
	}
}
