package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/yongDataScince/rest-api"
)

type Autorization interface {
	CreateUser(user rest.User) (int, error)
}

type TodoList interface {

}

type TodoItem interface {

}

type Repository struct {
	Autorization
	TodoList
	TodoItem
}


func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Autorization: NewAuthPostgres(db),
	}
}