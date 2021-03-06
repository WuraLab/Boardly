package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wuraLab/boardly/src/backend/internal/controllers"
)

// SetupRouter ... setup the router
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// log := logger.configLogger()

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

		auth := group.Group("/auth")
		{
			admin := auth.Group("/admin")
			{
				admin.POST("/register", controllers.RegisterAdmin)
			}
			// 	materials.POST("/", controllers.CreateMaterial)
			// 	materials.GET("/:id", controllers.GetMaterialByID)
			// 	materials.PUT("/:id", controllers.UpdateMaterial)
			// 	materials.DELETE("/:id", controllers.DeleteMaterial)
		}

	}
	return r
}
