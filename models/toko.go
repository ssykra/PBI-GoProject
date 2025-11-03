package models

import "time"

type Toko struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint      `json:"user_id" gorm:"not null"` // FK ke tabel users.user_id
	NamaToko  string    `json:"nama_toko" gorm:"type:varchar(255);not null"`
	UrlToko   string    `json:"url_toko" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
