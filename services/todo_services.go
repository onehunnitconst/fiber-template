package services

import (
	"github.com/onehunnitconst/fiber-template/dto"
	"github.com/onehunnitconst/fiber-template/models"
)

func CreateTodo(form dto.CreateTodoDTO) models.Todo {
	todo := models.Todo{
		Title: form.Title,
		Completed: form.Completed,
	}
	models.Database.Create(&todo)
	return todo;
}

func GetTodos() []models.Todo {
	var todos []models.Todo
	models.Database.Find(&todos)
	return todos;
}