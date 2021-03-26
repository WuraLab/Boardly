package middlewares

import (
	"github.com/wuraLab/boardly/src/backend/internal/errors"
	log "github.com/sirupsen/logrus"
	"time"

	"github.com/wuraLab/boardly/src/backend/internal/controllers"
	"github.com/wuraLab/boardly/src/backend/internal/models"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
const (
	identityKey = "ID"
)
func JWTMiddleware(DB *gorm.DB,secret string) *jwt.GinJWTMiddleware {
		authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
			Realm:       "boardly",
			Key:         []byte(secret),
			Timeout:     (time.Minute * 15),
			MaxRefresh:  time.Hour,
			IdentityKey: identityKey,
			PayloadFunc: func(data interface{}) jwt.MapClaims {
				if v, ok := data.(*models.User); ok {
					return jwt.MapClaims{
						identityKey: v.ID,
					}
				}
				return jwt.MapClaims{}
			},
			IdentityHandler: func(c *gin.Context) interface{} {
				claims := jwt.ExtractClaims(c)
				user := models.User{}
				log.Print(claims)
				user.ID = uint(claims[identityKey].(float64))
				return &user
			},
			Authenticator: func(c *gin.Context) (interface{}, error) {
				userController := controllers.User{
					DB: DB,
				}
				user, err := userController.Login(c); 
				if err != nil {
					log.Error(err)
					return nil, jwt.ErrFailedAuthentication
				}

				return user, nil
			},
			Authorizator: func(data interface{}, c *gin.Context) bool {
				if u, ok := data.(*models.User); ok && u.ID != 0 {
					return true
				}
				return false
			},
			Unauthorized: func(c *gin.Context, code int, message string) {
				log.Error(errors.NewBadRequest(message))
				c.JSON(code, gin.H{
					"code":    code,
					"message": message,
				})
			},
			// TokenLookup is a string in the form of "<source>:<name>" that is used
			// to extract token from the request.
			// Optional. Default value "header:Authorization".
			// Possible values:
			// - "header:<name>"
			// - "query:<name>"
			// - "cookie:<name>"
			// - "param:<name>"
			TokenLookup: "header: Authorization, query: token, cookie: jwt",
			// TokenLookup: "query:token",
			// TokenLookup: "cookie:token",

			// TokenHeadName is a string in the header. Default value is "Bearer"
			TokenHeadName: "Bearer",

			// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
			TimeFunc: time.Now,
		})

		if err != nil {
			log.Fatal("JWT Error:" + err.Error())
		}

		// When you use jwt.New(), the function is already automatically called for checking,
		// which means you don't need to call it again.
		errInit := authMiddleware.MiddlewareInit()

		if errInit != nil {
			log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
		}
		return authMiddleware
}