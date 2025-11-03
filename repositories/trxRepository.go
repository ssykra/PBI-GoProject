package repositories

import (
	"authentication/models"
	"gorm.io/gorm"
)

type TrxRepository struct {
	DB *gorm.DB
}

func NewTrxRepository(db *gorm.DB) *TrxRepository {
	return &TrxRepository{DB: db}
}

func (r *TrxRepository) Create(trx *models.Trx) error {
	return r.DB.Create(trx).Error
}

func (r *TrxRepository) GetByID(trxID uint, userID uint) (*models.Trx, error) {
	var trx models.Trx
	if err := r.DB.Preload("DetailTrx").First(&trx, "id = ? AND id_user = ?", trxID, userID).Error; err != nil {
		return nil, err
	}
	return &trx, nil
}

func (r *TrxRepository) GetAllByUser(userID uint) ([]models.Trx, error) {
	var trxs []models.Trx
	if err := r.DB.Preload("DetailTrx").Where("id_user = ?", userID).Find(&trxs).Error; err != nil {
		return nil, err
	}
	return trxs, nil
}
