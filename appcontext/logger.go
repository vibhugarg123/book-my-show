package appcontext

import (
	"github.com/rs/zerolog"
	"github.com/vibhugarg123/book-my-show/config"
	"os"
)

var Logger zerolog.Logger

func SetupLogger() {
	zerolog.SetGlobalLevel(getLogLevel())
	Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
}

func getLogLevel() zerolog.Level {
	level, err := zerolog.ParseLevel(config.LogLevel())
	if err != nil {
		return zerolog.InfoLevel
	}
	return level
}
