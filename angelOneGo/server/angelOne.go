package server

import (
	"fmt"
	"time"

	SmartApi "github.com/angel-one/smartapigo"
	"github.com/pquerna/otp/totp"
	"github.com/vishal2911/algoTrading/angelOneGo/model"
)

// // model.Credentials struct to store client credentials
// type model.Credentials struct {
// 	ClientCode    string
// 	Password      string
// 	TradingKey    string
// 	HistoryAPIKey string
// 	TOTPSecret    string
// }

// // HistoryDataClient struct for handling historical data operations
// type HistoryDataClient struct {
// 	apiClient *SmartApi.Client
// 	creds     model.Credentials
// }

// //TradingClient struct for handling trading operations
// type TradingClient struct {
// 	apiClient *SmartApi.Client
// 	session   *SmartApi.UserSession
// 	creds     model.Credentials
// }

// NewHistoryDataClient initializes a new history data client
func NewHistoryDataClient(creds model.Credentials) *model.HistoryDataClient {
	apiClient := SmartApi.New(creds.ClientCode, creds.Password, creds.HistoryAPIKey)
	return &model.HistoryDataClient{
		ApiClient: apiClient,
		Creds:     creds,
	}
}

// NewTradingClient initializes a new trading client
func NewTradingClient(creds model.Credentials) *model.TradingClient {
	apiClient := SmartApi.New(creds.ClientCode, creds.Password, creds.TradingKey)
	return &model.TradingClient{
		ApiClient: apiClient,
		Creds:     creds,
	}
}

// GenerateTOTP generates a TOTP code
func (server *Server) GenerateTOTP() (string, error) {
	return totp.GenerateCode(server.ABTradingClient.Creds.TOTPSecret, time.Now())
}

// Login logs in the user and generates a session
func (server *Server) Login() error {
	totpCode, err := server.GenerateTOTP()
	if err != nil {
		return fmt.Errorf("error generating TOTP code: %v", err)
	}

	session, err := server.ABTradingClient.ApiClient.GenerateSession(totpCode)
	if err != nil {
		return fmt.Errorf("error generating session: %v", err)
	}
	server.ABTradingClient.Session = &session
	return nil
}

// RenewAccessToken renews the user tokens using the refresh token
func (server *Server) RenewAccessToken() error {
	tokens, err := server.ABTradingClient.ApiClient.RenewAccessToken(server.ABTradingClient.Session.RefreshToken)
	if err != nil {
		return fmt.Errorf("error renewing access token: %v", err)
	}
	server.ABTradingClient.Session.UserSessionTokens = tokens
	return nil
}

// GetUserProfile retrieves the user's profile
func (server *Server) GetUserProfile() (*SmartApi.UserProfile, error) {
	profile, err := server.ABTradingClient.ApiClient.GetUserProfile()
	if err != nil {
		return nil, fmt.Errorf("error getting user profile: %v", err)
	}
	return &profile, nil
}

// PlaceOrder places a new order
func (server *Server) PlaceOrder(params SmartApi.OrderParams) (*SmartApi.OrderResponse, error) {
	orderResponse, err := server.ABTradingClient.ApiClient.PlaceOrder(params)
	if err != nil {
		return nil, fmt.Errorf("error placing order: %v", err)
	}
	return &orderResponse, nil
}

// // GetHistoricalData retrieves historical data
// func (c *HistoryDataClient) GetHistoricalData(symbol string, from time.Time, to time.Time) ([]SmartApi.Candle, error) {
//     data, err := server.ABTradingClient.ApiClient.GetHistoricalData(symbol, from, to)
//     if err != nil {
//         return nil, fmt.Errorf("error getting historical data: %v", err)
//     }
//     return data, nil
// }

// func (server *Server) testing() {

// 	// historyDataClient := NewHistoryDataClient(creds)
// 	// tradingClient := NewTradingClient(creds)

// 	// Use historyDataClient to retrieve historical data
// 	// historicalData, err := historyDataClient.GetHistoricalData("SBIN-EQ", time.Now().AddDate(0, 0, -7), time.Now())
// 	// if err != nil {
// 	//     fmt.Println("Error getting historical data:", err)
// 	// } else {
// 	//     fmt.Println("Historical Data:", historicalData)
// 	// }

// 	// // Use tradingClient to perform trading operations
// 	// if err := tradingClient.Login(); err != nil {
// 	// 	fmt.Println("Error logging in:", err)
// 	// 	return
// 	// }

// 	// if err := tradingClient.RenewAccessToken(); err != nil {
// 	// 	fmt.Println("Error renewing access token:", err)
// 	// 	return
// 	// }

// 	// profile, err := tradingClient.GetUserProfile()
// 	// if err != nil {
// 	// 	fmt.Println("Error getting user profile:", err)
// 	// 	return
// 	// }

// 	// fmt.Println("User Profile :- ", profile)
// 	// fmt.Println("User Session Tokens :- ", tradingClient.session.UserSessionTokens)

// 	// orderParams := SmartApi.OrderParams{
// 	// 	Variety:         "NORMAL",
// 	// 	TradingSymbol:   "SBIN-EQ",
// 	// 	SymbolToken:     "3045",
// 	// 	TransactionType: "BUY",
// 	// 	Exchange:        "NSE",
// 	// 	OrderType:       "LIMIT",
// 	// 	ProductType:     "INTRADAY",
// 	// 	Duration:        "DAY",
// 	// 	Price:           "19500",
// 	// 	SquareOff:       "0",
// 	// 	StopLoss:        "0",
// 	// 	Quantity:        "1",
// 	// }

// 	// order, err := tradingClient.PlaceOrder(orderParams)
// 	// if err != nil {
// 	// 	fmt.Println("Error placing order:", err)
// 	// 	return
// 	// }

// 	// fmt.Println("Placed Order ID and Script :- ", order)

// 	holdings, err := server.ABTradingClient.ApiClient.GetHoldings()
// 	if err != nil {
// 		fmt.Printf("err : %v", err)
// 	} else {
// 		fmt.Println()
// 		for _, holds := range holdings {
// 			fmt.Printf("%v\n", holds)
// 		}

// 	}

// }
