package router

import (
	"food-app/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(authenticationController *controller.AuthenticationController) *gin.Engine {
	service := gin.Default()

	service.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello just from router testing...!!")
	})

	router := service.Group("/api")

	authenticationRouter := router.Group("/auth")
	authenticationRouter.POST("/register", authenticationController.Register)
	authenticationRouter.POST("/login", authenticationController.LogIn)

	return service
}
