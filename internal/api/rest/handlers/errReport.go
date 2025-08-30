package handlers

import (
	"gifma-backend/internal/api/rest"
	"gifma-backend/internal/api/service"

	"github.com/gofiber/fiber/v2"
)

type ErrReportHandler struct {
	svc service.ErrReportService
}

func SetupErrReportRoutes(rh *rest.RestHandler) {
	app := rh.App

	// create an instance of errReport service & inject to handler
	svc := service.ErrReportService{
		// Repo:       repository.NewAdminRepository(rh.DB, rh.Config),
		Config:     rh.Config,
		Crypto:     rh.Crypto,
		TokenMaker: rh.TokenMaker,
	}
	handler := ErrReportHandler{
		svc: svc,
	}

	app.Post("/errReport", handler.CreateErrReport)
}

func (h *ErrReportHandler) CreateErrReport(c *fiber.Ctx) error {
	return h.svc.CreateErrReport(c)
}
