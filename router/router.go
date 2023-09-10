package router

import (
	"github.com/gofiber/fiber/v2"
	"go-middleware-recipe/handler"
	"go-middleware-recipe/model/core"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/login", handler.Login)

	userHandler := handler.Impl[core.User]{Path: "/user"}
	userHandler.RegisterAll(app)
}
