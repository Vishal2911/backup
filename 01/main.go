

package main

import (
	"fmt"
	"time"

	SmartApi "github.com/angel-one/smartapigo"
	"github.com/pquerna/otp/totp"
)

// Credentials struct to store client credentials
type Credentials struct {
	ClientCode    string
	Password      string
	TradingKey    string
	HistoryAPIKey string
	TOTPSecret    string
}

// HistoryDataClient struct for handling historical data operations
type HistoryDataClient struct {
	apiClient *SmartApi.Client
	creds     Credentials
}

// TradingClient struct for handling trading operations
type TradingClient struct {
	apiClient *SmartApi.Client
	session   *SmartApi.UserSession
	creds     Credentials
}

// NewHistoryDataClient initializes a new history data client
func NewHistoryDataClient(creds Credentials) *HistoryDataClient {
	apiClient := SmartApi.New(creds.ClientCode, creds.Password, creds.HistoryAPIKey)
	return &HistoryDataClient{
		apiClient: apiClient,
		creds:     creds,
	}
}

// NewTradingClient initializes a new trading client
func NewTradingClient(creds Credentials) *TradingClient {
	apiClient := SmartApi.New(creds.ClientCode, creds.Password, creds.TradingKey)
	return &TradingClient{
		apiClient: apiClient,
		creds:     creds,
	}
}

// GenerateTOTP generates a TOTP code
func (c *TradingClient) GenerateTOTP() (string, error) {
	return totp.GenerateCode(c.creds.TOTPSecret, time.Now())
}

// Login logs in the user and generates a session
func (c *TradingClient) Login() error {
	totpCode, err := c.GenerateTOTP()
	if err != nil {
		return fmt.Errorf("error generating TOTP code: %v", err)
	}

	session, err := c.apiClient.GenerateSession(totpCode)
	if err != nil {
		return fmt.Errorf("error generating session: %v", err)
	}
	c.session = &session
	return nil
}

// RenewAccessToken renews the user tokens using the refresh token
func (c *TradingClient) RenewAccessToken() error {
	tokens, err := c.apiClient.RenewAccessToken(c.session.RefreshToken)
	if err != nil {
		return fmt.Errorf("error renewing access token: %v", err)
	}
	c.session.UserSessionTokens = tokens
	return nil
}

// GetUserProfile retrieves the user's profile
func (c *TradingClient) GetUserProfile() (*SmartApi.UserProfile, error) {
	profile, err := c.apiClient.GetUserProfile()
	if err != nil {
		return nil, fmt.Errorf("error getting user profile: %v", err)
	}
	return &profile, nil
}

// PlaceOrder places a new order
func (c *TradingClient) PlaceOrder(params SmartApi.OrderParams) (*SmartApi.OrderResponse, error) {
	orderResponse, err := c.apiClient.PlaceOrder(params)
	if err != nil {
		return nil, fmt.Errorf("error placing order: %v", err)
	}
	return &orderResponse, nil
}

// // GetHistoricalData retrieves historical data
// func (c *HistoryDataClient) GetHistoricalData(symbol string, from time.Time, to time.Time) ([]SmartApi.Candle, error) {
//     data, err := c.apiClient.GetHistoricalData(symbol, from, to)
//     if err != nil {
//         return nil, fmt.Errorf("error getting historical data: %v", err)
//     }
//     return data, nil
// }

func main() {
	creds := Credentials{
		ClientCode:    "V58985289",
		Password:      "1250",
		HistoryAPIKey: "NlvfasR9",
		TradingKey:    "2t2TvbKy",
		TOTPSecret:    "B3EHYOJOOCFEUVWK774H5GLC4I",
	}

	// historyDataClient := NewHistoryDataClient(creds)
	tradingClient := NewTradingClient(creds)

	// Use historyDataClient to retrieve historical data
	// historicalData, err := historyDataClient.GetHistoricalData("SBIN-EQ", time.Now().AddDate(0, 0, -7), time.Now())
	// if err != nil {
	//     fmt.Println("Error getting historical data:", err)
	// } else {
	//     fmt.Println("Historical Data:", historicalData)
	// }

	// Use tradingClient to perform trading operations
	if err := tradingClient.Login(); err != nil {
		fmt.Println("Error logging in:", err)
		return
	}

	if err := tradingClient.RenewAccessToken(); err != nil {
		fmt.Println("Error renewing access token:", err)
		return
	}

	// profile, err := tradingClient.GetUserProfile()
	// if err != nil {
	// 	fmt.Println("Error getting user profile:", err)
	// 	return
	// }

	// fmt.Println("User Profile :- ", profile)
	// fmt.Println("User Session Tokens :- ", tradingClient.session.UserSessionTokens)

	// orderParams := SmartApi.OrderParams{
	// 	Variety:         "NORMAL",
	// 	TradingSymbol:   "SBIN-EQ",
	// 	SymbolToken:     "3045",
	// 	TransactionType: "BUY",
	// 	Exchange:        "NSE",
	// 	OrderType:       "LIMIT",
	// 	ProductType:     "INTRADAY",
	// 	Duration:        "DAY",
	// 	Price:           "19500",
	// 	SquareOff:       "0",
	// 	StopLoss:        "0",
	// 	Quantity:        "1",
	// }

	// order, err := tradingClient.PlaceOrder(orderParams)
	// if err != nil {
	// 	fmt.Println("Error placing order:", err)
	// 	return
	// }

	// fmt.Println("Placed Order ID and Script :- ", order)

	holdings , err := tradingClient.apiClient.GetHoldings()
	if err != nil {
		fmt.Printf("err : %v" , err)
	}else{
		fmt.Println()
		for _ , holds := range holdings {
			fmt.Printf("%v\n",holds)
		}
	
	}



}
