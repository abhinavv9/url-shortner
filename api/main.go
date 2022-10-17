package main

import (
	"fmt"
	"log"
	"os"

	"github.com/abhinavv9/url-shortner/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
	app.Get("/", routes.Test)
}

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("jdfhjsd")

	app := fiber.New()

	app.Use(logger.New())

	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
