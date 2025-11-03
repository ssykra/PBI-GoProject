package services

import (
	"authentication/models"
	"authentication/repositories"
	"errors"
)

type ProdukService struct {
	Repo *repositories.ProdukRepository
}

func NewProdukService(repo *repositories.ProdukRepository) *ProdukService {
	return &ProdukService{Repo: repo}
}

func (s *ProdukService) CreateProduk(produk *models.Produk, idToko uint) error {
	produk.IDToko = idToko
	return s.Repo.Create(produk)
}

func (s *ProdukService) UpdateProduk(produkID uint, idToko uint, updateData map[string]interface{}) error {
	produk, err := s.Repo.GetByID(produkID, idToko)
	if err != nil {
		return errors.New("produk tidak ditemukan atau bukan milik Anda")
	}
	return s.Repo.Update(produk, updateData)
}

func (s *ProdukService) GetProduk(produkID uint, idToko uint) (*models.Produk, error) {
	return s.Repo.GetByID(produkID, idToko)
}

func (s *ProdukService) GetAllProduk(idToko uint) ([]models.Produk, error) {
	return s.Repo.GetAllByToko(idToko)
}
