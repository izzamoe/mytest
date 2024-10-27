package middlewares

import (
	"testku/entities"
	"testku/helpers"

	"github.com/gofiber/fiber/v2"
)

func AdminOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract the role from the context
		role := c.Locals("role")

		// Check if the role is "admin"
		if role != entities.RoleAdmin {
			return helpers.ErrorResponse(c, "Forbidden", fiber.StatusForbidden, nil)
		}

		// Continue to the next middleware or handler
		return c.Next()
	}
}
