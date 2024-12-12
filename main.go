package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/onehunnitconst/fiber-template/routes"
	"github.com/onehunnitconst/fiber-template/models"
)

func main() {
	app := fiber.New()

	db := models.Connect();

	routes.TodoRoute(app, db)

	log.Fatal(app.Listen(":3000"))
}