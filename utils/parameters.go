package utils

import (
	"os"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"github.com/tinyhui/CryptoTrader/WebEngine"
)

type Parameters struct {
	ServerConfig   WebEngine.ServerConfig   `yaml:"server"`
	TemplateConfig WebEngine.TemplateConfig `yaml:"public"`
}

func LoadParameters() Parameters {
	parametersFile := os.Getenv("config")
	if parametersFile == "" {
		log.Fatal("Config file path missing")
	}

	yamlFile, err := ioutil.ReadFile(parametersFile)
	if err != nil {
		log.Fatalf("configFile %s .Get err #%v", parametersFile, err)
	}

	parameters := Parameters{}
	err = yaml.Unmarshal(yamlFile, &parameters)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return parameters
}
