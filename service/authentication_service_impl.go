package service

import (
	"errors"
	"fmt"
	"food-app/config"
	"food-app/data/request"
	"food-app/helper"
	"food-app/model"
	"food-app/repository"
	"food-app/utils"

	"github.com/go-playground/validator/v10"
)

type AuthenticationServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewAuthenticationServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) AuthenticationService {
	return &AuthenticationServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

// LogIn implements AuthenticationService.
func (a *AuthenticationServiceImpl) LogIn(user request.LoginRequest) (string, error) {
	//Find the username in db
	new_user, user_err := a.UserRepository.FindByUserName(user.Username)
	if user_err != nil {
		return "", errors.New("invalid username or password.")
	}

	config, _ := config.LoadConfig(".")
	verify_error := utils.VerifyPassword(new_user.Password, user.Password)
	if verify_error != nil {
		return "", fmt.Errorf("the username or password is incorrect")
	}

	token, token_err := utils.GenerateToken(config.TokenExpiresIn, new_user.Id, config.SecretKey)
	helper.ErrorPanic(token_err)
	return token, nil

}

// Register implements AuthenticationService.
func (a *AuthenticationServiceImpl) Register(user request.CreateUserRequest) {
	hashedPassword, err := utils.HashPassword(user.Password)
	helper.ErrorPanic(err)
	newUser := model.User{
		UserName: user.Username,
		Email:    user.Email,
		Password: hashedPassword,
	}
	a.UserRepository.Save(newUser)

}
