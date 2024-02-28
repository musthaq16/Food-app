package controller

import (
	"fmt"
	"food-app/data/response"
	"food-app/helper"
	"food-app/model"
	"food-app/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepository repository.UserRepository
}

func NewUserController(repository repository.UserRepository) *UserController {
	return &UserController{
		userRepository: repository,
	}
}

func (controller *UserController) GetAllUsers(ctx *gin.Context) {
	users := controller.userRepository.FindAll()
	resp := response.WebResponse{
		Code:    200,
		Status:  "OK",
		Message: "The all Available users are fetched.",
		Data:    users,
	}
	ctx.JSON(http.StatusOK, resp)
}
func (controller *UserController) UpdateMe(ctx *gin.Context) {

	UpdateUserRequest := model.User{}
	err := ctx.ShouldBindJSON(&UpdateUserRequest)
	helper.ErrorPanic(err)
	var id interface{}
	id, _ = ctx.Get("currentId")
	currentId, _ := strconv.Atoi(fmt.Sprint(id))
	fmt.Println(currentId)

	prevUpdated, err := controller.userRepository.FindById(currentId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Status": "NOT OKAY", "Message": "The user is not vailable on db."})
	}
	fmt.Println("prev Updated::::", prevUpdated.Id)
	controller.userRepository.Update(UpdateUserRequest, prevUpdated)
	// fmt.Println(UpdateUserRequest)
	resp := response.WebResponse{
		Code:    200,
		Status:  "OK",
		Message: "Updated Successfully",
	}
	ctx.JSON(http.StatusOK, resp)
}

func (controller *UserController) DeleteUser(ctx *gin.Context) {

	var id interface{}
	id, _ = ctx.Get("currentId")
	currentId, _ := strconv.Atoi(fmt.Sprint(id))
	fmt.Println(currentId)

	controller.userRepository.Delete(currentId)

	resp := response.WebResponse{
		Code:    200,
		Status:  "OK",
		Message: "Deleted your Profile Successfully",
	}
	ctx.JSON(http.StatusOK, resp)
}
