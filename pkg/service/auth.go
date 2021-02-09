package service

import "github.com/yongDataScince/rest-api/pkg/repository"

type AuthService struct {
	repo repository.Repository
}

func NewAuthService(repo repository.Repository) *AuthService {
	return &AuthService{}
}

func (s *AuthService) CreateUser(user rest.User) (int, error) {
	return s.repo.CreateUser(user)
}
