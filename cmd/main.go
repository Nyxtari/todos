package main

import (
	"context"
	"todos/config"
	"todos/modules/todos"
	"todos/modules/user"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

/*
	TODO:
	- [ ] Create config for server that "reads" env variables
	- [X] Create endpoints for todos - CRUD
	- [X] Create module of TODO
	- [ ] Create Services for the TODO
	- [X] Link to postgres
		- This is using prisma, in future use my own db please
	- [ ] Create tests
*/

func main() {
	err := godotenv.Load()
	conn, err := config.Connect()
	if err != nil {
		return
	}
	app := fiber.New()

	app.Post("/todos", todos.CreateTodo)
	app.Get("/todos/:id", todos.GetTodo)
	app.Get("/todos", todos.GetTodos)
	app.Put("/todos/:id", todos.UpdateTodo)
	app.Delete("/todos/:id", todos.DeleteTodo)

	app.Get("users", user.GetUsers(conn))

	app.Listen(":3000")

	defer conn.Close(context.Background())
}
