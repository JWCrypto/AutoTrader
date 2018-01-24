package CrawlerEngine

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
	"fmt"
)

func Test_getSubsUrl(t *testing.T) {
	t.Run("getSubsUrl", func(t *testing.T) {
		assert.Equal(t, "https://min-api.cryptocompare.com/data/subs?fsym=ETH&tsyms=ToCurrency", getSubsUrl("ETH", "ToCurrency"))
	})
}

func TestTradeCrawler_getSubscriptionsResp(t *testing.T) {
	t.Run("getSubscriptionsResp", func(t *testing.T) {
		connector := NewTradeCrawler(CurrenciesMap{})
		resp, _ := connector.getSubscriptionsResp(Currency{Name: "ETH"}, Currency{Name: "EUR"})

		assert.Equal(t, 14, len(resp["EUR"].Trades))
	})
}

func TestTradeCrawler_establishSocket(t *testing.T) {
	t.Run("establishSocket", func(t *testing.T) {
		connector := NewTradeCrawler(CurrenciesMap{})
		assert.False(t, connector.IsConnected())

		c := make(chan *TradeRecord)

		tradeList := SubsTradesResp{
			"0~Bitstamp~ETH~EUR",
			"0~Coinbase~ETH~EUR",
			"0~Cexio~ETH~EUR",
			"0~BTCE~ETH~EUR",
			"0~Kraken~ETH~EUR",
			"0~HitBTC~ETH~EUR",
			"0~Gatecoin~ETH~EUR",
			"0~Exmo~ETH~EUR",
			"0~BitBay~ETH~EUR",
			"0~TheRockTrading~ETH~EUR",
			"0~Quoine~ETH~EUR",
			"0~WavesDEX~ETH~EUR",
			"0~Lykke~ETH~EUR",
			"0~Coinroom~ETH~EUR",
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
