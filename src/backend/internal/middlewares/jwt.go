package middlewares

import (
	"github.com/wuraLab/boardly/src/backend/internal/errors"
	log "github.com/sirupsen/logrus"
	"time"
	"net/http"
	"github.com/wuraLab/boardly/src/backend/internal/controllers"
	"github.com/wuraLab/boardly/src/backend/internal/models"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
const (
	identityKey = "ID"
)

func JWTMiddleware(DB *gorm.DB,secret string, secureCookie bool, httpOnly bool) *jwt.GinJWTMiddleware {
		authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
			Realm:       "boardly",
			Key:         []byte(secret),
			SigningAlgorithm: "HS256",
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
					"status":    code,
					"message": message,
				})
			},
			LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
				c.JSON(http.StatusOK, gin.H{
					"status":   http.StatusOK,
					"message": "Logged in successfully"
					"token":  token,
					"expire": expire.Format(time.RFC3339),
				})
			},
			SendCookie: true,
			SecureCookie: secureCookie,
			CookieHTTPOnly: httpOnly,
			TokenLookup: "header: Authorization, query: token, cookie: jwt",
			TokenHeadName: "Bearer",
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