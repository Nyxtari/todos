package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"todos/modules/todos"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

/*
	TODO:
	- [ ] Create config for server that "reads" env variables
	- [X] Create endpoints for todos - CRUD
	- [X] Create module of TODO
	- [ ] Create Services for the TODO
	- [X] Link to postgres
		- This is using prisma, in future use my own db please
	- [ ] Create tests
*/

type MyUser struct {
	id    int
	name  string
	email string
}

func main() {
	fmt.Println("APP STARTING")
	err := godotenv.Load()
	conn, err := connect()
	if err != nil {
		return
	}
	defer conn.Close(context.Background())
	fmt.Println("PRIMA CONNECTED")

	queryData(conn)

	app := fiber.New()
	fmt.Println("FIBER STARTING")

	app.Post("/todos", todos.CreateTodo)
	app.Get("/todos/:id", todos.GetTodo)
	app.Get("/todos", todos.GetTodos)
	app.Put("/todos/:id", todos.UpdateTodo)
	app.Delete("/todos/:id", todos.DeleteTodo)

	app.Listen(":3000")

	fmt.Println("Dying")

}

func connect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func queryData(conn *pgx.Conn) {
	query := `
        SELECT id, name, email FROM "User"
    `
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		log.Printf("Error Querying the Table")
		return
	}
	defer rows.Close()

	fmt.Println("ROWS ARE --- %s", rows)

	var users []MyUser

	for rows.Next() {
		var user MyUser
		err := rows.Scan(&user.id, &user.name, &user.email)
		if err != nil {
			log.Printf("Error Fetching Users Details - " + err.Error())
			return
		}
		users = append(users, user)
	}

	fmt.Println(users)

	return
}
