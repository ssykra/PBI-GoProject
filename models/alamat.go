package models

import "time"

type Alamat struct {
	ID            uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID        uint      `json:"user_id" gorm:"not null"`
	NamaPenerima  string    `json:"nama_penerima" gorm:"type:varchar(100);not null"`
	NoHPPenerima  string    `json:"no_hp_penerima" gorm:"type:varchar(20);not null"`
	ProvinsiID    string    `json:"provinsi_id"`
	KotaID        string    `json:"kota_id"`
	KecamatanID   string    `json:"kecamatan_id"`
	KelurahanID   string    `json:"kelurahan_id"`
	DetailAlamat  string    `json:"detail_alamat" gorm:"type:text"`
	KodePos       string    `json:"kode_pos"`
	IsDefault     bool      `json:"is_default" gorm:"default:false"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
