package services

import (
	"authentication/models"
	"authentication/repositories"
)

type KategoriService struct {
	KategoriRepo *repositories.KategoriRepository
}

func NewKategoriService(repo *repositories.KategoriRepository) *KategoriService {
	return &KategoriService{KategoriRepo: repo}
}

func (s *KategoriService) CreateKategori(k *models.Kategori) error {
	return s.KategoriRepo.Create(k)
}

func (s *KategoriService) UpdateKategori(k *models.Kategori) error {
	return s.KategoriRepo.Update(k)
}

func (s *KategoriService) DeleteKategori(id uint) error {
	return s.KategoriRepo.Delete(id)
}

func (s *KategoriService) GetKategoriByID(id uint) (*models.Kategori, error) {
	return s.KategoriRepo.GetByID(id)
}

func (s *KategoriService) GetAllKategori(page, limit int, search string) ([]models.Kategori, int64, error) {
	return s.KategoriRepo.GetAll(page, limit, search)
}
