package handlers

import (
	"gifma-backend/internal/api/rest"
	"gifma-backend/internal/api/service"

	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	svc service.OrderService
}

func SetupOrderRoutes(rh *rest.RestHandler) {
	app := rh.App

	// create an instance of order service & inject to handler
	svc := service.OrderService{
		// Repo:       repository.NewAdminRepository(rh.DB, rh.Config),
		Config:     rh.Config,
		Crypto:     rh.Crypto,
		TokenMaker: rh.TokenMaker,
	}
	handler := OrderHandler{
		svc: svc,
	}

	app.Post("/order", handler.CreateOrder)
}

func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	return h.svc.CreateOrder(c)
}
