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
	authenticationRouter.POST("/forgetPassword", authenticationController.ForgetPassword)
	authenticationRouter.POST("/resetPassword/:otpToken", authenticationController.ResetPassword)

	userRouter := router.Group("/user")
	userRouter.GET("/Getallusers", middleware.DeserializeUser(userRepository), userController.GetAllUsers)
	userRouter.PATCH("/UpdateMe", middleware.DeserializeUser(userRepository), userController.UpdateMe)
	userRouter.DELETE("/DeleteMe", middleware.DeserializeUser(userRepository), userController.DeleteUser)

	return service
}
