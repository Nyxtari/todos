package main

import (
	"todos/modules/todos"

	"github.com/gofiber/fiber/v3"
)

/*
	TODO:
	- [ ] Create config for server that "reads" env variables
	- [ ] Create endpoints for todos - CRUD
	- [ ] Create module of TODO
	- [ ] Link to postgres
	- [ ] Create tests
*/

func main() {
	app := fiber.New()

	app.Post("/todos", todos.CreateTodo)
	app.Get("/todos/:id", todos.GetTodo)
	app.Get("/todos", todos.GetTodos)
	app.Put("/todos/:id", todos.UpdateTodo)
	app.Delete("/todos/:id", todos.DeleteTodo)

	app.Listen(":3000")
}
