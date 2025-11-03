package repositories

import (
	"authentication/models"
	"gorm.io/gorm"
)

type KategoriRepository struct {
	DB *gorm.DB
}

func NewKategoriRepository(db *gorm.DB) *KategoriRepository {
	return &KategoriRepository{DB: db}
}

func (r *KategoriRepository) Create(k *models.Kategori) error {
	return r.DB.Create(k).Error
}

func (r *KategoriRepository) Update(k *models.Kategori) error {
	return r.DB.Save(k).Error
}

func (r *KategoriRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Kategori{}, id).Error
}

func (r *KategoriRepository) GetByID(id uint) (*models.Kategori, error) {
	var kategori models.Kategori
	if err := r.DB.First(&kategori, id).Error; err != nil {
		return nil, err
	}
	return &kategori, nil
}

// Pagination + search
func (r *KategoriRepository) GetAll(page, limit int, search string) ([]models.Kategori, int64, error) {
	var kategoris []models.Kategori
	var total int64

	query := r.DB.Model(&models.Kategori{})
	if search != "" {
		query = query.Where("nama LIKE ?", "%"+search+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := query.Limit(limit).Offset(offset).Find(&kategoris).Error; err != nil {
		return nil, 0, err
	}

	return kategoris, total, nil
}
