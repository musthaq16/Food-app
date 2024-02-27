package repository

import (
	"errors"
	"fmt"
	"food-app/data/request"
	"food-app/helper"
	"food-app/model"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepositoryImpl(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: Db}
}

// Delete implements UserRepository.
func (u *UserRepositoryImpl) Delete(userId int) {
	var user model.User
	result := u.DB.Where("id = ?", userId).Delete(&user)
	helper.ErrorPanic(result.Error)
}

// FindAll implements UserRepository.
func (u *UserRepositoryImpl) FindAll() []model.User {
	var users []model.User
	result := u.DB.Find(&users)
	helper.ErrorPanic(result.Error)
	return users
}

// FindById implements UserRepository.
func (u *UserRepositoryImpl) FindById(userId int) (model.User, error) {
	var user model.User

	/* result := u.DB.Where("id = ?", userId).First(&user)
	helper.ErrorPanic(result.Error)
	return user, nil */
	//otherwise

	result := u.DB.Find(&user, userId)
	if result != nil {
		return user, nil
	} else {
		return user, errors.New("User is Not found")
	}

}

// FindByUserName implements UserRepository.
func (u *UserRepositoryImpl) FindByUserName(username string) (model.User, error) {
	var user model.User

	result := u.DB.First(&user, "username = ?", username)

	if result.Error != nil {
		fmt.Println("username is not found....")
		return user, errors.New("No username is found ")
	}
	return user, nil

}

// Save implements UserRepository.
func (u *UserRepositoryImpl) Save(user model.User) {
	result := u.DB.Create(&user)
	helper.ErrorPanic(result.Error)
}

// Update implements UserRepository.
func (u *UserRepositoryImpl) Update(user model.User) {
	var updateUser = request.UpdateUserRequest{
		Id:       user.Id,
		Username: user.UserName,
		Email:    user.Email,
		Password: user.Password,
	}
	result := u.DB.Model(&user).Updates(updateUser)
	helper.ErrorPanic(result.Error)
}
