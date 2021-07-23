package logging

import "github.com/sirupsen/logrus"

type Logger struct {
	prefix string
}

func NewLogger(prefix string) *Logger {
	return &Logger{prefix: prefix}
}

func (l *Logger) Trace(message string) {
	logrus.Tracef("%s | %s", l.prefix, message)
}

func (l *Logger) Debug(message string) {
	logrus.Debugf("%s | %s", l.prefix, message)
}

func (l *Logger) Info(message string) {
	logrus.Infof("%s | %s", l.prefix, message)
}

func (l *Logger) Warn(message string) {
	logrus.Warnf("%s | %s", l.prefix, message)
}

func (l *Logger) Error(message string) {
	logrus.Errorf("%s | %s", l.prefix, message)
}

func (l *Logger) Fatal(message string) {
	logrus.Fatalf("%s | %s", l.prefix, message)
}

func (l *Logger) Panic(message string) {
	logrus.Panicf("%s | %s", l.prefix, message)
}
