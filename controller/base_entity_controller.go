package controller

import (
	"github.com/gofiber/fiber/v2"
	"go-middleware-recipe/repository"
)

type BaseEntityController[T any] interface {
	RegisterAll(app *fiber.App)
	Create(app *fiber.App)
	Read(app *fiber.App)
	Update(app *fiber.App)
	Delete(app *fiber.App)
}

type Impl[T any] struct {
	Path string
}

func (c *Impl[T]) RegisterAll(app *fiber.App) {
	c.Create(app)
	c.Read(app)
}

func (c *Impl[T]) Create(app *fiber.App) {
	app.Post(c.Path, func(ctx *fiber.Ctx) error {
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

func (c *Impl[T]) Read(app *fiber.App) {
	app.Get(c.Path+"/:id", func(c *fiber.Ctx) error {
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
