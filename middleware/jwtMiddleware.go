package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/udonmeet-apigateway-server/utils"
)

func AuthMiddleware(c *fiber.Ctx) error {

	jwt, err := utils.GetTokenString(c)
	utils.HandleErr(err)

	// _, user, err := utils.ValidateToken(string(jwt), c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "expired token",
		})
	}
	fmt.Println(jwt)
	if err != nil {
		return c.SendStatus(401)
	}
	return c.Next()
}