package controller

import (
	"fmt"
	"food-app/data/request"
	"food-app/data/response"
	"food-app/helper"
	"food-app/service"
	"net/http"

	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
	passwordvalidator "github.com/wagslane/go-password-validator"
)

type AuthenticationController struct {
	AuthenticationService service.AuthenticationService
}

func NewAuthenticationController(service service.AuthenticationService) *AuthenticationController {
	return &AuthenticationController{AuthenticationService: service}
}

func (controller *AuthenticationController) LogIn(ctx *gin.Context) {
	loginRequest := request.LoginRequest{}
	err := ctx.ShouldBindJSON(&loginRequest)
	helper.ErrorPanic(err)

	token, err_token := controller.AuthenticationService.LogIn(loginRequest)
	if err_token != nil {
		webResponse := response.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Invalid username or password",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	resp := response.LoginResponse{
		TokenType: "Bearer",
		Token:     token,
	}
	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Succesfully Log in!!",
		Data:    resp,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *AuthenticationController) Register(ctx *gin.Context) {
	createUserRequest := request.CreateUserRequest{}

	err := ctx.ShouldBindJSON(&createUserRequest)
	// helper.ErrorPanic(err)
	if err != nil {
		fmt.Println("errrrorrrrrr", err)
		resp := response.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "NOT OK",
			Message: "Failed to create user.",
		}
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	//email validating
	email_VAlidate_err := checkmail.ValidateFormat(createUserRequest.Email)
	if email_VAlidate_err != nil {
		fmt.Println("wrong email format @...")
		resp := response.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "NOT OK",
			Message: "Mail Format is Wrong...",
		}
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	password_Validate_err := passwordvalidator.Validate(createUserRequest.Password, 50)
	if password_Validate_err != nil {
		fmt.Println("Password is not strong..")
		resp := response.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "NOT OK",
			Message: fmt.Sprintf("%s", password_Validate_err),
		}
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	err = controller.AuthenticationService.Register(createUserRequest)
	if err != nil {
		fmt.Println("in controller auth..the account already exist..")
		resp := response.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "NOT OK",
			Message: "Account Already exist",
		}
		ctx.JSON(http.StatusOK, resp)
		return

	}
	resp := response.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Succesfully Created user",
	}
	ctx.JSON(http.StatusOK, resp)
}
