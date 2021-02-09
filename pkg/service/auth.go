package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/yongDataScince/rest-api"
	"github.com/yongDataScince/rest-api/pkg/repository"
)

const (
	salt = "jhdn3jnjkd$&lkma23"
	tockenTTL = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

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

func (s *AuthService) CreateTocken(username, password string) (string, error) {
	// get user from db
	user, err := s.repo.GetUser(username, genHashPassword(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims {
			ExpiresAt: time.Now().Add(tockenTTL).Unix(),
			IssuedAt: time.Now().Unix(),
		},
		user.ID,
	})

	signKey := []byte("kdnjsndjnd*jdnj212md")
	
	tkn, err := token.SignedString(signKey)
	return tkn, err
}

func genHashPassword(password string) string {
	hash_pass := sha1.New()
	hash_pass.Write([]byte(password))

	return fmt.Sprintf("%x", hash_pass.Sum([]byte(salt)))
}