package models

import "time"

type Produk struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	NamaProduk   string    `json:"nama_produk" gorm:"type:varchar(255);not null"`
	Slug         string    `json:"slug" gorm:"type:varchar(255);unique"`
	HargaReseller int       `json:"harga_reseller"`
	HargaKonsumen int       `json:"harga_konsumen"`
	Stok         int       `json:"stok"`
	IDToko       uint      `json:"id_toko"`    // FK ke Toko
	IDKategori   uint      `json:"id_category"` // FK ke Kategori
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	FotoProduk   []FotoProduk `json:"foto_produk" gorm:"foreignKey:IDProduk"`
}
