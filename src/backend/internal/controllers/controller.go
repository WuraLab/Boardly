package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wuraLab/boardly/src/backend/internal/models"
)

//UserController ...
type UserController struct{}

var userModel = new(models.UserModel)

// GetWelcome shows greeting message
func GetWelcome(c *gin.Context) {
	message := "Hello world :)..."

	c.JSON(http.StatusOK, message)
}

// RegisterAdmin function to create an admin account
func RegisterAdmin(c *gin.Context) {
	var user models.User

	if c.ShouldBindJSON(&user) != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Invalid data format"})
		return
	}

	err := userModel.RegisterAdmin(&user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully registered", "user": user})
}
