package api

import (
	"gifma-backend/config"
	"gifma-backend/internal/api/rest"
	"gifma-backend/internal/api/rest/handlers"
	"gifma-backend/internal/helper"
	"gifma-backend/token"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	// cors configuration
	c := cors.New(cors.Config{
		AllowOrigins: "http://localhost:3030",
		AllowHeaders: "Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, PATCH, DELETE, OPTIONS",
	})

	app.Use(c)

	auth := helper.SetupAuth(config.AppSecret)
	crypto := helper.SetupCrypto(config.AppCryptKey)

	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Fatalf("cannot create token maker: %v\n", err)
	}

	rh := &rest.RestHandler{
		App: app,
		// DB:         db,
		Auth:       auth,
		Config:     config,
		Crypto:     crypto,
		TokenMaker: tokenMaker,
	}

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	setupRoutes(rh)

	log.Printf("Server starting on port %s", config.ServerAddress)
	err = app.Listen(config.ServerAddress)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func setupRoutes(rh *rest.RestHandler) {
	// user handler
	handlers.SetupUserRoutes(rh)
}
