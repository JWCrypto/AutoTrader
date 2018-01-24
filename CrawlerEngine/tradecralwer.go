package CrawlerEngine

import (
	"time"
	"net/http"
	"fmt"
	"encoding/json"
	"net"
	"github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

const (
	SUB_URL_BASE = "https://min-api.cryptocompare.com/data/subs"
	STREAM_URL   = "streamer.cryptocompare.com"
	MAX_TIMEOUT  = 10 * time.Second
	MAX_DELAY    = 5 * time.Second
)

// cc stands for CryptoCompare
type TradeCrawler interface {
	Crawler
	getSubscriptionsResp(from Currency, to Currency) (SubsResp, error)
	establishSocket(tradesList SubsTradesResp, tradeRecordChannel chan *TradeRecord)
	Run() chan *TradeRecord
}

type tradecrawler struct {
	httpClient     *http.Client
	fromCurrencies []Currency
	toCurrency     Currency
	connected      bool
}

func NewTradeCrawler(currenciesMap CurrenciesMap) *tradecrawler {
	var fromCurrencies = make([]Currency, 0, len(currenciesMap.From))

	for _, currency := range currenciesMap.From {
		fromCurrencies = append(fromCurrencies, Currency{Name: currency, Symbol: SYMBOL[currency]})
	}

	return &tradecrawler{
		fromCurrencies: fromCurrencies,
		toCurrency:     Currency{Name: currenciesMap.To, Symbol: SYMBOL[currenciesMap.To]},
		httpClient: &http.Client{
			Timeout: MAX_TIMEOUT,
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					Timeout: MAX_TIMEOUT,
				}).Dial,
				TLSHandshakeTimeout: MAX_TIMEOUT,
			},
		},
		connected: false,
	}
}

func getSubsUrl(fsym string, tsym string) string {
	return fmt.Sprintf("%s?fsym=%s&tsyms=%s", SUB_URL_BASE, fsym, tsym)
}

func (c *tradecrawler) getSubscriptionsResp(from Currency, to Currency) (SubsResp, error) {
	url := getSubsUrl(from.Name, to.Name)
	logger.Debugf("Retrieving subscription list from %s", url)
	var resp = new(SubsResp)

	r, err := c.httpClient.Get(url)
	if err != nil {
		logger.Error("Not able to get subscription list from crypto compare")
		return *resp, err
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(resp)

	if !validSubsResp(*resp, to) {
		logger.Error("Trader list should not empty")
		return *resp, &SubListRetrieveError{}
	}

	return *resp, nil
}

func (c *tradecrawler) establishSocket(tradeList SubsTradesResp, tradeRecordChannel chan *TradeRecord) {
	client, err := gosocketio.Dial(
		gosocketio.GetUrl(STREAM_URL, 443, true),
		transport.GetDefaultWebsocketTransport(),
	)
	if err != nil {
		logger.Fatalf("Failed to dial %s, %v", STREAM_URL, err)
	}

	subSignal := map[string]SubsTradesResp{
		"subs": tradeList,
	}
	client.Emit("SubAdd", subSignal)

	client.On("m", func(h *gosocketio.Channel, lastAction string) {
		if getType(lastAction) == TYPE["TRADE"] {
			record := toTradeRecord(lastAction)

			if record != nil {
				tradeRecordChannel <- record
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
			close(tradeRecordChannel)
			break
		}
	}

	defer client.Close()
}

func (c *tradecrawler) IsConnected() bool {
	return c.connected
}

func (c *tradecrawler) Run() chan *TradeRecord {
	liveTradeChannel := make(chan *TradeRecord)

	for _, fromCurrency := range c.fromCurrencies {
		subsResp, err := c.getSubscriptionsResp(fromCurrency, c.toCurrency)
		if err != nil {
			logger.Fatalf("Cannot retrieve subscriptions list, %v", err)
		}
		tradesList := subsResp[c.toCurrency.Name].Trades
		logger.Infof("Retrieved subscriptions list for %s", fromCurrency.Name)

		logger.Infof("Retrieving trade log from CryptoCompare for %s", fromCurrency.Name)
		go c.establishSocket(tradesList, liveTradeChannel)
	}

	return liveTradeChannel
}
