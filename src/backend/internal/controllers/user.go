package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/wuraLab/boardly/src/backend/internal/errors"
	"github.com/wuraLab/boardly/src/backend/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//User  Controller
type User struct {
	DB   *gorm.DB
	Role string
}

//Register ...
func (ctrl *User) Registration(c *gin.Context) error {
	user := models.User{}

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Error(err)
	}
	//check if there is a binding error or empty firstname or lastname
	if err != nil || len(user.FirstName) == 0 || len(user.LastName) == 0 {
		errMsg := "All fields are required"
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": errMsg})
		return errors.NewBadRequest(errMsg)
	}
	hashPassword, err := HashPassword(user.Password)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return err
	}
	user.Password = hashPassword
	//assign default role
	user.Role = ctrl.Role
	rows, err := user.Create(ctrl.DB)

	if err != nil && rows > 0 {
		log.Error(err)
		c.JSON(http.StatusConflict, gin.H{"message": "User already exists"})
	} else if err == nil && rows == 1 {
		c.JSON(http.StatusOK, gin.H{"message": "Successfully registered"})
		return nil
	} else {
		log.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	return err
}

//Register ...
func (ctrl *User) RegisterEmployee(c *gin.Context) {
	ctrl.Role = models.DEFAULT_ROLE
	ctrl.Registration(c)
}

func (ctrl *User) RegisterAdmin(c *gin.Context) {
	//check if a User is empty
	err := ctrl.DB.First(&models.User{}, 1).Error
	if errors.Is(err, gorm.ErrRecordNotFound) == false {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "Admin has already been created"})
		return
	}

	ctrl.Role = models.ADMIN_ROLE
	ctrl.Registration(c)
}

// //Register admin user...
// func (ctrl *User) RegisterAdmin(c *gin.Context) {
// 	user := models.User{}

// 	err := c.ShouldBindJSON(&user)

// 	if err != nil {
// 		log.Error(err)
// 	}

// 	// user.UserRole points to user.Role.UserRole
// 	// this is known as a field promotion: where all the nested anonymous
// 	// struct fields are automatically available on the parent struct. cool:)
// 	if len(user.UserRole) == 0 {
// 		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "'Role' is required"})
// 		return
// 	}

// 	//check if there is a binding error or empty firstname or lastname
// 	if err != nil || len(user.FirstName) == 0 || len(user.LastName) == 0 {
// 		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": "All fields are required"})
// 		return
// 	}
// 	hashPassword, err := HashPassword(user.Password)
// 	if err != nil {
// 		log.Error(err)
// 		c.JSON(http.StatusInternalServerError, gin.H{})
// 		return
// 	}
// 	user.Password = hashPassword
// 	rows, err := user.Create(ctrl.DB)

// 	if err != nil && rows > 0 {
// 		log.Error(err)
// 		c.JSON(http.StatusConflict, gin.H{"message": "User already exists"})

// 	} else if err == nil && rows == 1 {
// 		c.JSON(http.StatusOK, gin.H{"message": "Successfully registered"})
// 	} else {
// 		log.Error(err)
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
// 	}
// }

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
