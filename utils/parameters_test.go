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
	})
}
