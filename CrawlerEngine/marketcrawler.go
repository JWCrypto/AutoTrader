package CrawlerEngine

import (
	"github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"time"
	"fmt"
)

type MarketCrawler interface {
	Crawler
	getSubscriptionsResp(from []Currency, to Currency) SubsMarketResp
	establishSocket(marketList SubsMarketResp, marketRecordChannel chan *MarketRecord)
	Run() chan *MarketRecord
}

type marketcrawler struct {
	fromCurrencies []Currency
	toCurrency     Currency
	connected      bool
}

func NewMarketCrawler(currenciesMap CurrenciesMap) *marketcrawler {
	var fromCurrencies = make([]Currency, 0, len(currenciesMap.From))

	for _, currency := range currenciesMap.From {
		fromCurrencies = append(fromCurrencies, Currency{Name: currency, Symbol: SYMBOL[currency]})
	}

	return &marketcrawler{
		fromCurrencies: fromCurrencies,
		toCurrency:     Currency{Name: currenciesMap.To, Symbol: SYMBOL[currenciesMap.To]},
	}
}

func (c *marketcrawler) getSubscriptionsResp(from []Currency, to Currency) SubsMarketResp {
	subsList := make([]string, 0, len(from))

	for _, currency := range from {
		subsList = append(subsList, fmt.Sprintf("5~CCCAGG~%s~%s", currency.Name, to.Name))
	}

	return subsList
}

func (c *marketcrawler) establishSocket(marketList SubsMarketResp, marketRecordChannel chan *MarketRecord) {
	client, err := gosocketio.Dial(
		gosocketio.GetUrl(STREAM_URL, 443, true),
		transport.GetDefaultWebsocketTransport(),
	)
	if err != nil {
		logger.Fatalf("Failed to dial %s, %v", STREAM_URL, err)
	}

	subSignal := map[string]SubsMarketResp{
		"subs": marketList,
	}
	client.Emit("SubAdd", subSignal)

	client.On("m", func(h *gosocketio.Channel, lastAction string) {
		if getType(lastAction) == TYPE["CURRENTAGG"] {
			record := toMarketRecord(lastAction)

			if record != nil {
				marketRecordChannel <- record
			}
		}
	})

	err = client.On(gosocketio.OnConnection, func(h *gosocketio.Channel) {
		c.connected = true
	})

	err = client.On(gosocketio.OnDisconnection, func(h *gosocketio.Channel) {
		c.connected = false
	})

	for {
		time.Sleep(MAX_DELAY)

		if !c.connected {
			close(marketRecordChannel)
			break
		}
	}

	defer client.Close()
}

func (c *marketcrawler) IsConnected() bool {
	return c.connected
}

func (c *marketcrawler) Run() chan *MarketRecord {
	liveMarketChannel := make(chan *MarketRecord)

	marketList := c.getSubscriptionsResp(c.fromCurrencies, c.toCurrency)
	go c.establishSocket(marketList, liveMarketChannel)

	return liveMarketChannel
}
