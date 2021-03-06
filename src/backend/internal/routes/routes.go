package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wuraLab/boardly/src/backend/internal/controllers"
	"github.com/wuraLab/boardly/src/backend/internal/middlewares"
	"gorm.io/gorm"
)

// SetupRouter setup routing here
func SetupRouter() *gin.Engine {
	var DB *gorm.DB
	//Start the default gin server
	r := gin.Default()

	// Middlewares
	r.Use(middlewares.ErrorHandler)
	r.Use(middlewares.CORSMiddleware())

	api := r.Group("/api/v1")

	// routes
	{
		/*** START USER ***/
		user := controllers.User{
			Base: controllers.Base{
				DB: DB,
			},
		}

		// v1.POST("/user/login", user.Login)
		api.POST("/user/register", user.Register)
		// v1.GET("/user/logout", user.Logout)
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
