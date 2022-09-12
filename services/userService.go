package services

import (
	"app/auth"
	"app/dto"
	"app/models"
	"app/repositories"
)

type UserService struct {
	UserRepository repositories.UserRepositoryDb
}

type UserServiceInterface interface {
	RegisterUser(user dto.UserRequestDto) (models.User, error)
	LoginUser(userDto dto.LoginRequestDto) (string, error)
}

func (h UserService) RegisterUser(userDto dto.UserRequestDto) (models.User, error) {
	user, err := h.UserRepository.Create(userDto)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (h UserService) LoginUser(userDto dto.LoginRequestDto) (string, error) {
	user, err := h.UserRepository.CheckUser(userDto.Email, userDto.Password)

	if err != nil {
		return "", err
	}

	token, err := auth.GenerateJWT(user)

	if err != nil {
		return "", err
	}

	return token, nil
}

func NewUserService(userRepo repositories.UserRepositoryDb) UserService {
	return UserService{UserRepository: userRepo}
}
