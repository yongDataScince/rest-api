package service

import (
	"github.com/yongDataScince/rest-api"
	"github.com/yongDataScince/rest-api/pkg/repository"
)

type Autorization interface {
	CreateUser(user rest.User) (int, error) 
	CreateTocken(username, password string) (string, error) 
}

type TodoList interface {

}

type TodoItem interface {

}

type Service struct {
	Autorization
	TodoList
	TodoItem
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Autorization: NewAuthService(r.Autorization),
	}
}

