package main

import (
	"log"
	"simple-fiber-server/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Static("/static", "./public")
	routes.Register(app)

	log.Println("Server is running on http://localhost:3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
