package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go-middleware-recipe/config"
	"go-middleware-recipe/controller"
	"go-middleware-recipe/database"
	"go-middleware-recipe/model/core"
	"log"
	"os"
)

func main() {
	checkApplicationProfile()

	database.ConnectDatabase()

	app := fiber.New()
	app.Use(cors.New())

	c := controller.Impl[core.User]{Path: "/user"}
	c.RegisterAll(app)

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
