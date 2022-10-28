package routes

import (
	"github.com/abhinavv9/url-shortner/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func ResolveURL(c *fiber.Ctx) error {

	url := c.Params("url")

	r := database.CreateClient(1)
	defer r.Close()

	val, err := r.Get(database.Ctx, url).Result()

	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Short not found"})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}
	rInr := database.CreateClient(1)
	defer rInr.Close()
	_ = rInr.Incr(database.Ctx, "counter")

	return c.Redirect(val, 301)
}
