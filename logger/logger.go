package logger

import (
	"github.com/sirupsen/logrus"
)

// Log is the exported object for logging
var Log *logrus.Logger

type AnalyticsLogger interface {
	Logf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

type analyticsLogger struct{}

func (l *analyticsLogger) Logf(format string, args ...interface{}) {
	Log.Infof(format, args...)
}

func (l *analyticsLogger) Errorf(format string, args ...interface{}) {
	Log.Errorf(format, args...)
}

func NewAnalyticsLogger() AnalyticsLogger {
	return &analyticsLogger{}
}

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	Log = logrus.New()
}
