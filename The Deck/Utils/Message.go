package utils

import "github.com/gofiber/fiber/v2"

func MessageJSON(c *fiber.Ctx, code int, message string, status string) error {
	return c.Status(code).JSON(fiber.Map{"status": status, "message": message})
}
