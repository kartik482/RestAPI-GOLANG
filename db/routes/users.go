package routes

import (
	"fmt"
	"net/http"

	"example.com/api/models"
	"example.com/api/utils"
	"github.com/gin-gonic/gin"
)

func SignUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "could not parse request"})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not save user"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
func Login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "could not parse request"})
		return
	}
	if err = user.ValidateCredentials(); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": token})
}
