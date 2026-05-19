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

	app.Get("/", todos.CreateTodo)

	app.Listen(":3000")
}
