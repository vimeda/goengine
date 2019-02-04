package zap

import (
	goengine_dev "github.com/hellofresh/goengine-dev"
	"go.uber.org/zap"
)

var _ goengine_dev.Logger = &Wrapper{}

// Wrapper a struct that embeds the zap.Logger in order to implement log.Logger
type Wrapper struct {
	*zap.Logger
}

// Wrap wraps a zap.Logger
func Wrap(logger *zap.Logger) *Wrapper {
	return &Wrapper{logger}
}

// Error writes a log with log level error
func (w *Wrapper) Error(msg string) {
	w.Logger.Error(msg)
}

// Warn writes a log with log level warning
func (w *Wrapper) Warn(msg string) {
	w.Logger.Warn(msg)
}

// Info writes a log with log level info
func (w *Wrapper) Info(msg string) {
	w.Logger.Info(msg)
}

// Debug writes a log with log level debug
func (w *Wrapper) Debug(msg string) {
	w.Logger.Debug(msg)
}

// WithField Adds a field to the log entry
func (w *Wrapper) WithField(key string, val interface{}) goengine_dev.Logger {
	return Wrap(w.Logger.With(zap.Any(key, val)))
}

//WithFields Adds a set of fields to the log entry
func (w *Wrapper) WithFields(fields goengine_dev.Fields) goengine_dev.Logger {
	zapFields := make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		zapFields = append(zapFields, zap.Any(k, v))
	}

	return Wrap(w.Logger.With(zapFields...))
}

// WithError Add an error as single field to the log entry
func (w *Wrapper) WithError(err error) goengine_dev.Logger {
	return Wrap(w.Logger.With(zap.Error(err)))
}