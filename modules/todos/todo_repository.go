package todos

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func getTodosRepository(conn *pgx.Conn) ([]Todo, error) {
	query := `
        SELECT id, isDone, title, description, startDate, endDate, difficulty FROM app.todos
    `
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		log.Printf("Error Querying the Table")
		return []Todo{}, err
	}
	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.Id, &todo.IsDone, &todo.Title, &todo.Description, &todo.StartDate, &todo.EndDate, &todo.Difficulty)
		if err != nil {
			log.Printf("error fetching Todos details - %s", err.Error())
			return []Todo{}, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}
