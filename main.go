package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(
		logger.New(),
	)

	// Add route for send message to Service 1.
	app.Get("/status", func(c *fiber.Ctx) error {
		
		return nil
	})

	// Start Fiber API server.
	log.Fatal(app.Listen(":3000"))
}