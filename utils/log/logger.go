package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()
var initialed = false

func GetLogger() *logrus.Logger {
	if !initialed {
		log.Out = os.Stdout

		if os.Getenv("GIN_MODE") == "release" {
			log.SetLevel(logrus.InfoLevel)
			log.Info("Running in the release environment")
		} else {
			log.SetLevel(logrus.DebugLevel)
			log.Debug("Running in the debug environment")
		}
	}

	initialed = true
	return log
}
