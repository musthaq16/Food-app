package repository

import (
	"errors"
	"fmt"
	"food-app/helper"
	"food-app/model"
	"food-app/utils"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepositoryImpl(DB *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: DB}
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
	fmt.Println(username)

	result := u.DB.First(&user, u.DB.Where("user_name = ? ", username))

	if result.Error != nil {
		fmt.Println("username is not found....")
		return user, errors.New("No username is found ")
	}
	return user, nil

}

// Save implements UserRepository.
func (u *UserRepositoryImpl) Save(user model.User) error {
	result := u.DB.Create(&user)
	if result.Error != nil {
		fmt.Println("there is error in storing model", result.Error)
		return result.Error
	}
	// helper.ErrorPanic(result.Error)
	return nil

}

// Update implements UserRepository.
func (u *UserRepositoryImpl) Update(user, prevUpdated model.User) {
	if user.Id == 0 {
		user.Id = prevUpdated.Id
	}
	if user.Email == "" {
		user.Email = prevUpdated.Email
	}
	if user.UserName == "" {
		user.UserName = prevUpdated.UserName
	}
	if user.Password != "" {
		user.Password, _ = utils.HashPassword(user.Password)
	}
	if user.Password == "" {
		user.Password = prevUpdated.Password
	}
	// var updateUser = request.UpdateUserRequest{
	// 	Id:       user.Id,
	// 	Username: user.UserName,
	// 	Email:    user.Email,
	// 	Password: user.Password,
	// }
	// fmt.Println("Updated userr", updateUser)
	result := u.DB.Save(&user)
	helper.ErrorPanic(result.Error)
}
