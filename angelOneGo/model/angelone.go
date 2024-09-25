package model

import (
	SmartApi "github.com/angel-one/smartapigo"
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
	ApiClient *SmartApi.Client
	Creds     Credentials
}

// TradingClient struct for handling trading operations
type TradingClient struct {
	ApiClient *SmartApi.Client
	Session   *SmartApi.UserSession
	Creds     Credentials
}
