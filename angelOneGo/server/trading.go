package server

import (
	"fmt"
	"net/http"

	SmartApi "github.com/angel-one/smartapigo"
	"github.com/gin-gonic/gin"
	"github.com/vishal2911/algoTrading/angelOneGo/model"
	"github.com/vishal2911/algoTrading/angelOneGo/util"
)

func (server *Server) GetHoldings(c *gin.Context) (*SmartApi.Holdings, error) {

	//validation is to be done here
	//DB call
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetUsers, "reading all Holdings data", nil)
	holdings, err := server.ABTradingClient.ApiClient.GetHoldings()
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.GetUsers,
			"error while reading holdings data from angelbroking", err)
		return &holdings, fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetUsers,
		"returning all holdings data to api and setting response", holdings)
	c.JSON(http.StatusOK, holdings)
	return &holdings, nil

}
