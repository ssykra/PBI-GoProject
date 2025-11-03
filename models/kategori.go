package models

import "time"

type Kategori struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Nama      string    `json:"nama" gorm:"type:varchar(100);not null;unique"`
	Deskripsi string    `json:"deskripsi" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
