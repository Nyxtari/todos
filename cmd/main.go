package main

import (
	"context"
	"fmt"
	"todos/config"
	"todos/modules/todos"
	"todos/modules/user"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

/*
	TODO:
	- [X] Create config for server that "reads" env variables
	- [X] Create endpoints for todos - CRUD
	- [X] Create module of TODO
	- [X] Create Services for the TODO
	- [X] Link to postgres
		- This is using prisma, in future use my own db please
	- [ ] Create tests
*/

func main() {
	err := godotenv.Load()
	conn, err := config.Connect()
	if err != nil {
		fmt.Println("PROBLEM CONNECTING :: " + err.Error())
		return
	}
	app := fiber.New()

	app.Post("/todos", todos.CreateTodo(conn))
	app.Get("/todo/:id", todos.GetTodo(conn))
	app.Get("/todos", todos.GetTodos(conn))
	app.Put("/todos", todos.UpdateTodo(conn))
	app.Delete("/todos/:id", todos.DeleteTodo(conn))

	app.Get("users", user.GetUsers(conn))

	app.Listen(":3000")

	defer conn.Close(context.Background())
}
