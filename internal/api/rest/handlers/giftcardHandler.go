package handlers

import (
	"gifma-backend/internal/api/rest"
	"gifma-backend/internal/api/service"

	"github.com/gofiber/fiber/v2"
)

type GiftcardHandler struct {
	svc service.GiftcardService
}

func SetupGiftcardRoutes(rh *rest.RestHandler) {
	app := rh.App

	// create an instance of giftcard service & inject to handler
	svc := service.GiftcardService{
		// Repo:       repository.NewAdminRepository(rh.DB, rh.Config),
		Config:     rh.Config,
		Crypto:     rh.Crypto,
		TokenMaker: rh.TokenMaker,
	}
	handler := GiftcardHandler{
		svc: svc,
	}

	app.Post("/giftcard", handler.CreateGiftcard)
}

func (h *GiftcardHandler) CreateGiftcard(c *fiber.Ctx) error {
	return h.svc.CreateGiftcard(c)
}
