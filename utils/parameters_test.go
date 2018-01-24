package utils

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
)

func TestLoadParameters(t *testing.T) {
	t.Run("Load parameters yaml file", func(t *testing.T) {
		os.Setenv("config", "../config/test/parameters.yaml")
		parameters := LoadParameters()

		assert.Equal(t, "localhost", parameters.ServerConfig.Location,"Location error")
		assert.Equal(t, uint(5000), parameters.ServerConfig.Port,"Port error")

		assert.Equal(t, 2, len(parameters.CurrencyConfig.From))
		assert.Equal(t, "ETH", parameters.CurrencyConfig.From[0])
		assert.Equal(t, "XRP", parameters.CurrencyConfig.From[1])
		assert.Equal(t, "ToCurrency", parameters.CurrencyConfig.To)
	})
}
