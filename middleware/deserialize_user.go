package middleware

import (
	"fmt"
	"food-app/config"
	"food-app/helper"
	"food-app/repository"
	"food-app/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func DeserializeUser(userRepository repository.UserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token string
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		}

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "Fail", "message": "Not logged in."})
			return
		}

		config, err := config.LoadConfig(".")
		if err != nil {
			fmt.Println("Error in loading config", err)
			return
		}
		sub, err := utils.ValidateToken(token, config.SecretKey)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "Fail", "message": err.Error()})
			return
		}

		id, err := strconv.Atoi(fmt.Sprint(sub))
		helper.ErrorPanic(err)

		result, err := userRepository.FindById(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "Fail", "message": "the user belong to this token no longer exist.."})
			return
		}
		ctx.Set("currentUser", result.UserName)
		ctx.Set("currentId", result.Id)
		ctx.Next()

	}
}
