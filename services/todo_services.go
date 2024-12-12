package services

import (
	"github.com/onehunnitconst/fiber-template/dto"
	"github.com/onehunnitconst/fiber-template/models"
	"gorm.io/gorm"
)

func CreateTodo(form dto.CreateTodoDTO, db *gorm.DB) models.Todo {
	todo := models.Todo{
		Title: form.Title,
		Completed: form.Completed,
	}
	db.Create(&todo)
	return todo;
}

func GetTodos(db *gorm.DB) []models.Todo {
	var todos []models.Todo
	db.Find(&todos)
	return todos;
}