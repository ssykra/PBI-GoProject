package models

import "time"

type Trx struct {
	ID              uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	IDUser          uint      `json:"id_user"`  // FK ke User
	AlamatPengiriman uint     `json:"alamat_pengiriman"` // FK ke Alamat
	HargaTotal      int       `json:"harga_total"`
	KodeInvoice     string    `json:"kode_invoice"`
	MethodBayar     string    `json:"method_bayar"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DetailTrx       []DetailTrx `json:"detail_trx" gorm:"foreignKey:IDTrx"`
}
