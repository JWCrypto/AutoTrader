package CrawlerEngine

import (
	"testing"
	"fmt"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestMarketcrawler_getSubscriptionsResp(t *testing.T) {
	t.Run("getSubscriptionsResp", func(t *testing.T) {
		connector := NewMarketCrawler(CurrenciesMap{})

		from := []Currency{
			{Name: "BTC"},
			{Name: "ETH"},
		}

		to := Currency{
			Name: "EUR",
		}

		resp := connector.getSubscriptionsResp(from, to)

		assert.Equal(t, 2, len(resp))
		assert.Equal(t, "5~CCCAGG~BTC~EUR", resp[0])
		assert.Equal(t, "5~CCCAGG~ETH~EUR", resp[1])
	})
}

func TestMarketCrawler_establishSocket(t *testing.T) {
	t.Run("establishSocket", func(t *testing.T) {
		connector := NewMarketCrawler(CurrenciesMap{})
		assert.False(t, connector.IsConnected())

		c := make(chan *MarketRecord)

		tradeList := SubsMarketResp{
			"5~CCCAGG~BTC~EUR",
			"5~CCCAGG~ETH~EUR",
		}
		go connector.establishSocket(tradeList, c)

		go func() {
			for {
				record := <-c
				fmt.Println(record)
				assert.True(t, connector.IsConnected())
			}
		}()
		time.Sleep(30 * time.Second)
	})
}
