package user

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5"
)

func getUsersService(c fiber.Ctx, conn *pgx.Conn) ([]User, error) {

	return getUsersRepository(conn)
}
