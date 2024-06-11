package config

import (
	"os"
)

const (
	apiKey         = "CONVERTER_COIN_MARKET_CAP_API_KEY"
	requestTimeout = "CONVERTER_COIN_MARKET_CAP_REQUEST_TIMEOUT"
)

type Config struct {
	CoinMarketCapAPIKey string
	RequestTimeout      string
}

func LoadConfig() *Config {
	return &Config{
		CoinMarketCapAPIKey: os.Getenv(apiKey),
		RequestTimeout:      os.Getenv(requestTimeout),
	}
}
