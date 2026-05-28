package todos

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5"
)

func CreateTodo(conn *pgx.Conn) fiber.Handler {
	return func(c fiber.Ctx) error {
		todo, err := CreateTodoService(c, conn)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(todo)
	}
}

func GetTodo(conn *pgx.Conn) fiber.Handler {
	return func(c fiber.Ctx) error {
		todo, err := GetTodoService(c, conn)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(todo)
	}
}

func GetTodos(conn *pgx.Conn) fiber.Handler {
	return func(c fiber.Ctx) error {
		todos, err := GetTodosService(c, conn)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(todos)
	}
}

func UpdateTodo(conn *pgx.Conn) fiber.Handler {
	return func(c fiber.Ctx) error {
		todo, err := UpdateTodoService(c, conn)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(todo)
	}
}

func DeleteTodo(conn *pgx.Conn) fiber.Handler {
	return func(c fiber.Ctx) error {
		err := DeleteTodoService(c, conn)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON("Deleted Todo")
	}
}
