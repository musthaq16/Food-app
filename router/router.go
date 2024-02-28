package router

import (
	"food-app/controller"
	"food-app/middleware"
	"food-app/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(userRepository repository.UserRepository, authenticationController *controller.AuthenticationController, userController *controller.UserController) *gin.Engine {
	service := gin.Default()

	service.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello just from router testing...!!")
	})

	router := service.Group("/api")

	authenticationRouter := router.Group("/auth")
	authenticationRouter.POST("/register", authenticationController.Register)
	authenticationRouter.POST("/login", authenticationController.LogIn)

	userRouter := router.Group("/user")
	userRouter.GET("/Getallusers", middleware.DeserializeUser(userRepository), userController.GetAllUsers)
	userRouter.PATCH("/UpdateMe", middleware.DeserializeUser(userRepository), userController.UpdateMe)

	return service
}
