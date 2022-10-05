package main

import (
	"fmt"
	"log"
	"os"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.app) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}

	app := fiber.New()

	app.Use(logger.New())

	setupRoutes(app)

	log.Fatal(app.Listern(os.Getenv("APP_PORT")))
}