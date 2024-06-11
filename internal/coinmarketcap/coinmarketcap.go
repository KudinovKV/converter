package coinmarketcap

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/KudinovKV/converter/internal/domain"

	"github.com/shopspring/decimal"
)

const (
	url                     = "https://pro-api.coinmarketcap.com/"
	priceConversionEndpoint = "v2/tools/price-conversion"
	apiKeyHeader            = "X-CMC_PRO_API_KEY"

	defaultRequestTimeout = time.Second * 10
)

var errAPIKeyRequired = errors.New("api key is required")

type mapper interface {
	CoinMarketCapToDomainPrices(data *Data) domain.Prices
}

// Client is responsible for communication with CoinMarketCap.
type Client struct {
	apiKey         string
	requestTimeout time.Duration

	mapper mapper
}

func (c *Client) Convert(amount decimal.Decimal, currencyFrom string, currenciesTo []string) (domain.Prices, error) {
	req, err := http.NewRequest(http.MethodGet, url+priceConversionEndpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	q := req.URL.Query()
	q.Add("amount", amount.String())
	q.Add("symbol", currencyFrom)
	q.Add("convert", strings.Join(currenciesTo, ","))
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Accepts", "application/json")
	req.Header.Add(apiKeyHeader, c.apiKey)

	client := &http.Client{
		Timeout: c.requestTimeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var priceConversionResp PriceConversionResponse
	err = json.NewDecoder(resp.Body).Decode(&priceConversionResp)
	if err != nil {
		return nil, err
	}

	if priceConversionResp.Status.ErrorCode != 0 {
		return nil, errors.New(priceConversionResp.Status.ErrorMessage)
	}

	// TODO: must add network to specify element
	return c.mapper.CoinMarketCapToDomainPrices(&priceConversionResp.Data[0]), nil
}

func New(apiKey, rawRequestTimeout string, mapper mapper) (c *Client, err error) {
	if apiKey == "" {
		return nil, errAPIKeyRequired
	}

	requestTimeout := defaultRequestTimeout
	if rawRequestTimeout != "" {
		requestTimeout, err = time.ParseDuration(rawRequestTimeout)
		if err != nil {
			return nil, fmt.Errorf("parsing request timeout: %w", err)
		}
	}

	return &Client{
		apiKey:         apiKey,
		requestTimeout: requestTimeout,
		mapper:         mapper,
	}, nil
}
