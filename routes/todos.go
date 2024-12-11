package routes

import (
	"github.com/onehunnitconst/fiber-template/dto"
	"github.com/onehunnitconst/fiber-template/services"
	"github.com/gofiber/fiber/v3"
)

func TodoRoute(app *fiber.App) {
	app.Post("/todos", func(ctx fiber.Ctx) error {
		form := new(dto.CreateTodoDTO)

		if err := ctx.Bind().Body(form); err != nil {
			return ctx.Status(400).SendString(err.Error())
		}

		services.CreateTodo(*form);
		return ctx.SendStatus(201);
	})

	app.Get("/todos", func(ctx fiber.Ctx) error {
		todos := services.GetTodos();
		return ctx.JSON(todos);
	})
}