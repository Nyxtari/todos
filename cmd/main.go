package main

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

/*
	TODO:
	- [ ] Create config for server that "reads" env variables
	- [ ] Create endpoints for todos - CRUD
	- [ ] Create module of TODO
	- [ ] Link to postgres
	- [ ] Create tests
*/

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		fmt.Printf("what is happening!!")
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
