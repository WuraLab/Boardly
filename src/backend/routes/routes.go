package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/robogeek95/ginCrud/controllers"
)

// SetupRouter ... setup the router
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// log := logger.configLogger()

	// Fixes CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		// log.InfoLogger.Println("handled cors")

		c.Next()
	})

	group := r.Group("/api/v1/")
	{
		group.GET("/", controllers.GetWelcome)
		// All auth endpoints
		users := group.Group("/auth")
		{
			users.POST("/register", controllers.RegisterAdmin)
		}
	}
	return r
}
