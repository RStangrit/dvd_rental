package logger

import (
	"os"

	"github.com/rs/zerolog"
	logzerolog "github.com/rs/zerolog/log"
)

func initZerolog() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logzerolog.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

}

func GetZerologger() zerolog.Logger {
	initZerolog()
	return logzerolog.Logger
}
