package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/yongDataScince/rest-api"
	"github.com/yongDataScince/rest-api/pkg/repository"
)

const salt = "jhdn3jnjkd$&lkma23"

// AuthService ...
type AuthService struct {
	repo repository.Autorization
}

// NewAuthService ...
func NewAuthService(repo repository.Autorization) *AuthService {
	return &AuthService{repo: repo}
}

// CreateUser ...
func (s *AuthService) CreateUser(user rest.User) (int, error) {
	user.Password = genHashPassword(user.Password)
	
	return s.repo.CreateUser(user)
}

func genHashPassword(password string) string {
	hash_pass := sha1.New()
	hash_pass.Write([]byte(password))

	return fmt.Sprintf("%x", hash_pass.Sum([]byte(salt)))
}