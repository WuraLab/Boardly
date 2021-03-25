package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wuraLab/boardly/src/backend/internal/controllers"
	"github.com/wuraLab/boardly/src/backend/internal/middlewares"
	"github.com/wuraLab/boardly/src/backend/internal/config"
	"gorm.io/gorm"
)

// SetupRouter setup routing here
func SetupRouter(DB *gorm.DB, config *config.Config) *gin.Engine {
	//Start the default gin server
	r := gin.New()

	// Middlewares
	authMiddleware := middlewares.JWTMiddleware(DB,config.Server.JWT)
	r.Use(middlewares.ErrorHandler)
	r.Use(middlewares.CORSMiddleware())

	api := r.Group("/api/v1")
	{
		userController := controllers.User{
			DB: DB,
		}
		api.POST("/user/register", userController.Register)

		api.POST("/user/login", authMiddleware.LoginHandler)

	}
	auth := r.Group("/api/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.POST("/refresh_token", authMiddleware.RefreshHandler)
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
