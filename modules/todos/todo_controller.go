package todos

import (
	"fmt"

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

func GetTodo(c fiber.Ctx) error {
	todo, err := GetTodoService(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(todo)
}

func GetTodos(c fiber.Ctx) error {
	todos, err := GetTodosService(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(todos)
}

func UpdateTodo(c fiber.Ctx) error {
	id, err := UpdateTodoService(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fmt.Sprintf("Todo with id %d was correctly updated", id))
}

func DeleteTodo(c fiber.Ctx) error {
	id, err := DeleteTodoService(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fmt.Sprintf("Todo with id %d was correctly updated", id))
}
