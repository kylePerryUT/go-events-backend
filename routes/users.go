package routes

import (
	"net/http"

	"example.com/example/models"
	"github.com/gin-gonic/gin"
)

func createUser(context *gin.Context) {
	var user models.USER
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"There was an error reading the user data": err.Error()})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"There was an error saving the user.": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully!", "user": user})
}
