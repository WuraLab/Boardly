package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/wuraLab/boardly/src/backend/internal/controllers"
	// "github.com/wuraLab/boardly/src/backend/internal/middlewares"
	"github.com/wuraLab/boardly/src/backend/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// SetupRouter setup routing here
func SetupRouter(DB *gorm.DB) *gin.Engine {
	//Start the default gin server
	r := gin.Default()
	gin.SetMode(gin.TestMode)

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
	r.GET("/", controllers.Home)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  404,
			"message": "Route Not Found",
		})
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

	testCases := []struct{
		input          models.User
		expected int
	  }{
		//missing email
		{
		  input: models.User{
						FirstName: "testing",
						LastName:  "tester",
						Email:     "",
						Password:  "testPassword123$",
		  },
		  expected: http.StatusUnprocessableEntity,
		},
		//missing password
		{
			input: models.User{
						  FirstName: "testing",
						  LastName:  "tester",
						  Email:     "tester@test.com",
						  Password:  "",
			},
			expected: http.StatusUnprocessableEntity,
		},
		//
		{
			input: models.User{
						  FirstName: "",
						  LastName:  "tester",
						  Email:     "tester@test.com",
						  Password:  "testPassword123$",
			},
			expected: http.StatusUnprocessableEntity,
		},
		//compltely filled out
		{
			input: models.User{
						  FirstName: "testing",
						  LastName:  "tester",
						  Email:     "tester@test.com",
						  Password:  "testPassword123$",
			},
			expected: http.StatusOK,
		},
	}
   for _, testCase := range testCases {

		data, _ := json.Marshal(testCase.input)

		req, err := http.NewRequest("POST", "/api/v1/user/register", bytes.NewBufferString(string(data)))
		req.Header.Set("Content-Type", "application/json")

		if err != nil {
			fmt.Println(err)
		}

		resp := httptest.NewRecorder()

		testRouter.ServeHTTP(resp, req)
		assert.Equal(t, testCase.expected, resp.Code)
   }
}
