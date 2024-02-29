package service

import (
	"food-app/data/request"
)

type AuthenticationService interface {
	LogIn(user request.LoginRequest) (string, error)
	Register(user request.CreateUserRequest) error
	ForgetPassword(user request.ForgetPasswordRequest) error
	ResetPassword(user request.ResetPasswordRequest, otpToken string) error
}
