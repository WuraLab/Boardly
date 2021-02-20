package controllers

import (
	"database/sql"
	"fmt"
	"github.com/go-postgres-jwt-react-starter/server/config"
	"net/http"
	"time"

	"github.com/go-postgres-jwt-react-starter/server/db"
	"github.com/go-postgres-jwt-react-starter/server/errors"
	"github.com/go-postgres-jwt-react-starter/server/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("secret")

//Claims jwt claims struct
type Claims struct {
	db.User
	jwt.StandardClaims
}

// Pong tests that api is working
func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
}


// Login controller
func Login(c *gin.Context) {
	var user db.Login
	c.Bind(&user)

	row := db.DB.QueryRow(db.LoginQuery, user.Email)

	var id int
	var name, email, password, createdAt, updatedAt string

	err := row.Scan(&id, &name, &password, &email, &createdAt, &updatedAt)

	if err == sql.ErrNoRows {
		fmt.Println(sql.ErrNoRows, "err")
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "msg": "incorrect credentials"})
		return
	}

	match := db.CheckPasswordHash(user.Password, password)
	if !match {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "msg": "incorrect credentials"})
		return
	}

	//expiration time of the token ->30 mins
	expirationTime := time.Now().Add(30 * time.Minute)

	// Create the JWT claims, which includes the User struct and expiry time
	claims := &Claims{

		User: db.User{
			Name: name, Email: email, CreatedAt: createdAt, UpdatedAt: updatedAt,
		},
		StandardClaims: jwt.StandardClaims{
			//expiry time, expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT token string
	tokenString, err := token.SignedString(jwtKey)
	errors.HandleErr(c, err)
	// c.SetCookie("token", tokenString, expirationTime, "", "*", true, false)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	fmt.Println(tokenString)
	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "logged in succesfully", "user": claims.User, "token": tokenString})
}