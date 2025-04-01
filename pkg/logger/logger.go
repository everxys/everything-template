package logger

import (
	"sync"
)

func (l *Logger) SetLogger(in LoggerInterface) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.logger = in
}

type Logger struct {
	lock   sync.Mutex
	logger LoggerInterface
}

func NewLogger(level Level) *Logger {
	return newlogger(level)
}

func (l *Logger) Debugw(msg string, keysAndValues ...interface{}) {
	l.logger.Debugw(msg, keysAndValues...)
}

func (l *Logger) Errorw(msg string, keysAndValues ...interface{}) {
	l.logger.Errorw(msg, keysAndValues...)
}

func (l *Logger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.logger.Fatalw(msg, keysAndValues...)
}

func (l *Logger) Infow(msg string, keysAndValues ...interface{}) {
	l.logger.Infow(msg, keysAndValues...)
}

func (l *Logger) SetLevel(level Level) {
	l.logger.SetLevel(level)
}

func (l *Logger) With(keyValues ...interface{}) LoggerInterface {
	newLog := l.logger.With(keyValues...)
	return &Logger{logger: newLog}
}
