package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/vishal2911/algoTrading/angelOneGo/docs"
	"github.com/vishal2911/algoTrading/angelOneGo/server"
	"github.com/vishal2911/algoTrading/angelOneGo/store/pgress"
)

type APIRoutes struct {
	Server server.ServerOperation
}

func (api APIRoutes) StartApp(router *gin.Engine, server server.Server) {

	//Setup Server
	fmt.Println("Creating new server .....")
	api.Server = &server
	api.Server.NewServer(pgress.PgressStore{})

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// // user routs
	// api.UserRouts(router)

	// trading routs
	api.TradingRouts(router)

}
