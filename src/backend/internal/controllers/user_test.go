package controllers_test

import (
	"bytes"
	"encoding/json"
	log "github.com/sirupsen/logrus"
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
			DB: DB,
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

/**
* TestRegister
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
			log.Fatalln(err)
		}

		resp := httptest.NewRecorder()

		testRouter.ServeHTTP(resp, req)
		assert.Equal(t, testCase.expected, resp.Code)
   }
}


func TestLogin(t *testing.T) {

	testRouter := SetupRouter(DB)

	//create user first
	idealCase := models.User{
		FirstName: "loginfirstname",
		LastName:  "loginlastname",
		Email:     "login@test.com",
		Password:  "loginPassword123$",		
	}
	data, _ := json.Marshal(idealCase)

	req, err := http.NewRequest("POST", "/api/v1/user/register", bytes.NewBufferString(string(data)))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		log.Fatalln(err)
	}
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	testCases := []struct{
		input          models.User
		expected int
	  }{
		//missing email
		{
		  input: models.User{
						Email:     "",
						Password:  idealCase.Password,
		  },
		  expected: http.StatusUnprocessableEntity,
		},
		//missing password
		{
			input: models.User{
						  Email:     idealCase.Email,
						  Password:  "",
			},
			expected: http.StatusUnprocessableEntity,
		},
		//non existing email
		{
			input: models.User{
						  Email:     "wrong@test.com",
						  Password:  idealCase.Password,
			},
			expected: http.StatusUnauthorized,
		},
		//wrong password
		{
			input: models.User{
						  Email:     idealCase.Email,
						  Password:  "wrongpassword",
			},
			expected: http.StatusUnauthorized,
		},
		//completely filled out
		{
			input: models.User{
						  Email:     idealCase.Email,
						  Password:  idealCase.Password,
			},
			expected: http.StatusOK,
		},
	}
   for _, testCase := range testCases {

		data, _ := json.Marshal(testCase.input)

		req, err := http.NewRequest("POST", "/api/v1/user/login", bytes.NewBufferString(string(data)))
		req.Header.Set("Content-Type", "application/json")

		if err != nil {
			log.Fatalln(err)
		}

		resp := httptest.NewRecorder()

		testRouter.ServeHTTP(resp, req)
		assert.Equal(t, testCase.expected, resp.Code)
   }
}
