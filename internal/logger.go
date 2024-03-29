package internal

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

var log *zerolog.Logger


func Get() *zerolog.Logger {

	if log != nil {
		return log
	}

	zerolog.TimeFieldFormat = time.RFC3339Nano
	var logLevel int

	config, err:= NewData()
	if err != nil {
		logLevel = int(zerolog.InfoLevel)
	} else {
		logLevel = config.LogLevel
	}

	var output io.Writer = zerolog.ConsoleWriter{
		Out:		os.Stdout,
		TimeFormat: time.RFC3339,
	}

	if config.LogOutput == "file" {
		file, err := os.OpenFile(
			config.LogFile,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0664,
		)

		if err == nil {
            output = zerolog.MultiLevelWriter(file)
		}
	}

	logger := zerolog.New(output).
		Level(zerolog.Level(logLevel)).
		With().
		Timestamp().
		Logger()

	log = & logger
	return log
}
