package errors

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// ValidationErrors Users input validation errros
var ValidationErrors = []string{}

//HandleErr //generic error handler, logs error and Os.Exit(1)
func HandleErr(c *gin.Context, err error) error {
	if err != nil {
		log.Error(c.Error(err))
	}
	return err
}
