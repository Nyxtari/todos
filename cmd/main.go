package main

import (
	"todos/modules/todos"

	"github.com/gofiber/fiber/v3"
)

/*
	TODO:
	- [ ] Create config for server that "reads" env variables
	- [X] Create endpoints for todos - CRUD
	- [X] Create module of TODO
	- [ ] Create Services for the TODO
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
