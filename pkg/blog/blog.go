// blog is generic logger used for this project

package blog

import (
	"github.com/sirupsen/logrus"
)

// Level :
type Level int8

const (
	// DebugLevel  : by - default debug level if off
	DebugLevel Level = iota - 1

	// InfoLevel default log level
	InfoLevel
)

// Logger represents logger
// it just interface
// implantation TODO
type Logger interface {
	Infof(string, ...interface{})
	Debugf(string, ...interface{})
	Errorf(string, ...interface{})
}

// New Create blog logger
func New(lvl Level) Logger {
	l := logrus.New()
	var level logrus.Level
	if lvl == DebugLevel {
		level = logrus.DebugLevel
	} else {
		level = logrus.InfoLevel
	}
	l.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
		PadLevelText:     true,
	})
	l.SetLevel(level)
	return &logger{l}
}

// NewWithFields :
func NewWithFields(l Logger, fields map[string]interface{}) Logger {
	return l.(*logger).WithFields(logrus.Fields(fields))
}

type logger struct {
	*logrus.Logger
}
