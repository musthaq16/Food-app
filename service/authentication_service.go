package service

import (
	"food-app/data/request"
)

type AuthenticationService interface {
	LogIn(user request.LoginRequest) (string, error)
	Register(user request.CreateUserRequest) error
}
