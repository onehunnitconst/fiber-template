package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/onehunnitconst/fiber-template/dto"
	"github.com/onehunnitconst/fiber-template/services"
	"gorm.io/gorm"
)

func TodoRoute(app *fiber.App, db *gorm.DB) {
	app.Post("/todos", func(ctx fiber.Ctx) error {
		form := new(dto.CreateTodoDTO)

		if err := ctx.Bind().Body(form); err != nil {
			return ctx.Status(400).SendString(err.Error())
		}

		services.CreateTodo(*form, db);
		return ctx.SendStatus(201);
	})

	app.Get("/todos", func(ctx fiber.Ctx) error {
		todos := services.GetTodos(db);
		return ctx.JSON(todos);
	})
}