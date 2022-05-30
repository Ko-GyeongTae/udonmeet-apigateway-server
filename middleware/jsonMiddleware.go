package middleware

import "github.com/gofiber/fiber/v2"

func JSONMiddleware(c *fiber.Ctx) error {
	c.Response().Header.Set("Content-Type", "application/json")
	c.Response().Header.Set("Access-Control-Allow-Origin", "*")
	return c.Next()
}