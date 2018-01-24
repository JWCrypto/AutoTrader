package CrawlerEngine

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/shopspring/decimal"
)

func Test_toTradeRecord(t *testing.T) {
	t.Run("getSubsUrl", func(t *testing.T) {
		action := "0~Coinroom~ETH~EUR~1~15171648710001~1517164871~0.7962174~833.44~663.599429856~1f"

		record := toTradeRecord(action)

		assert.Equal(t, "Coinroom", record.Market)
		assert.Equal(t, "15171648710001", record.Id)
		assert.Equal(t, "ETH", record.Price.Currency.Name)
		assert.Equal(t, SYMBOL["ETH"], record.Price.Currency.Symbol)
		assert.Equal(t, decimal.NewFromFloat(833.44), record.Price.Value)
		assert.Equal(t, BUY, record.Action)
		assert.Equal(t, decimal.NewFromFloat(0.7962174), record.Quantity)
		assert.Equal(t, Currency{SYMBOL["ETH"], "ETH"}, record.TotalPrice.Currency)
		assert.Equal(t, decimal.NewFromFloat(663.599429856), record.TotalPrice.Value)
	})
}
