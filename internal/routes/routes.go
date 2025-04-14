package routes

import (
	"github.com/gofiber/fiber/v2"
	"dummyImplementation/internal/handlers"
)

func SetupRoutes(app *fiber.App) {
	exampleHandler := handlers.ExampleHandler{}

	app.Get("/example", exampleHandler.HandleExample)
}