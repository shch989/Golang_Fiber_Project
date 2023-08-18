package controllers

import "github.com/gofiber/fiber/v2"

func UserHandlerRead(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"hello": "world!",
	})
}
