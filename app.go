package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go-middleware-recipe/config"
	"go-middleware-recipe/database"
	"go-middleware-recipe/router"
)

func main() {
	config.CheckApplicationProfile()

	database.ConnectDatabase()

	app := fiber.New()
	app.Use(cors.New())
	router.SetupRoutes(app)

	if err := app.Listen(":8080"); err != nil {
		panic("Error starting the server")
	}
}
