package mapper

import (
	"github.com/KudinovKV/converter/internal/coinmarketcap"
	"github.com/KudinovKV/converter/internal/domain"

	"github.com/shopspring/decimal"
)

type Mapper struct{}

func (m *Mapper) CoinMarketCapToDomainPrices(data *coinmarketcap.Data) domain.Prices {
	result := make(map[string]decimal.Decimal, len(data.Quote))
	for currency, quote := range data.Quote {
		if quote.Price == nil {
			continue
		}

		result[currency] = decimal.NewFromFloat(*quote.Price)
	}

	return result
}

func New() *Mapper {
	return &Mapper{}
}
