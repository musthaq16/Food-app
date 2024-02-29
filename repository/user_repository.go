package repository

import (
	"food-app/model"
)

type UserRepository interface {
	Save(user model.User) error
	Update(user, prevUpdated model.User)
	Delete(userId int)
	FindById(userId int) (model.User, error)
	FindAll() []model.User
	FindByUserName(username string) (model.User, error)
	FindByEmail(email string) (model.User, error)
	SaveOtp(user model.OneTimePassword) error
	DeleteOtp(otp int) error
	OtpVerifyInToken(otp string) (model.OneTimePassword, error)
}
