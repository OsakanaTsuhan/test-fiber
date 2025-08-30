package handlers

import (
	"gifma-backend/internal/api/rest"
	"gifma-backend/internal/api/service"

	"github.com/gofiber/fiber/v2"
)

type WithdrawHandler struct {
	svc service.WithdrawService
}

func SetupWithdrawRoutes(rh *rest.RestHandler) {
	app := rh.App

	// create an instance of withdraw service & inject to handler
	svc := service.WithdrawService{
		// Repo:       repository.NewAdminRepository(rh.DB, rh.Config),
		Config:     rh.Config,
		Crypto:     rh.Crypto,
		TokenMaker: rh.TokenMaker,
	}
	handler := WithdrawHandler{
		svc: svc,
	}

	app.Post("/withdraw", handler.CreateWithdraw)
}

func (h *WithdrawHandler) CreateWithdraw(c *fiber.Ctx) error {
	return h.svc.CreateWithdraw(c)
}
