package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"os"
)

func Protected() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:     jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		ErrorHandler:   jwtError,
		ContextKey:     PrincipalCtxKey,
		SuccessHandler: PrincipalContext(),
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if c.GetReqHeaders()["Authorization"] == "" {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Error{Code: fiber.StatusUnauthorized, Message: err.Error()})
	}

	c.Status(fiber.StatusBadRequest)
	return c.JSON(fiber.Error{Code: fiber.StatusBadRequest, Message: err.Error()})
}
