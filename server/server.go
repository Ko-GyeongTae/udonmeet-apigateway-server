package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/udonmeet-apigateway-server/utils"
)

func Start() {
	// Basic Setting of server
	app := fiber.New()

	// Middleware setting
	app.Use(cors.New())
	app.Use(pprof.New())
	app.Use(recover.New())

	//app.Use(csrf.New(utils.Csrf()))
	app.Use(limiter.New(utils.Limiter()))
	app.Use(logger.New(utils.ConsoleLogger()))

	//Routing
	app.Get("/status", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"isActive": true,
		})
	})
	log.Fatal(app.Listen(":" + fmt.Sprint(8080)))
}