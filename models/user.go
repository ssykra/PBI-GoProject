package models

import "time"

type User struct {
	ID             uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Nama           string    `json:"nama" validate:"required,min=2,max=100"`
	Password       string    `json:"password" validate:"required,min=6"`
	NoTelp         string    `json:"notelp" validate:"required"`
	TanggalLahir   string    `json:"tanggal_lahir" validate:"required"`
	JenisKelamin   string    `json:"jenis_kelamin"`
	Tentang        string    `json:"tentang"`
	Pekerjaan      string    `json:"pekerjaan"`
	Email          string    `json:"email" validate:"email,required" gorm:"unique"`
	IDProvinsi     string    `json:"id_provinsi"`
	IDKota         string    `json:"id_kota"`
	IsAdmin        bool      `json:"is_admin" gorm:"default:false"`
	Token          string    `json:"token,omitempty"`
	RefreshToken   string    `json:"refresh_token,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
