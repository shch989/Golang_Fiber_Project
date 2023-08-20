package middleware

import "github.com/gofiber/fiber/v2"

func Auth(c *fiber.Ctx) error {
	token := c.Get("x-token")
	if token == "" || token != "secret" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	return c.Next()
}

func PermissionCreate(c *fiber.Ctx) error {
	return c.Next()
}
