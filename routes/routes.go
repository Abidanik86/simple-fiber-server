package routes

import (
	"simple-fiber-server/handlers"

	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	// app.Get("/", handlers.Home)

	// Query Parameters example
	app.Get("/", handlers.Greet)
	app.Get("/hello", handlers.Hello)
	app.Get("/users/:id", handlers.GetUser)
	app.Post("/users", handlers.CreateUser)
	app.Put("/users/:id", handlers.UpdateUser)
	app.Delete("/users/:id", handlers.DeleteUser)

}
