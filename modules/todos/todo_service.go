package todos

import (
	"errors"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5"
)

/*
	TODO:

	- [ ] Needs some kind of connection to database
	- [ ] Needs to grab things from the request
	- [ ] Needs to implement the basic CRUD
*/

func CreateTodoService(c fiber.Ctx, conn *pgx.Conn) (Todo, error) {

	if !c.HasBody() {
		return Todo{}, errors.New("no body has been sent")
	}

	todo := new(Todo)

	if err := c.Bind().Body(todo); err != nil {
		return Todo{}, err
	}

	return createTodoRepository(conn, *todo)
}

func GetTodoService(c fiber.Ctx, conn *pgx.Conn) (Todo, error) {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Fatal("failed to convert string to integer", err)
	}

	return getTodoRepository(conn, id)
}

func GetTodosService(c fiber.Ctx, conn *pgx.Conn) ([]Todo, error) {
	return getTodosRepository(conn)
}

func UpdateTodoService(c fiber.Ctx, conn *pgx.Conn) (Todo, error) {
	if !c.HasBody() {
		return Todo{}, errors.New("no body has been sent")
	}

	todo := new(Todo)

	if err := c.Bind().Body(todo); err != nil {
		return Todo{}, err
	}

	return updateTodoRepository(conn, *todo)
}

func DeleteTodoService(c fiber.Ctx, conn *pgx.Conn) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Fatal("failed to convert string to integer", err)
	}

	return deleteTodoRepository(conn, id)
}
