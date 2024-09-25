package api

import (
	"github.com/gin-gonic/gin"
)

func (api APIRoutes) TradingRouts(router *gin.Engine) {
	// Define routes
	userapi := router.Group("/trading")
	{
		userapi.GET("/holings", api.GetHoldings)
	}

}

// Handler to get all holings
// @router /trading/holings [get]
// @summary Get all holings
// @tags holings
// @produce json
// @success 200 {array} smartapigo.Holdings
// @Security ApiKeyAuth
func (api APIRoutes) GetHoldings(c *gin.Context) {
	api.Server.GetHoldings(c)
}
