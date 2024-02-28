package repository

import (
	"food-app/model"
)

type UserRepository interface {
	Save(user model.User) error
	Update(user model.User)
	Delete(userId int)
	FindById(userId int) (model.User, error)
	FindAll() []model.User
	FindByUserName(username string) (model.User, error)
}
