package main

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

// Logger provides contextual formated json logs for consumption by humans
// and machines. It is powered by zerolog.
type Logger struct {
	log zerolog.Logger
}

// logWriter uses the APP_ENV value to determine the io.Writer to use with
// zerolog.ConsoleWriter. By default it writes logs to os.Stdout, but will use
// io.Discard (/dev/null effectively) if env is set to "test".
func logWriter(env string) zerolog.ConsoleWriter {
	var output io.Writer
	switch os.Getenv("APP_ENV") {
	case "test":
		output = io.Discard
	default:
		output = os.Stdout
	}

	return zerolog.ConsoleWriter{Out: output, TimeFormat: time.RFC3339Nano}
}

// NewLogger returns a new pointer to Logger.
func NewLogger(conf Config) *Logger {
	return &Logger{
		log: zerolog.New(logWriter(conf.AppEnv)).
			With().Timestamp().
			Str("APP_NAME", conf.AppName).
			Str("APP_ENV", conf.AppEnv).
			Logger(),
	}
}

// LogError writes to the error log when err is not nil. It also returns the
// error, so it can be chained with other error handling code.
func (logger *Logger) LogError(err error, msg string) error {
	if err != nil {
		logger.log.Error().Err(err).Msg(msg)
	}

	return err
}
