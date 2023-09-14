package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go-middleware-recipe/model/core"
	"go-middleware-recipe/repository"
	"os"
	"time"
)

func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Identity string `json:"identity"`
		Password string `json:"password"`
	}

	var input LoginInput

	if err := c.BodyParser(&input); err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	identity := input.Identity
	pass := input.Password

	userRepository := repository.NewUserRepository()
	user, err := userRepository.FindByEmail(identity)

	if err != nil {
		return c.JSON(fiber.Error{Code: fiber.StatusUnauthorized, Message: "Invalid user"})
	}

	if !user.IsPasswordValid(pass) {
		return c.JSON(fiber.Error{Code: fiber.StatusUnauthorized, Message: "Invalid password"})
	}

	token := generateJWT(user)

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	c.Set("X-Auth-Token", t)
	return c.JSON(user)
}

func generateJWT(user core.User) *jwt.Token {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	return token
}
