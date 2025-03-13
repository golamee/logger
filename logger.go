package logger

import "github.com/golamee/logger/log"

var l *log.Log

func NewLog(Path ...string) *log.Log {
	return log.New(Path...)
}

func Init(Path ...string) *log.Log {
	l = log.New(Path...)

	return l
}

func Info(Format string, v ...any) {
	l.Info(Format, v...)
}

func Notice(Format string, v ...any) {
	l.Notice(Format, v...)
}

func Alert(Format string, v ...any) {
	l.Alert(Format, v...)
}

func Warning(Format string, v ...any) {
	l.Warning(Format, v...)
}

func Error(Format string, v ...any) {
	l.Error(Format, v...)
}

func Critical(Format string, v ...any) {
	l.Critical(Format, v...)
}
