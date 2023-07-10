package utils_adapters

import (
	"encoding/json"
	"log"

	"github.com/sirupsen/logrus"
	utils_protocols "github.com/vm-championships-manager/backend-service/internal/utils/protocols"
)

type LogAdapter struct {
	metadata interface{}
	log      *logrus.Logger
}

func NewLoggerAdapter() *LogAdapter {
	return &LogAdapter{
		log: logrus.New(),
	}
}

func (l *LogAdapter) Metadata(metadata map[string]interface{}) utils_protocols.Log {
	m, err := json.Marshal(metadata)
	if err != nil {
		log.Fatalln(err)
	}

	l.metadata = string(m)
	return l
}

func (l *LogAdapter) Error(msg string) {
	l.log.WithFields(logrus.Fields{
		"metadata": l.metadata,
	}).Error(msg)
}

func (l *LogAdapter) Warn(msg string) {
	l.log.WithFields(logrus.Fields{
		"metadata": l.metadata,
	}).Warn(msg)
}
