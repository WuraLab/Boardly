package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/wuraLab/boardly/src/backend/internal/commons"
	"github.com/wuraLab/boardly/src/backend/internal/errors"
	"github.com/wuraLab/boardly/src/backend/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//User  Controller
type User struct {
	DB *gorm.DB
}

//Register ...
func (ctrl *User) Register(c *gin.Context) {
	user := models.User{}

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Error(err)
	}
	//check if there is a binding error or empty firstname or lastname
	if err != nil || len(user.FirstName) == 0 || len(user.LastName) == 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "All fields are required"})
		return
	}
	hashPassword, err := HashPassword(user.Password)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
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

//Login ...
func (ctrl *User) Login(c *gin.Context) (*models.User, error) {
	var errMsg string
	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Error(err)
		// c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "All fields are required"})
		return nil, err
	}
	//check if user exists
	storedCreds := &models.User{Email: user.Email}
	if !storedCreds.Exists(ctrl.DB) {
		errMsg = "Invalid username or password"
		// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": errMsg})
		return nil, errors.NewBadRequest(errMsg)
	}
	if !CheckPasswordHash(user.Password, storedCreds.Password) {
		errMsg = "Invalid username or password"
		// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": errMsg})
		return nil, errors.NewBadRequest(errMsg)
	}
	// c.JSON(http.StatusOK, gin.H{"message": "Successfully Logged In"})
	return storedCreds, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// ForgotPassword
func (ctrl *User) ForgotPassword(c *gin.Context) {
	var errMsg string
	user := models.ForgotPassword{}

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Error(err)
		errMsg = "Email is required"
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": errMsg})
		return
	}

	// check if user exists
	storedCreds := &models.User{Email: user.Email}
	if !storedCreds.Exists(ctrl.DB) {
		errMsg = "User does not exist"
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": errMsg})
		return
	}

	// Generate the token that will be used to reset the password
	// resetToken, _ := services.GenerateNonAuthToken(storedCreds.Email)

	// The link to password the password reset
	link := "http://localhost:5000/api/v1/password-reset?reset_token="
	// Define the body of the email
	body := "Here is your reset <a href='" + link + "'>link</a>"
	html := "<strong>" + body + "</strong>"

	// Initialize email sendout
	email := commons.SendMail("Reset Password", storedCreds.Email, html)

	// If email was sent, return 200 status code
	if email == true {
		c.AbortWithStatusJSON(200, gin.H{"messsage": "Check mail"})
		return
		// Return 500 status when something wrong happened
	} else {
		c.AbortWithStatusJSON(500, gin.H{"message": "An issue occurred while sending reset instructions email"})
		return
	}
}
