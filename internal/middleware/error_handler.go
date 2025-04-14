package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// ErrorHandler middleware to handle errors
func ErrorHandler(c *fiber.Ctx) error {
	// Example error handling logic
	err := c.Next() // Proceed to the next middleware or route handler
	if err != nil {
		// Handle the error and send a response
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return nil
}
