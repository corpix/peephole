package logger

import (
	"github.com/corpix/loggers"
	"github.com/corpix/loggers/logger/logrus"
)

type Config logrus.Config

func FromConfig(c Config) (loggers.Logger, error) {
	return logrus.NewFromConfig(logrus.Config(c))
}
