package logging

import "github.com/sirupsen/logrus"

type Logger struct {
	prefix string
}

func NewLogger(prefix string) *Logger {
	return &Logger{prefix: prefix}
}

func (l *Logger) Trace(err error) {
	logrus.Tracef("%s | %s", l.prefix, err.Error())
}

func (l *Logger) Debug(err error) {
	logrus.Debugf("%s | %s", l.prefix, err.Error())
}

func (l *Logger) Info(err error) {
	logrus.Infof("%s | %s", l.prefix, err.Error())
}

func (l *Logger) Warn(err error) {
	logrus.Warnf("%s | %s", l.prefix, err.Error())
}

func (l *Logger) Error(err error) {
	logrus.Errorf("%s | %s", l.prefix, err.Error())
}

func (l *Logger) Fatal(err error) {
	logrus.Fatalf("%s | %s", l.prefix, err.Error())
}

func (l *Logger) Panic(err error) {
	logrus.Panicf("%s | %s", l.prefix, err.Error())
}
