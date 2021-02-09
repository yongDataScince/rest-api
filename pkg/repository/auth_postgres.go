package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/yongDataScince/rest-api"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user rest.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Usename, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}