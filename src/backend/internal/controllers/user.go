package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/wuraLab/boardly/src/backend/internal/models"
	"golang.org/x/crypto/bcrypt"
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
	hashPassword, err := HashPassword("hello")
	user.Password = hashPassword
	rows, err := user.Create(ctrl.DB)

	if err != nil && rows > 0 {
		log.Error(err)
		c.JSON(http.StatusConflict, gin.H{"message": "User already exists"})

	} else if err == nil && rows == 1 {
		c.JSON(http.StatusOK, gin.H{"message": "Successfully registered"})
	} else {
		log.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

}

//Register ...
func (ctrl *User) Login(c *gin.Context) {
	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Error(err)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Invalid form"})
		return
	}
	//check if user exists
	storedCreds := &models.User{Email: user.Email}
	if !storedCreds.Exists(ctrl.DB) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid username or password"})
		return
	}
	if !CheckPasswordHash(user.Password, storedCreds.Password) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid name or password"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully Logged In"})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
