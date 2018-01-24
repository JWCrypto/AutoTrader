package utils

import (
	"os"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"github.com/tinyhui/CryptoTrader/WebEngine"
	"github.com/tinyhui/CryptoTrader/utils/log"
	"github.com/tinyhui/CryptoTrader/CrawlerEngine"
)

var logger = log.GetLogger()

type Parameters struct {
	ServerConfig   WebEngine.ServerConfig      `yaml:"server"`
	TemplateConfig WebEngine.TemplateConfig    `yaml:"public"`
	CurrencyConfig CrawlerEngine.CurrenciesMap `yaml:"tradeCurrencies"`
}

func LoadParameters() Parameters {
	parametersFile := os.Getenv("config")
	if parametersFile == "" {
		logger.Fatalln("Config file path missing")
	}

	yamlFile, err := ioutil.ReadFile(parametersFile)
	if err != nil {
		logger.Fatalf("configFile %s .Get err #%v", parametersFile, err)
	}

	parameters := Parameters{}
	err = yaml.Unmarshal(yamlFile, &parameters)
	if err != nil {
		logger.Fatalf("Unmarshal: %v", err)
	}

	return parameters
}
