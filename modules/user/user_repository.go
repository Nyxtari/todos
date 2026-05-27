package user

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func getUsersRepository(conn *pgx.Conn) ([]User, error) {
	query := `
        SELECT id, name, email FROM app.users
    `
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		log.Printf("Error Querying the Table")
		return []User{}, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			log.Printf("error fetching Users details - %s", err.Error())
			return []User{}, err
		}
		users = append(users, user)
	}

	return users, nil
}
