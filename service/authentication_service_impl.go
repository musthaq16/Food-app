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
	"net/smtp"

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
		fmt.Println("in impllll")
		return "", errors.New("invalid username or password")
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
func (a *AuthenticationServiceImpl) Register(user request.CreateUserRequest) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	helper.ErrorPanic(err)
	newUser := model.User{
		UserName: user.Username,
		Email:    user.Email,
		Password: hashedPassword,
	}
	ifAlready_error := a.UserRepository.Save(newUser)

	return ifAlready_error

}

func (a *AuthenticationServiceImpl) ForgetPassword(user request.ForgetPasswordRequest) error {
	//Find the username in db
	new_user, user_err := a.UserRepository.FindByEmail(user.Email)
	if user_err != nil {
		return errors.New("The email is Not Registered")

	}
	config, _ := config.LoadConfig(".")

	otpPassword := utils.OtpGenerate()

	auth := smtp.PlainAuth("", config.UserMailID, config.MailPassword, "smtp.gmail.com")
	to := []string{user.Email}
	body := fmt.Sprintf("Copy the following otp to reset your password: %s", otpPassword)
	msg := []byte(
		"Subject: Password Reset\n\n" + body)

	err := smtp.SendMail("smtp.gmail.com:587", auth, config.UserMailID, to, msg)
	if err != nil {
		return errors.New("Failed to send email.")
	}
	OtpModel := model.OneTimePassword{
		Otp:    otpPassword,
		UserId: new_user.Id,
	}
	saveOtp_err := a.UserRepository.SaveOtp(OtpModel)

	return saveOtp_err

}

func (a *AuthenticationServiceImpl) ResetPassword(resetPasswordReq request.ResetPasswordRequest, otpToken string) error {

	//verifying the otpPassword
	otpModel, err := a.UserRepository.OtpVerifyInToken(otpToken)
	if err != nil {
		return err
	}

	//Check the given two password are same.
	if resetPasswordReq.Password != resetPasswordReq.ConfirmPassword {
		return errors.New("The Confirm Password is not same")
	}
	//get the previous user to save the hashed password in table
	prevUser, err := a.UserRepository.FindById(otpModel.UserId)
	helper.ErrorPanic(err)
	newUser := model.User{
		Password: resetPasswordReq.ConfirmPassword,
	}
	a.UserRepository.Update(newUser, prevUser)
	helper.ErrorPanic(err)

	err = a.UserRepository.DeleteOtp(otpModel.Id)

	helper.ErrorPanic(err)

	return nil
}
