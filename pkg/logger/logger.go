package logger

import (
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path"
)

// Logger represents zerolog.Logger interface.
type Logger struct {
	*zerolog.Logger
}

// New returns a new logger instance.
func New(config *Config) *Logger {
	var writers []io.Writer

	if config.LogToConsole {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	}

	if config.LogToFile {
		writers = append(writers, newRollingFile(config))
	}

	mw := io.MultiWriter(writers...)

	logger := zerolog.New(mw).With().Timestamp().Logger()

	logger.Info().
		Bool("log_to_file", config.LogToFile).
		Bool("json_log", config.EncodeLogsAsJson).
		Str("log_dir", config.Directory).
		Str("filename", config.Filename).
		Int("max_size_mb", config.MaxSize).
		Int("max_backups", config.MaxBackups).
		Int("max_age_days", config.MaxAge).
		Msg("Logger has been configured.")

	return &Logger{
		Logger: &logger,
	}
}

func newRollingFile(config *Config) io.Writer {
	return &lumberjack.Logger{
		Filename:   path.Join(config.Directory, config.Filename),
		MaxBackups: config.MaxBackups,
		MaxSize:    config.MaxSize,
		MaxAge:     config.MaxAge,
	}
}
