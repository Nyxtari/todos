package todos

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5"
)

/*
	TODO:

	- [ ] Needs some kind of connection to database
	- [ ] Needs to grab things from the request
	- [ ] Needs to implement the basic CRUD
*/

func CreateTodoService(c fiber.Ctx) (Todo, error) {
	return Todo{
		Title: "example",
	}, nil
}

func GetTodoService(c fiber.Ctx) (Todo, error) {
	return Todo{}, nil
}

func GetTodosService(c fiber.Ctx, conn *pgx.Conn) ([]Todo, error) {
	return getTodosRepository(conn)
}

func UpdateTodoService(c fiber.Ctx) (int32, error) {
	return 1, nil
}

func DeleteTodoService(c fiber.Ctx) (int32, error) {
	return 1, nil
}
