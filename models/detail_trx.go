package models

import "time"

type DetailTrx struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	IDTrx       uint      `json:"id_trx"`        // FK ke Trx
	IDProduk    uint      `json:"id_produk"`
	IDLogProduk uint      `json:"id_log_produk"` // FK ke LogProduk
	IDToko      uint      `json:"id_toko"`       // FK ke Toko
	Kuantitas   int       `json:"kuantitas"`
	HargaTotal  int       `json:"harga_total"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
