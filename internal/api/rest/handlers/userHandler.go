package handlers

import (
	"gifma-backend/internal/api/dto"
	"gifma-backend/internal/api/rest"
	"gifma-backend/internal/api/service"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	svc service.UserService
}

func SetupUserRoutes(rh *rest.RestHandler) {
	app := rh.App

	// create an instance of user service & inject to handler
	svc := service.UserService{
		// Repo:       repository.NewAdminRepository(rh.DB, rh.Config),
		Config:     rh.Config,
		Crypto:     rh.Crypto,
		TokenMaker: rh.TokenMaker,
	}
	handler := UserHandler{
		svc: svc,
	}

	app.Post("/user/signup", handler.Signup)
	app.Post("/user/login", handler.Login)
}

func (h *UserHandler) Signup(c *fiber.Ctx) error {
	var req dto.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if req.Password != req.ConfirmPassword {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "password and confirm password do not match",
		})
	}

	result, err := h.svc.CreateUser(req.Email, req.Password, req.Role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	rs := dto.CreateUserResponse{
		UserID: result.UserID,
	}
	return c.Status(fiber.StatusOK).JSON(rs)
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req dto.LoginUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	loginUser, err := h.svc.Login(req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	rsp := dto.LoginUserResponse{
		UserID:                loginUser.User.UserID,
		AccessToken:           loginUser.AccessToken,
		AccressTokenExpiresAt: loginUser.AccessTokenExpiresAt.Format(time.RFC3339),
		RefreshToken:          loginUser.RefreshToken,
		RefreshTokenExpiresAt: loginUser.RefreshTokenExpiresAt.Format(time.RFC3339),
	}
	return c.Status(fiber.StatusOK).JSON(rsp)

}
