package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/vishal2911/algoTrading/angelOneGo/model"
	"github.com/vishal2911/algoTrading/angelOneGo/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Middleware function for token authentication
func (api APIRoutes) AuthMiddlewareComplete() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader(model.Authorization)
		if tokenString == "" {
			util.Log(model.LogLevelInfo, model.ApiPackageLevel,
				model.AuthMiddlewareComplete, "token string empty", nil)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token"})
			c.Abort()
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return model.SecretKey, nil
		})
		if err != nil || !token.Valid {
			util.Log(model.LogLevelError, model.ApiPackageLevel,
				model.AuthMiddlewareComplete, "token value is not valid", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			c.Abort()
			return
		}
		util.Log(model.LogLevelInfo, model.ApiPackageLevel,
			model.AuthMiddlewareComplete, "token value parsed", token)

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok {
			email := claims["email"].(string)
			password := claims["password"].(string)
			util.Log(model.LogLevelInfo, model.ApiPackageLevel,
				model.AuthMiddlewareComplete, "token value for email and password", email+" password = "+password)
			db, err := gorm.Open(postgres.Open(model.DSN), &gorm.Config{})
			defer func() {
				sqldb, err := db.DB()
				if err != nil {
					util.Log(model.LogLevelError, model.ApiPackageLevel,
						model.AuthMiddlewareComplete, "error in geting sql object for middleware", nil)
				}
				sqldb.Close()
				util.Log(model.LogLevelInfo, model.ApiPackageLevel,
					model.AuthMiddlewareComplete, "middleware db connection closed", nil)
			}()
			if err != nil {
				util.Log(model.LogLevelError, model.ApiPackageLevel, model.NewStore,
					"error while connecting database", err)
				panic(err)
			}
			var user model.User
			util.Log(model.LogLevelInfo, model.ApiPackageLevel, model.AuthMiddlewareComplete,
				"reading user data from db based on email", email+" pass = "+password)
			resp := db.Where("email = ? AND password = ?", email, password).First(&user)
			if resp.Error != nil {
				util.Log(model.LogLevelError, model.ApiPackageLevel, model.AuthMiddlewareComplete,
					"error while reading user data", resp.Error)
				c.JSON(http.StatusUnauthorized, gin.H{"error": "error in geting user data from Database"})
				c.Abort()
				return
			}
			if !validateUserCredability(user, c.Request.URL.Path, c.Request.Method) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "user is not valid for request"})
				c.Abort()
				return
			}
			util.Log(model.LogLevelInfo, model.ApiPackageLevel, model.AuthMiddlewareComplete,
				"returning user data", user)
		}
		c.Next()
	}
}

func validateUserCredability(user model.User, url, protocal string) bool {
	util.Log(model.LogLevelInfo, model.ApiPackageLevel, model.AuthMiddlewareComplete,
		"validating user credability", user.Email+" Password="+user.Password+" url="+url+" protocal="+protocal)
	if user.Type == model.SuperAdminUser {
		return true
	} else if user.Type == model.AdminUser {

	} else if user.Type == model.NormalUser && protocal == "get" {
		if strings.Contains(url, model.ActuatorsURL) {

		}

	} else {

	}
	return true
}
