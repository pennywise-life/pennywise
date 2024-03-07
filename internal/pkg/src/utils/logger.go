package utils

import "github.com/sirupsen/logrus"

type Log struct {
	log *logrus.Logger
}

func NewLog() *Log {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
		DisableColors: false,
	})
	return &Log{
		log: log,
	}
}
