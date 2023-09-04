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
	checkAppProfile()

	database.ConnectDatabase()

	app := fiber.New()
	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	log.Fatal(app.Listen(":8080"))
}

func checkAppProfile() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./your_app <profile>")
		os.Exit(1)
	}

	profile := os.Args[1]

	switch profile {
	case config.Local, config.Test:
		if err := loadProfileEnvironments(profile); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	default:
		fmt.Println("Invalid profile. Available profiles: ", config.Profiles)
		os.Exit(1)
	}
}

func loadProfileEnvironments(profile string) error {
	envFile := fmt.Sprintf("config/%s.env", profile)
	return godotenv.Load(envFile)
}
