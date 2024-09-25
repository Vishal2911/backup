package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vishal2911/algoTrading/angelOneGo/api"
	"github.com/vishal2911/algoTrading/angelOneGo/server"
)

// @title User API
// @version 1.0
// @description API for managing users
// @host localhost:8000
// @BasePath /
// @schemes http https
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-Token
func main() {

	api := api.APIRoutes{}
	router := gin.Default()
	server := server.Server{}
	//Starting app
	api.StartApp(router, server)
	// Start the server
	router.Run(":8080")

}
