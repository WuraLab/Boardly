package routes

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/wuraLab/boardly/src/backend/internal/config"
	"github.com/wuraLab/boardly/src/backend/internal/controllers"
	"github.com/wuraLab/boardly/src/backend/internal/middlewares"
	"gorm.io/gorm"
)

// SetupRouter setup routing here
func SetupRouter(DB *gorm.DB, config *config.Config) *gin.Engine {
	//Start the default gin server
	r := gin.New()

	// Middlewares
	authMiddleware := middlewares.JWTMiddleware(DB, config.Server.JWT, true, true)
	r.Use(middlewares.ErrorHandler)
	r.Use(middlewares.CORSMiddleware())

	//authorization
	e, _ := casbin.NewEnforcer(config.Casbin.Model, config.Casbin.Policy)

	api := r.Group("/api/v1")
	{
		userController := controllers.User{
			DB: DB,
		}
		api.POST("/user/register", userController.RegisterEmployee)
		api.POST("/user/register-admin", userController.RegisterAdmin)
		api.POST("/user/login", authMiddleware.LoginHandler)
	}
	auth := r.Group("/api/v1/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	auth.Use(middlewares.NewAuthorizer(e))
	{
		auth.POST("/refresh-token", authMiddleware.RefreshHandler)
		api.POST("/logout", authMiddleware.LogoutHandler)
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
