package response

import "github.com/gofiber/fiber/v2"

func BadRequestResponse(c *fiber.Ctx) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"status":  "Failed",
		"message": "Invalid Request Data",
	})
}

func NotFoundResponse(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"status":  "Failed",
		"message": "Data Not Found",
	})
}

func InternalServerError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"status":  "Failed",
		"message": "Internal Server Error",
	})
}

func UnauthorizedResponse(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"status":  "Failed",
		"message": "Invalid Credential",
	})
}

func SuccessResponse(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "successfully",
	})
}
