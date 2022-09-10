package logger

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type loggerFactory struct {
	level string
}

// NewFactory creates a new logger factory instance
func NewFactory(level string) *loggerFactory {
	return &loggerFactory{level}
}

// Create creates a new logger receibing the log level as a parameter
func (l loggerFactory) Create() Log {
	return NewLogger(l.level)
}

type Log interface {
	Debug(messageFormat string, v ...interface{})
	Info(messageFormat string, v ...interface{})
	Error(message string, err ...error)
	Fatal(message string, err ...error)
	InfoWithFields(message string, fields map[string]interface{}, err ...error)
	AddFieldToContext(message string, field interface{})
}

type Logger struct {
	logger zerolog.Logger
}

// NewLogger take the level parameter to specify the setGlobalLevel for the logs
// and return an instance of the logger with a processId guid per context
func NewLogger(level string) *Logger {
	switch strings.ToLower(level) {
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	return &Logger{log.With().Str("processId", uuid.New().String()).Logger()}
}

// Debug logs a message to the console if log level is debug or lower
// most used to trace the process inside the application
func (l *Logger) Debug(messageFormat string, v ...interface{}) {
	event := l.logger.Debug()
	event.Msgf(messageFormat, v...)
}

// Info logs a message to the console if log level is info or lower
// show information of the execution
func (l *Logger) Info(messageFormat string, v ...interface{}) {
	event := l.logger.Info()
	event.Msgf(messageFormat, v...)
}

// Error logs a message to the console if log level is error or lower
// most used to send business errors
func (l *Logger) Error(message string, err ...error) {
	event := l.logger.Error()
	if err != nil {
		event = event.Err(errors.New(collect(err)))
	}
	event.Msg(message)
}

// Fatal logs a message to the console and stop the execution
// an error happens that the application can't continue with the execution
func (l *Logger) Fatal(message string, err ...error) {
	event := l.logger.Fatal()
	if err != nil {
		event = event.Err(errors.New(collect(err)))
	}
	event.Msg(message)
}

// collect is a helper function to concatenate errors messages
func collect(errors []error) string {
	var errorStrings []string
	for _, singleError := range errors {
		errorStrings = append(errorStrings, singleError.Error())
	}
	return strings.Join(errorStrings, " - ")
}

// InfoWithFields allows to send an information message to the console and
// add any custom fields, and errors to the message
func (l *Logger) InfoWithFields(message string, fields map[string]interface{}, err ...error) {
	event := l.logger.Info()
	if err != nil {
		event = event.Err(errors.New(collect(err)))
	}
	for key, value := range fields {
		switch val := value.(type) {
		case string:
			event = event.Str(key, val)
		case int:
			event = event.Int(key, val)
		case bool:
			event = event.Bool(key, val)
		case float32:
			event = event.Float32(key, val)
		case float64:
			event = event.Float64(key, val)
		}
	}
	event.Msg(message)
}

// AddFieldToContext add custom fields to the context so it's going to be present in all messages
// generated from the context
func (l *Logger) AddFieldToContext(fieldName string, field interface{}) {
	switch val := field.(type) {
	case string:
		l.logger = l.logger.With().Str(fieldName, val).Logger()
	case int:
		l.logger = l.logger.With().Int(fieldName, val).Logger()
	case uint:
		l.logger = l.logger.With().Uint(fieldName, val).Logger()
	case bool:
		l.logger = l.logger.With().Bool(fieldName, val).Logger()
	case float32:
		l.logger = l.logger.With().Float32(fieldName, val).Logger()
	case float64:
		l.logger = l.logger.With().Float64(fieldName, val).Logger()
	}
}
