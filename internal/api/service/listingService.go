package service

import (
	"gifma-backend/config"
	"gifma-backend/internal/api/domain"
	"gifma-backend/internal/helper"
	"gifma-backend/token"

	"github.com/gofiber/fiber/v2"
)

type ListingService struct {
	// Repo       repository.UserRepository
	Config     config.AppConfig
	Crypto     helper.Crypto
	TokenMaker token.Maker
}

func (s *ListingService) CreateListing(c *fiber.Ctx) error {
	return nil
}

func (s *ListingService) GetAllListing() ([]*domain.Listing, error) {
	// TODO get all listing]
	sampleListing := []*domain.Listing{
		{
			ID:      123,
			BrandNo: 5,
			Amount:  10000,
		},
		{
			ID:      124,
			BrandNo: 4,
			Amount:  2000,
		},
	}
	return sampleListing, nil
}
