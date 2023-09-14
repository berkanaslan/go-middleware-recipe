package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go-middleware-recipe/config"
	"go-middleware-recipe/database"
	"go-middleware-recipe/router"
	"os"
)

func main() {
	config.CheckApplicationProfile()

	database.ConnectDatabase()
	InitializeData()

	app := fiber.New()
	app.Use(cors.New())
	router.SetupRoutes(app)

	if err := app.Listen(":" + os.Getenv("PORT")); err != nil {
		panic("Error starting the server")
	}
}
