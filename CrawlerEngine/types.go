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

type HistoryPrice struct {
	Open  decimal.Decimal
	Close decimal.Decimal
	HIGH  decimal.Decimal
	LOW   decimal.Decimal
}

type Price struct {
	Currency Currency
	Value    decimal.Decimal
}

type CryptoCurrencies map[Currency]Price

type SubsResp map[string]SubsRespContent

type SubsRespContent struct {
	Trades     SubsTradesResp  `json:"TRADES"`
	Current    SubsCurrentResp `json:"CURRENT"`
	CurrentAgg string          `json:"CURRENTAGG"`
}

type SubsTradesResp []string
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

type SubListRetrieveError struct {
	error
}
