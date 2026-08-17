package service

import (
	"time"

	"github.com/piipiets/go-library/model"
	"github.com/piipiets/go-library/model/request"
	"github.com/piipiets/go-library/repository"
)

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(
	repository *repository.UserRepository,
) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) CreateUser(userReques request.AddUserRequest) (string, error) {
	var userModel model.Users

	userModel.Username = userReques.Username
	userModel.Password = userReques.Password
	userModel.CreatedAt = time.Now()
	userModel.CreatedBy = "admin"
	userModel.ModifiedAt = time.Now()
	userModel.ModifiedBy = "admin"

	user, err := s.repository.CreateUser(userModel)
	if err != nil {
		return "", err
	}

	return user.Username, nil
}
