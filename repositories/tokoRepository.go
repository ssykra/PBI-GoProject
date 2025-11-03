package repositories

import (
	"authentication/models"
	"gorm.io/gorm"
)

type TokoRepository struct {
	DB *gorm.DB
}

func NewTokoRepository(db *gorm.DB) *TokoRepository {
	return &TokoRepository{DB: db}
}

func (r *TokoRepository) CreateToko(toko *models.Toko) error {
	return r.DB.Create(toko).Error
}

func (r *TokoRepository) GetTokoByUserID(userID uint) (*models.Toko, error) {
	var toko models.Toko
	if err := r.DB.Where("user_id = ?", userID).First(&toko).Error; err != nil {
		return nil, err
	}
	return &toko, nil
}

func (r *TokoRepository) UpdateToko(toko *models.Toko) error {
	return r.DB.Save(toko).Error
}

func (r *TokoRepository) GetAllToko(limit, offset int) ([]models.Toko, error) {
	var tokos []models.Toko
	if err := r.DB.Limit(limit).Offset(offset).Find(&tokos).Error; err != nil {
		return nil, err
	}
	return tokos, nil
}
