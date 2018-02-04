package CrawlerEngine

import (
	"github.com/shopspring/decimal"
	"time"
)

type Crawler interface {
	IsConnected() bool
}

type CurrenciesMap struct {
	From []string `yaml:"from"`
	To   string   `yaml:"to"`
}

type Currency struct {
	Symbol string
	Name   string
}

type Price struct {
	Currency Currency
	Value    decimal.Decimal
}

type MarketPrice struct {
	Open decimal.Decimal
	High decimal.Decimal
	Low  decimal.Decimal
}

type MarketVolume map[Currency]Price

type Market struct {
	Price  MarketPrice
	Volume MarketVolume
}

type CryptoCurrencies map[Currency]Price

type SubsResp map[string]SubsRespContent

type SubsRespContent struct {
	Trades     SubsTradesResp  `json:"TRADES"`
	Current    SubsCurrentResp `json:"CURRENT"`
	CurrentAgg string          `json:"CURRENTAGG"`
}

type SubsTradesResp []string
type SubsMarketResp []string
type SubsCurrentResp []string

type CurrencyNotFoundError struct {
	error
}

type TradeAction int

const (
	BUY     TradeAction = 1
	SELL    TradeAction = 2
	UNKNOWN TradeAction = 4
)

type TradeRecord struct {
	Market     string
	Id         string
	Timestamp  time.Time
	Action     TradeAction
	Price      Price
	Quantity   decimal.Decimal
	TotalPrice Price
}

type MarketRecord struct {
	Day  Market
	Hour Market
}

type SubListRetrieveError struct {
	error
}
