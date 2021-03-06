package controllers

import (
	"fmt"
	"net/http"

	"github.com/wuraLab/boardly/src/backend/internal/db"
	"github.com/wuraLab/boardly/src/backend/internal/errors"
	"github.com/wuraLab/boardly/src/backend/internal/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("secret")

//Claims jwt claims struct
type Claims struct {
	models.User
	jwt.StandardClaims
}

// Pong tests that api is working
func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
}

// checkUserExists function
func checkUserExists(user models.User) bool {
	rows, err := db.DB.Query(db.CheckUserExists, user.Email)
	if err != nil {
		return false
	}
	if !rows.Next() {
		return false
	}
	return true
}

//Create new user
func Create(c *gin.Context) {
	var user models.User
	c.Bind(&user)
	exists := checkUserExists(user)

	valErr := models.ValidateUser(user, errors.ValidationErrors)
	if exists == true {
		valErr = append(valErr, "email already exists")
	}
	fmt.Println(valErr)
	if len(valErr) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": valErr})
		return
	}
	_, err := db.DB.Query(db.CreateUserQuery, user.FirstName, user.LastName, user.Password, user.Email)
	errors.HandleErr(c, err)
	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "User created succesfully"})
}
