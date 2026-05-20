package todos

import "github.com/gofiber/fiber/v3"

func CreateTodoService(c fiber.Ctx) (Todo, error) {
	return Todo{
		Title: "example",
	}, nil
}

func GetTodoService(c fiber.Ctx) (Todo, error) {
	return Todo{}, nil
}

func GetTodosService(c fiber.Ctx) ([]Todo, error) {
	return []Todo{}, nil
}

func UpdateTodoService(c fiber.Ctx) (int32, error) {
	return 1, nil
}

func DeleteTodoService(c fiber.Ctx) (int32, error) {
	return 1, nil
}
