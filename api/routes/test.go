package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/abhinavv9/url-shortner/database"
)

func Test(c *fiber.Ctx) error {

	r2 := database.CreateClient(0)
	defer r2.Close()
	r2.Set(database.Ctx, "abc", "efg", 3600)
	val, err := r2.Get(database.Ctx, "abc").Result()
	fmt.Println(val, err)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": "True"})
}
