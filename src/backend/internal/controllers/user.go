package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/wuraLab/boardly/src/backend/internal/models"
)

//User  Controller
type User struct {
	Base
}

var userModel = models.User{}

//getUserID ...
func getUserID(c *gin.Context) (userID int64) {
	//MustGet returns the value for the given key if it exists, otherwise it panics.
	return c.MustGet("userID").(int64)
}

//Register ...
func (ctrl *User) Register(c *gin.Context) {
	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Error(err)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form"})
		return
	}

	rows, err := user.Create(ctrl.DB)

	if err != nil && rows > 0 {
		c.JSON(http.StatusConflict, gin.H{"message": "User already exists", "err": err.Error()})

	} else if err == nil && rows == 1 {
		c.JSON(http.StatusOK, gin.H{"message": "Successfully registered"})
	} else {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
	}

}
