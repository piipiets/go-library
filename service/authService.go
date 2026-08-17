package service

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/piipiets/go-library/repository"
)

type AuthService struct {
	userRepository *repository.UserRepository
}

func NewAuthService(
	userRepository *repository.UserRepository,
) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

func (s *AuthService) Login(
	username string,
	password string,
) (string, error) {

	user, err := s.userRepository.FindByUsername(username)

	if err != nil {
		return "", err
	}

	if user.Password != password {
		return "", errors.New("invalid username or password")
	}

	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	tokenString, err := token.SignedString(
		[]byte("secret-key"),
	)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
