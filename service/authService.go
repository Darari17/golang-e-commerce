package service

import (
	"errors"

	"github.com/Darari17/golang-e-commerce/dto"
	"github.com/Darari17/golang-e-commerce/helper"
	"github.com/Darari17/golang-e-commerce/repository"
	"github.com/Darari17/golang-e-commerce/security"
)

type IAuthService interface {
	Login(input dto.LoginRequest) (string, error)
}

type authService struct {
	userRepository repository.IUserRepository
	jwtHandler     security.IJWTHandler
}

func NewAuthService(userRepository repository.IUserRepository, jwtHandler security.IJWTHandler) IAuthService {
	return &authService{
		userRepository: userRepository,
		jwtHandler:     jwtHandler,
	}
}

// Login implements IAuthService.
func (a *authService) Login(input dto.LoginRequest) (string, error) {
	user, err := a.userRepository.FindUserByEmail(input.Email)
	if err != nil {
		return "", err
	}

	if !helper.CheckPassword(user.Password, input.Password) {
		return "", errors.New("incorrect email or password")
	}

	token, err := a.jwtHandler.CreateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil

}
