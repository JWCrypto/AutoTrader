package CrawlerEngine

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/shopspring/decimal"
)

func TestNewCrawlerEngine(t *testing.T) {
	t.Run("New CrawlerEngine", func(t *testing.T) {
		config := CurrenciesMap{
			From: []string{"ETH", "XRP"},
			To:   "ToCurrency",
		}

		engine := NewCrawlerEngine(config)
		currencies := *engine.cryptoCurrencies

		assert.Equal(t, decimal.Zero, currencies[Currency{Name: "ETH", Symbol: SYMBOL["ETH"]}].Value)
		assert.Equal(t, decimal.Zero, currencies[Currency{Name: "XRP", Symbol: SYMBOL["XRP"]}].Value)
	})
}

func TestCrawlerEngine_GetCurrentPrice(t *testing.T) {
	t.Run("New CrawlerEngine", func(t *testing.T) {
		config := CurrenciesMap{
			From: []string{"ETH", "XRP"},
			To:   "ToCurrency",
		}

		engine := NewCrawlerEngine(config)

		var e error
		_, e = engine.GetCurrentPrice("ETH")
		assert.Nil(t, e, "Error not nil")
		_, e = engine.GetCurrentPrice("BTC")
		assert.NotNil(t, e, "Expected error")
	})
}
