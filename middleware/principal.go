package middleware

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

const PrincipalCtxKey = "RequestPrincipal"

type Principal struct {
	context.Context
	*fiber.Ctx
	UserID uint
	Email  string
}

func PrincipalContext() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims := c.Locals(PrincipalCtxKey)

		claimsMap := claims.(*jwt.Token)
		userID := uint(claimsMap.Claims.(jwt.MapClaims)["user_id"].(float64))
		email := claimsMap.Claims.(jwt.MapClaims)["email"].(string)

		customCtx := Principal{
			Ctx:    c,
			UserID: userID,
			Email:  email,
		}

		c.Locals(PrincipalCtxKey, customCtx)
		return c.Next()
	}
}
