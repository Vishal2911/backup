package server

import (
	"fmt"

	SmartApi "github.com/angel-one/smartapigo"
	"github.com/gin-gonic/gin"
	"github.com/vishal2911/algoTrading/angelOneGo/model"
	"github.com/vishal2911/algoTrading/angelOneGo/store/pgress"
	"github.com/vishal2911/algoTrading/angelOneGo/util"
	// "github.com/sirupsen/logrus"
)

type Server struct {
	Pgress              pgress.StoreOperation
	ABTradingClient     *model.TradingClient
	ABHistoryDataClient *model.HistoryDataClient
	AngeloneClinet      *SmartApi.Client
}

func (s *Server) NewServer(store pgress.PgressStore) {
	creds := model.Credentials{
		ClientCode:    "V58985289",
		Password:      "1250",
		HistoryAPIKey: "NlvfasR9",
		TradingKey:    "2t2TvbKy",
		TOTPSecret:    "B3EHYOJOOCFEUVWK774H5GLC4I",
	}
	s.Pgress = &store
	util.SetLoger()
	util.Logger.Info("Logger Setup Done at server level")
	fmt.Println("Creating new Store .....")

	s.ABTradingClient = NewTradingClient(creds)
	s.Login()
	// createSession()
	// s.Pgress.NewStore()

	// s.testing()
}

type ServerOperation interface {
	//Postgress config methods
	NewServer(store pgress.PgressStore)

	GetHoldings(c *gin.Context) (*SmartApi.Holdings, error)

	// //middleware
	// AuthMiddleware() gin.HandlerFunc
}
