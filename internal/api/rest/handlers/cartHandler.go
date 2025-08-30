package handlers

import (
	"gifma-backend/internal/api/rest"
	"gifma-backend/internal/api/service"

	"github.com/gofiber/fiber/v2"
)

type CartHandler struct {
	svc service.CartService
}

func SetupCartRoutes(rh *rest.RestHandler) {
	app := rh.App

	// create an instance of cart service & inject to handler
	svc := service.CartService{
		// Repo:       repository.NewAdminRepository(rh.DB, rh.Config),
		Config:     rh.Config,
		Crypto:     rh.Crypto,
		TokenMaker: rh.TokenMaker,
	}
	handler := CartHandler{
		svc: svc,
	}

	app.Post("/cart", handler.CreateCart)
}

func (h *CartHandler) CreateCart(c *fiber.Ctx) error {
	return h.svc.CreateCart(c)
}
