package service

import (
	"gifma-backend/config"
	"gifma-backend/internal/helper"
	"gifma-backend/token"

	"github.com/gofiber/fiber/v2"
)

type GiftcardService struct {
	// Repo       repository.UserRepository
	Config     config.AppConfig
	Crypto     helper.Crypto
	TokenMaker token.Maker
}

func (s *GiftcardService) CreateGiftcard(c *fiber.Ctx) error {
	return nil
}
