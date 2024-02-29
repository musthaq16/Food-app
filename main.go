package main

import (
	"fmt"
	"food-app/config"
	"food-app/controller"
	"food-app/helper"
	"food-app/model"
	"food-app/repository"
	"food-app/router"
	"food-app/service"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func main() {

	loadConfig, err := config.LoadConfig(".")

	if err != nil {
		fmt.Println("there is error in loading configfile..", err)
		return
	}

	db := config.ConnectDB(&loadConfig)
	validate := validator.New()

	db.AutoMigrate(&model.User{}, &model.OneTimePassword{})

	//initialize Repository
	userRepository := repository.NewUserRepositoryImpl(db)

	//initialize Serice
	authenticationService := service.NewAuthenticationServiceImpl(userRepository, validate)

	//initialize Controller
	authenticationController := controller.NewAuthenticationController(authenticationService)
	usersController := controller.NewUserController(userRepository)

	//Routes
	routes := router.NewRouter(userRepository, authenticationController, usersController)

	server := &http.Server{
		Addr:    loadConfig.ServerPort,
		Handler: routes,
	}

	server_err := server.ListenAndServe()

	helper.ErrorPanic(server_err)

}
