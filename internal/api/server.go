package api

import (
	"gifma-backend/config"
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
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	var err error
	log.Printf("Server starting on port %s", config.ServerAddress)
	err = app.Listen(config.ServerAddress)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
