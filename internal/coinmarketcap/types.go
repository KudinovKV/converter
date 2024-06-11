package coinmarketcap

import (
	"time"
)

type PriceConversionResponse struct {
	Data   []Data `json:"data,omitempty"`
	Status Status `json:"status"`
}

type Status struct {
	Timestamp    time.Time `json:"timestamp"`
	ErrorCode    int       `json:"error_code"`
	ErrorMessage string    `json:"error_message"`
	Elapsed      int       `json:"elapsed"`
	CreditCount  int       `json:"credit_count"`
	Notice       string    `json:"notice"`
}

type Data struct {
	Id          int              `json:"id"`
	Symbol      string           `json:"symbol"`
	Name        string           `json:"name"`
	Amount      float64          `json:"amount"`
	LastUpdated time.Time        `json:"last_updated"`
	Quote       map[string]Quote `json:"quote"`
}

type Quote struct {
	Price       *float64  `json:"price,omitempty"`
	LastUpdated time.Time `json:"last_updated"`
}
