package todos

import (
	"github.com/gofiber/fiber/v3"
)

func CreateTodo(c fiber.Ctx) error {
	todo, err := CreateTodoService(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(todo)
}
