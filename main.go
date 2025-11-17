package main

import (
	"log"
	"simple-fiber-server/middleware"
	"simple-fiber-server/routes"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"

	"simple-fiber-server/config"
)

func main() {

	cfg := config.Load()

	app := fiber.New(fiber.Config{
		AppName: "Simple Fiber Server",
	})

	// Register middleware
	middleware.Register(app, cfg)

	app.Static("/static", "./public")
	routes.Register(app)

	log.Printf("Server running on http://localhost%s (env: %s)\n", cfg.AppPort, cfg.AppEnv)
	if err := app.Listen(cfg.AppPort); err != nil {
		log.Fatal(err)
	}
}
