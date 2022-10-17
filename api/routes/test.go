package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Test(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": "True"})
}
