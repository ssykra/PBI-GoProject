package services

import (
	"authentication/models"
	"errors"
	"time"

	"gorm.io/gorm"
)

type TrxService struct {
	DB *gorm.DB
}

// Constructor
func NewTrxService(db *gorm.DB) *TrxService {
	return &TrxService{DB: db}
}

// CreateTrx membuat transaksi beserta detail dan log produk
func (s *TrxService) CreateTrx(trx *models.Trx, userID uint) error {
	trx.IDUser = userID
	trx.CreatedAt = time.Now()
	trx.UpdatedAt = time.Now()

	return s.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(trx).Error; err != nil {
			return err
		}

		for i := range trx.DetailTrx {
			// Ambil produk asli berdasarkan ID produk
			var produk models.Produk
			if err := tx.First(&produk, "id = ?", trx.DetailTrx[i].IDProduk).Error; err != nil {
				return errors.New("produk tidak ditemukan")
			}

			// Simpan snapshot ke log_produk
			log := models.LogProduk{
				NamaProduk:    produk.NamaProduk,
				Slug:          produk.Slug,
				HargaReseller: produk.HargaReseller,
				HargaKonsumen: produk.HargaKonsumen,
				Stok:          produk.Stok,
				IDToko:        produk.IDToko,
				IDKategori:    produk.IDKategori,
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
			}
			if err := tx.Create(&log).Error; err != nil {
				return err
			}

			// Simpan detail transaksi
			trx.DetailTrx[i].IDTrx = trx.ID
			trx.DetailTrx[i].IDLogProduk = log.ID
			trx.DetailTrx[i].IDToko = produk.IDToko
			trx.DetailTrx[i].HargaTotal = trx.DetailTrx[i].Kuantitas * produk.HargaKonsumen // contoh perhitungan
			trx.DetailTrx[i].CreatedAt = time.Now()
			trx.DetailTrx[i].UpdatedAt = time.Now()

			if err := tx.Create(&trx.DetailTrx[i]).Error; err != nil {
				return err
			}

			// Update stok produk
			if produk.Stok < trx.DetailTrx[i].Kuantitas {
				return errors.New("stok produk tidak cukup")
			}
			produk.Stok -= trx.DetailTrx[i].Kuantitas
			if err := tx.Save(&produk).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

// GetTrx mengambil transaksi berdasarkan userID
func (s *TrxService) GetTrx(trxID uint, userID uint) (*models.Trx, error) {
	var trx models.Trx
	if err := s.DB.Preload("DetailTrx").First(&trx, "id = ? AND id_user = ?", trxID, userID).Error; err != nil {
		return nil, errors.New("transaksi tidak ditemukan")
	}
	return &trx, nil
}

// GetAllTrxByUser mengambil semua transaksi milik user
func (s *TrxService) GetAllTrxByUser(userID uint) ([]models.Trx, error) {
	var trxs []models.Trx
	if err := s.DB.Preload("DetailTrx").Where("id_user = ?", userID).Find(&trxs).Error; err != nil {
		return nil, err
	}
	return trxs, nil
}
