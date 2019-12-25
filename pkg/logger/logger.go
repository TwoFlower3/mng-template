package logger

import (
	"github.com/sirupsen/logrus"
)

// LogFormat dummy
type LogFormat string

// Logger struct
type Logger struct {
	*logrus.Entry
	log    *logrus.Logger
	Format LogFormat
}

const (
	// JSONLogFormat dummy
	JSONLogFormat LogFormat = "json"
	// TextLogFormat dummy
	TextLogFormat LogFormat = "text"
)

// New initialize logger
func New(version, defaultVersion string) *Logger {
	if version == "" {
		version = defaultVersion
	}
	log := logrus.New()

	entry := log.WithFields(logrus.Fields{
		"Version": version,
	})

	return &Logger{
		Entry:  entry,
		log:    log,
		Format: TextLogFormat,
	}
}

// SetLogMode debug/info
func (l *Logger) SetLogMode(level logrus.Level) {
	l.log.SetLevel(level)
}

// EnableTrace dummy
func (l *Logger) EnableTrace(trace bool) {
	l.log.SetReportCaller(trace)
}

// SetLogFormatter text/json
func (l *Logger) SetLogFormatter(format LogFormat) {
	switch format {
	case TextLogFormat:
		l.log.SetFormatter(&logrus.TextFormatter{})
	case JSONLogFormat:
		l.log.SetFormatter(&logrus.JSONFormatter{})
	}
	l.Format = format
}

// GetLevel dummy
func (l *Logger) GetLevel() logrus.Level {
	return l.log.GetLevel()
}
