package main

import (
	"flag"
	"log"

	"github.com/KudinovKV/converter/internal/coinmarketcap"
	"github.com/KudinovKV/converter/internal/config"
	"github.com/KudinovKV/converter/internal/converter"
	"github.com/KudinovKV/converter/internal/mapper"

	"github.com/shopspring/decimal"
)

func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) != 3 {
		log.Fatal("Usage: converter <amount> <source currency> <destination currency>")
	}

	amount, err := decimal.NewFromString(args[0])
	if err != nil {
		log.Fatalf("Parse amount failed: %v", err)
	}

	cfg := config.LoadConfig()

	coinMarketCapClient, err := coinmarketcap.New(cfg.CoinMarketCapAPIKey, cfg.RequestTimeout, mapper.New())
	if err != nil {
		log.Fatalf("Build CoinMarketCap client failed: %v", err)
	}

	converterManager := converter.New(coinMarketCapClient)

	resultAmount, err := converterManager.Convert(amount, args[1], args[2])
	if err != nil {
		log.Fatalf("Convert failed: %v", err)
	}

	log.Printf("%s %s is equal %s %s", amount, args[1], resultAmount.String(), args[2])
}
