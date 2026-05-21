package user

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5"
)

func GetUsers(conn *pgx.Conn) fiber.Handler {
	return func(c fiber.Ctx) error {
		users, err := getUsersService(c, conn)
		if err != nil {
			fmt.Printf("WHAT RIC....")
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(users)
	}
}
