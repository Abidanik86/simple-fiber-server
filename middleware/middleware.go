package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"simple-fiber-server/config"
)

func Register(app *fiber.App, cfg *config.Config) {

	// Recover middleware to handle panics
	app.Use(recover.New())

	// Logger middleware to log requests
	app.Use(logger.New())

	// CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: cfg.CORSAllowedOrigins,
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))
}
