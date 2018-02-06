package CrawlerEngine

import (
	"time"
	"github.com/JWCrypto/AutoTrader/utils/log"
	"github.com/shopspring/decimal"
)

var logger = log.GetLogger()

type CrawlerEngine interface {
	GetCurrentPrice(currency string) (*Price, error)
	IsHealth() bool
	Run() (chan *TradeRecord, chan *MarketRecord)
}

type crawlerEngine struct {
	health           bool
	cryptoCurrencies *CryptoCurrencies
	latestUpdate     time.Time
	tradeCrawler     TradeCrawler
	marketCrawler    MarketCrawler
}

func NewCrawlerEngine(config CurrenciesMap) *crawlerEngine {
	location, err := time.LoadLocation("UTC")
	if err != nil {
		logger.Fatalln("Time Location is incorrect")
	}

	cryptoCurrencies := CryptoCurrencies{}

	for _, currencyName := range config.From {
		cryptoCurrencies[Currency{Name: currencyName, Symbol: SYMBOL[currencyName]}] = Price{
			Currency: Currency{},
			Value:    decimal.Zero,
		}
	}

	engine := &crawlerEngine{
		health:           false,
		cryptoCurrencies: &cryptoCurrencies,
		latestUpdate:     time.Now().In(location),
		tradeCrawler:     NewTradeCrawler(config),
		marketCrawler:    NewMarketCrawler(config),
	}
	return engine
}

func (engine *crawlerEngine) hearBeat(duration time.Duration) {
	defer func() { engine.health = false }()
	for {
		time.Sleep(duration)

		if engine.tradeCrawler.IsConnected() && engine.marketCrawler.IsConnected() {
			logger.Debug("Crawler Engine Heart beat")
			engine.health = true
		} else {
			logger.Error("Socket with CryptoCompare disconnected")
			engine.health = false
		}
	}
}

func (engine *crawlerEngine) GetCurrentPrice(currency string) (*Price, error) {
	if price, ok := (*engine.cryptoCurrencies)[Currency{Name: currency, Symbol: SYMBOL[currency]}]; ok {
		return &price, nil
	}

	return &Price{
		Currency: Currency{},
		Value:    decimal.Zero,
	}, CurrencyNotFoundError{}
}

func (engine *crawlerEngine) IsHealth() bool {
	return engine.health
}

func (engine *crawlerEngine) Run() (chan *TradeRecord, chan *MarketRecord) {
	tradeRecord := engine.tradeCrawler.Run()
	marketRecord := engine.marketCrawler.Run()
	go engine.hearBeat(30 * time.Second)
	return tradeRecord, marketRecord
}
