package converter

import (
	"fmt"

	"github.com/KudinovKV/converter/internal/domain"

	"github.com/shopspring/decimal"
)

type priceConverter interface {
	Convert(amount decimal.Decimal, currencyFrom string, currenciesTo []string) (domain.Prices, error)
}

// Converter is responsible for converting currencies. Converter uses priceConverter to get prices.
type Converter struct {
	priceConverter priceConverter
}

func (c *Converter) Convert(amount decimal.Decimal, currencyFrom, currencyTo string) (decimal.Decimal, error) {
	prices, err := c.priceConverter.Convert(amount, currencyFrom, []string{currencyTo})
	if err != nil {
		return decimal.Zero, err
	}

	if prices == nil {
		return decimal.Zero, fmt.Errorf("price for pair %s-%s not found", currencyFrom, currencyTo)
	}

	price, ok := prices[currencyTo]
	if !ok {
		return decimal.Zero, fmt.Errorf("price for pair %s-%s not found", currencyFrom, currencyTo)
	}

	return price, nil
}

func New(priceConverter priceConverter) *Converter {
	return &Converter{priceConverter}
}
