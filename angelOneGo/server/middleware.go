package server

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang-jwt/jwt"
// 	"github.com/vishal2911/algoTrading/angelOneGo/model"
// 	"github.com/vishal2911/algoTrading/angelOneGo/util"
// )

// // Middleware function for token authentication
// func (server Server) AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		tokenString := c.GetHeader(model.Authorization)
// 		if tokenString == "" {
// 			util.Log(model.LogLevelInfo, model.ServerPackageLavel,
// 				model.AuthMiddleware, "token string empty", nil)
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token"})
// 			c.Abort()
// 			return
// 		}
// 		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 			return model.SecretKey, nil
// 		})
// 		if err != nil || !token.Valid {
// 			util.Log(model.LogLevelError, model.ServerPackageLavel,
// 				model.AuthMiddleware, "token value is not valid", err)
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
// 			c.Abort()
// 			return
// 		}
// 		util.Log(model.LogLevelInfo, model.ServerPackageLavel,
// 			model.AuthMiddleware, "token value parsed", token)
// 		// Proceed to the next handler
// 		claims, ok := token.Claims.(jwt.MapClaims)
// 		if ok {

// 			email := claims["email"].(string)
// 			password := claims["password"].(string)
// 			usersignin := model.UserSignIn{
// 				Email:    email,
// 				Password: password,
// 			}
// 			util.Log(model.LogLevelInfo, model.ServerPackageLavel,
// 				model.AuthMiddleware, "token value for email and password", "email = "+email+" password = "+password)
// 			user, err := server.Pgress.SingIn(usersignin)
// 			if err != nil {
// 				util.Log(model.LogLevelError, model.ServerPackageLavel,
// 					model.AuthMiddleware, "error in reading userdata based on email", err)
// 				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
// 				c.Abort()
// 				return
// 			} else {
// 				if user.Email != email || user.Password != password {
// 					util.Log(model.LogLevelInfo, model.ServerPackageLavel,
// 						model.AuthMiddleware, "userEmail and userPassword is not matched", user)
// 					c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
// 					c.Abort()

// 				}
// 			}
// 		}
// 		c.Next()
// 	}
// }
