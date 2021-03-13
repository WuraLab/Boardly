package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"

	"net/http/httptest"
	"testing"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wuraLab/boardly/src/backend/internal/controllers"
	"github.com/wuraLab/boardly/src/backend/internal/middlewares"
	"github.com/wuraLab/boardly/src/backend/internal/models"
	"gorm.io/gorm"
	"github.com/stretchr/testify/assert"
)

var DB *gorm.DB

// SetupRouter setup routing here
func SetupRouter(DB *gorm.DB) *gin.Engine {
	//Start the default gin server
	r := gin.New()

	// Middlewares
	r.Use(middlewares.ErrorHandler)
	r.Use(middlewares.CORSMiddleware())

	api := r.Group("/api/v1")
	{
		userController := controllers.User{
			Base: controllers.Base{
				DB: DB,
			},
		}

		api.POST("/user/register", userController.Register)

		api.POST("/user/login", userController.Login)

	}
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Welcome to Boardly",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{})
	})

	return r
}

func main() {
	r := SetupRouter(DB)
	r.Run()
}

/**
* TestRegister
* Test user registration
*
* Must return response code 200
 */
 func TestRegister(t *testing.T) {
	testRouter := SetupRouter(DB)

	var registerForm models.User

	registerForm.FirstName = "testing"
	registerForm.LastName = "tester"
	registerForm.Email = "tester@test.com"
	registerForm.Password = "testPassword123$"

	data, _ := json.Marshal(registerForm)

	req, err := http.NewRequest("POST", "/api/v1/user/register", bytes.NewBufferString(string(data)))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()

	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, http.StatusOK)
}