package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go-middleware-recipe/config"
	"go-middleware-recipe/database"
	"go-middleware-recipe/model/core"
	"go-middleware-recipe/repository"
	"log"
	"os"
)

func main() {
	checkApplicationProfile()

	database.ConnectDatabase()

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/user", func(c *fiber.Ctx) error {
		userRepository := repository.Impl[core.User]{}

		user := core.User{Email: "hello@berkan.io"}
		user, err := userRepository.Create(user)

		err = user.SetPassword("123456")

		if err != nil {
			return err
		}

		if err != nil {
			return c.Status(500).JSON(fiber.Error{Code: 500, Message: err.Error()})
		}

		return c.JSON(user)
	})

	app.Get("/user/:id", func(c *fiber.Ctx) error {
		userId, err := c.ParamsInt("id")

		userRepository := &repository.Impl[core.User]{}

		if err != nil {
			return c.Status(400).JSON(fiber.Error{Code: 400, Message: "Invalid user id"})
		}

		user, err := userRepository.Read(userId)

		if err != nil {
			return c.Status(500).JSON(fiber.Error{Code: 500, Message: err.Error()})
		}

		return c.JSON(user)
	})

	if err := app.Listen(":8080"); err != nil {
		panic("Error starting the server")
	}
}

func checkApplicationProfile() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./your_app <profile>")
		os.Exit(1)
	}

	profile := config.Profile(os.Args[1])

	switch profile {
	case config.Local, config.Test:
		err := loadProfileEnvironments(profile)

		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}

		config.ActiveProfile = profile
	default:
		fmt.Println("Invalid profile. Available profiles: ", config.Profiles)
		os.Exit(1)
	}
}

func loadProfileEnvironments(profile config.Profile) error {
	envFile := fmt.Sprintf("config/%s.env", profile)
	return godotenv.Load(envFile)
}
