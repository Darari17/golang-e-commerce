package service

import (
	"errors"

	"github.com/Darari17/golang-e-commerce/dto"
	"github.com/Darari17/golang-e-commerce/helper"
	"github.com/Darari17/golang-e-commerce/model"
	"github.com/Darari17/golang-e-commerce/repository"
	"github.com/go-playground/validator/v10"
)

type IUserService interface {
	Register(input dto.RegisterRequest) (dto.UserResponse, error)
	Profile(userId uint) (dto.UserResponse, error)
}

type userService struct {
	userRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &userService{
		userRepository: userRepository,
	}
}

// Profile implements IUserService.
func (u *userService) Profile(userId uint) (dto.UserResponse, error) {
	if userId == 0 {
		return dto.UserResponse{}, errors.New("invalid user id")
	}

	user, err := u.userRepository.FindUserById(userId)
	if err != nil {
		return dto.UserResponse{}, err
	}

	response := dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	return response, nil
}

// Register implements IUserService.
func (u *userService) Register(input dto.RegisterRequest) (dto.UserResponse, error) {
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		return dto.UserResponse{}, err
	}

	hashedPassword, err := helper.HashPassword(input.Password)
	if err != nil {
		return dto.UserResponse{}, err
	}

	register, err := u.userRepository.CreateUser(&model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
	})

	if err != nil {
		return dto.UserResponse{}, err
	}

	response := dto.UserResponse{
		ID:        register.ID,
		Name:      register.Name,
		Email:     register.Email,
		CreatedAt: register.CreatedAt,
	}

	return response, nil
}
