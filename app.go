package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go-middleware-recipe/config"
	"go-middleware-recipe/database"
	"log"
	"os"
)

func main() {
	checkApplicationProfile()

	database.ConnectDatabase()

	app := fiber.New()
	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	_ = app.Listen(":8080")
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
