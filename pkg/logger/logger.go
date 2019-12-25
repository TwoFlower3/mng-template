package logger

import (
	"github.com/sirupsen/logrus"
)

type LogFormat string

type Logger struct {
	*logrus.Entry
	log    *logrus.Logger
	Format LogFormat
}

const (
	JSONLogFormat LogFormat = "json"
	TextLogFormat LogFormat = "text"
)

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

func (l *Logger) SetLogMode(level logrus.Level) {
	l.log.SetLevel(level)
}

func (l *Logger) EnableTrace(trace bool) {
	l.log.SetReportCaller(trace)
}

func (l *Logger) SetLogFormatter(format LogFormat) {
	switch format {
	case TextLogFormat:
		l.log.SetFormatter(&logrus.TextFormatter{})
	case JSONLogFormat:
		l.log.SetFormatter(&logrus.JSONFormatter{})
	}
	l.Format = format
}

func (l *Logger) GetLevel() logrus.Level {
	return l.log.GetLevel()
}
