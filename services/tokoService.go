package services

import (
	"authentication/models"
	"authentication/repositories"
	"fmt"
	"strings"
)

type TokoService struct {
	TokoRepo *repositories.TokoRepository
}

func NewTokoService(tokoRepo *repositories.TokoRepository) *TokoService {
	return &TokoService{TokoRepo: tokoRepo}
}

// Otomatis buat toko saat user register
func (s *TokoService) CreateDefaultTokoForUser(userID uint, namaUser string) error {
	namaToko := fmt.Sprintf("Toko %s", namaUser)
	urlToko := fmt.Sprintf("https://toko.com/%s", strings.ToLower(strings.ReplaceAll(namaUser, " ", "_")))

	toko := &models.Toko{
		UserID:   userID, 
		NamaToko: namaToko,
		UrlToko:  urlToko,
	}

	return s.TokoRepo.CreateToko(toko)
}

// Ambil toko milik user login
func (s *TokoService) GetTokoByUserID(userID uint) (*models.Toko, error) {
	return s.TokoRepo.GetTokoByUserID(userID)
}

// Update toko milik user
func (s *TokoService) UpdateToko(userID uint, updatedData models.Toko) (*models.Toko, error) {
	toko, err := s.TokoRepo.GetTokoByUserID(userID)
	if err != nil {
		return nil, err
	}

	toko.NamaToko = updatedData.NamaToko
	toko.UrlToko = updatedData.UrlToko

	if err := s.TokoRepo.UpdateToko(toko); err != nil {
		return nil, err
	}

	return toko, nil
}

// Admin: ambil semua toko (pagination)
func (s *TokoService) GetAllToko(limit, offset int) ([]models.Toko, error) {
	return s.TokoRepo.GetAllToko(limit, offset)
}
