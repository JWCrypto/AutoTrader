package AnalyseEngine

import (
	"github.com/JWCrypto/AutoTrader/CrawlerEngine"
	"github.com/JWCrypto/AutoTrader/utils/log"
	"fmt"
)

var logger = log.GetLogger()

type AnalyseEngine interface {
	Run()
}

type analyseEngine struct {
	health     bool
	liveTrade  chan *CrawlerEngine.TradeRecord
	liveMarket chan *CrawlerEngine.MarketRecord
}

func NewAnalyseEngine(liveTrade chan *CrawlerEngine.TradeRecord, liveMarket chan *CrawlerEngine.MarketRecord) *analyseEngine {
	return &analyseEngine{
		liveTrade:  liveTrade,
		liveMarket: liveMarket,
	}
}

func (engine *analyseEngine) Run() {
	logger.Infof("AnalyseEngine is running")

	go func() {
		for {
			fmt.Println(<-engine.liveTrade)
		}
	}()
}
