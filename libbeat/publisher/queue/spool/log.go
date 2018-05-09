package spool

import (
	"fmt"
	"sync"

	"github.com/elastic/beats/libbeat/logp"
)

type logger interface {
	Debug(...interface{})
	Debugf(string, ...interface{})

	Info(...interface{})
	Infof(string, ...interface{})

	Error(...interface{})
	Errorf(string, ...interface{})
}

var _defaultLogger struct {
	singleton logger
	init      sync.Once
}

func defaultLogger() logger {
	_defaultLogger.init.Do(func() {
		_defaultLogger.singleton = logp.NewLogger("spool")
	})
	return _defaultLogger.singleton
}

// func defaultLogger() logger { return (*outLogger)(nil) }

type outLogger struct{}

func (l *outLogger) Debug(vs ...interface{})              { l.report("Debug", vs) }
func (l *outLogger) Debugf(fmt string, vs ...interface{}) { l.reportf("Debug: ", fmt, vs) }

func (l *outLogger) Info(vs ...interface{})              { l.report("Info", vs) }
func (l *outLogger) Infof(fmt string, vs ...interface{}) { l.reportf("Info", fmt, vs) }

func (l *outLogger) Error(vs ...interface{})              { l.report("Error", vs) }
func (l *outLogger) Errorf(fmt string, vs ...interface{}) { l.reportf("Error", fmt, vs) }

func (l *outLogger) report(level string, vs []interface{}) {
	args := append([]interface{}{level, ":"}, vs...)
	fmt.Println(args...)
}

func (*outLogger) reportf(level string, str string, vs []interface{}) {
	str = level + ": " + str
	fmt.Printf(str, vs...)
}
