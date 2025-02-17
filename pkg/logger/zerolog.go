package logger

import (
	"os"

	"github.com/rs/zerolog"
	logzerolog "github.com/rs/zerolog/log"
)

func initZerolog() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// logzerolog.Logger = logzerolog.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
	logzerolog.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

}

func GetZerologger() zerolog.Logger {
	initZerolog()
	return logzerolog.Logger
}
