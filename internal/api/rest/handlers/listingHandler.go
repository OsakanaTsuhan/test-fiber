package handlers

import (
	"gifma-backend/internal/api/rest"
	"gifma-backend/internal/api/service"

	"github.com/gofiber/fiber/v2"
)

type ListingHandler struct {
	svc service.ListingService
}

func SetupListingRoutes(rh *rest.RestHandler) {
	app := rh.App

	// create an instance of listing service & inject to handler
	svc := service.ListingService{
		// Repo:       repository.NewAdminRepository(rh.DB, rh.Config),
		Config:     rh.Config,
		Crypto:     rh.Crypto,
		TokenMaker: rh.TokenMaker,
	}
	handler := ListingHandler{
		svc: svc,
	}

	app.Post("/listing", handler.CreateListing)
}

func (h *ListingHandler) CreateListing(c *fiber.Ctx) error {
	return h.svc.CreateListing(c)
}

func (h *ListingHandler) GetAllListing(c *fiber.Ctx) error {
	result, err := h.svc.GetAllListing()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}
