package main

import (
	"github.com/tinyhui/CryptoTrader/WebEngine"
	"github.com/tinyhui/CryptoTrader/utils"
	"github.com/tinyhui/CryptoTrader/CrawlerEngine"
	"github.com/tinyhui/CryptoTrader/AnalyseEngine"
)

func main() {
	parameters := utils.LoadParameters()

	// crawler engine loaded
	crawlerEngine := CrawlerEngine.NewCrawlerEngine(parameters.CurrencyConfig)
	liveTradeChannel, liveMarketChannel := crawlerEngine.Run()

	// analyse engine loaded
	analyseEngine := AnalyseEngine.NewAnalyseEngine(liveTradeChannel, liveMarketChannel)
	analyseEngine.Run()

	// web engine loaded
	webEngineBuilder := WebEngine.NewWebEngineBuilder().WithServerConfig(parameters.ServerConfig).WithStaticConfig(parameters.TemplateConfig)
	webEngineBuilder = webEngineBuilder.WithAnalyseEngine(crawlerEngine)

	webEngine := webEngineBuilder.Build()
	webEngine.Run() // run as blocked
}
