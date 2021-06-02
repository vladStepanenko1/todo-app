package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/vladStepanenko1/todo-app"
)

type AuthPostgres struct {
	database *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{database: db}
}

func (ap *AuthPostgres) CreateUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values($1, $2, $3) RETURNING id", usersTable)
	row := ap.database.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
