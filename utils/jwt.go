package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

// GetToken function
func GetTokenString(c *fiber.Ctx) ([]byte, error) {
	// Get token from request token
	jwt := c.Request().Header.Peek("Authorization")

	// Token length validation
	if len(jwt) == 0 {
		c.SendStatus(401)
		return nil, errors.New("token cannot found")
	}

	// Return token with type []byte
	return jwt, nil
}
