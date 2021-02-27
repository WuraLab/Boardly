package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/wuraLab/boardly/src/backend/controller"
)

func LoginRoute(router *gin.Engine) {
	router.POST("/login", controller.Login)
}
