package handler

import (
	"github.com/gofiber/fiber/v2"
	"go-middleware-recipe/middleware"
	"go-middleware-recipe/repository"
)

type BaseEntityHandler[T any] interface {
	RegisterAll(app *fiber.App)
	Create(app fiber.Router)
	Read(app fiber.Router)
	Update(app fiber.Router)
	Delete(app fiber.Router)
}

type Impl[T any] struct {
	Path string
}

func (c *Impl[T]) RegisterAll(app *fiber.App) {
	userRoute := app.Group(c.Path, middleware.Protected())

	c.Create(userRoute)
	c.Read(userRoute)
}

func (c *Impl[T]) Create(app fiber.Router) {
	app.Post("/", func(ctx *fiber.Ctx) error {
		var requestBody T

		if err := ctx.BodyParser(&requestBody); err != nil {
			return err
		}

		r := repository.Impl[T]{}
		result, err := r.Create(requestBody)

		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return ctx.Status(fiber.StatusCreated).JSON(result)
	})
}

func (c *Impl[T]) Read(app fiber.Router) {
	app.Get("/:id", func(c *fiber.Ctx) error {
		idParam, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Error{Code: fiber.StatusBadRequest, Message: err.Error()})
		}

		r := repository.Impl[T]{}
		user, err := r.Read(idParam)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Error{
				Code:    fiber.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		return c.JSON(user)
	})
}
