package models

import "time"

type LogProduk struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	NamaProduk   string    `json:"nama_produk"`
	Slug         string    `json:"slug"`
	HargaReseller int      `json:"harga_reseller"`
	HargaKonsumen int      `json:"harga_konsumen"`
	Stok         int       `json:"stok"`
	IDToko       uint      `json:"id_toko"`
	IDKategori   uint      `json:"id_category"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
