package config

import (
	"fmt"
	"log"
	"authentication/models" 
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	user := "root"
	password := "" 
	host := "127.0.0.1"
	port := "3306"
	dbname := "evermos_db"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal konek ke database MySQL: %v", err)
	}

	// Auto-migrate semua model
	err = db.AutoMigrate(
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
		log.Fatalf("Gagal melakukan migrasi: %v", err)
	}

	log.Println("✅ Berhasil konek dan migrasi database MySQL")
	DB = db
	return db
}

