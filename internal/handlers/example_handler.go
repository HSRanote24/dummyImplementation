package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type ExampleHandler struct{}

func (h *ExampleHandler) HandleExample(c *fiber.Ctx) error {
	return c.SendString("Hello from ExampleHandler!")
}
