package log

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type LogInterface interface {
	Init()
	write(title string, format string, v ...any)
}

type Log struct {
	Path string
}

const DEFAULT_PATH string = "tmp/logs"

func New(Path ...string) *Log {

	var log *Log

	if len(Path) > 0 {
		log = &Log{Path: Path[0]}
	} else {
		log = &Log{Path: DEFAULT_PATH}
	}

	return log.Init()

}

func (l *Log) Init() *Log {

	// Buat folder log jika belum ada
	if err := os.MkdirAll(l.Path, 0755); err != nil {
		fmt.Println("github.com/golamee/logger: Failed to create a log directory:", err)
	}

	return l
}

func (l *Log) write(title string, format string, v ...any) {

	timeNow := time.Now().Local()
	date := timeNow.Format("2006-01-02")
	logFileName := filepath.Join(l.Path, date+".log")

	timestamp := timeNow.Format("2006-01-02 15:04:05 MST")
	logEntry := fmt.Sprintf("[%s] [%s] %s\n", timestamp, title, fmt.Sprintf(format, v...))

	if file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err == nil {
		defer file.Close()
		_, _ = file.WriteString(logEntry)
	}
}

func (l *Log) Info(Format string, v ...any) {
	l.write("INFO", Format, v...)
}

func (l *Log) Notice(Format string, v ...any) {
	l.write("NOTICE", Format, v...)
}

func (l *Log) Alert(Format string, v ...any) {
	l.write("ALERT", Format, v...)
}

func (l *Log) Warning(Format string, v ...any) {
	l.write("WARNING", Format, v...)
}

func (l *Log) Error(Format string, v ...any) {
	l.write("ERROR", Format, v...)
}

func (l *Log) Critical(Format string, v ...any) {
	l.write("CRITICAL", Format, v...)
}
