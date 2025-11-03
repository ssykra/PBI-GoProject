package repositories

import (
	"authentication/models"
	"gorm.io/gorm"
)

type ProdukRepository struct {
	DB *gorm.DB
}

func NewProdukRepository(db *gorm.DB) *ProdukRepository {
	return &ProdukRepository{DB: db}
}

func (r *ProdukRepository) Create(produk *models.Produk) error {
	return r.DB.Create(produk).Error
}

func (r *ProdukRepository) Update(produk *models.Produk, updateData map[string]interface{}) error {
	return r.DB.Model(produk).Updates(updateData).Error
}

func (r *ProdukRepository) GetByID(produkID uint, idToko uint) (*models.Produk, error) {
	var produk models.Produk
	if err := r.DB.Preload("FotoProduk").First(&produk, "id = ? AND id_toko = ?", produkID, idToko).Error; err != nil {
		return nil, err
	}
	return &produk, nil
}

func (r *ProdukRepository) GetAllByToko(idToko uint) ([]models.Produk, error) {
	var produks []models.Produk
	if err := r.DB.Where("id_toko = ?", idToko).Find(&produks).Error; err != nil {
		return nil, err
	}
	return produks, nil
}
