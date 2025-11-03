package models

import "time"

type FotoProduk struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	IDProduk  uint      `json:"id_produk"`  // FK ke Produk
	URL       string    `json:"url" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
