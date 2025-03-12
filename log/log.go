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

func New() *Log {
	log := &Log{
		Path: "tmp/logs/",
	}

	log.Init()

	return log
}

func (l *Log) Init() {

	// Buat folder log jika belum ada
	if err := os.MkdirAll(l.Path, 0755); err != nil {
		fmt.Println("Gagal membuat direktori log:", err)
	}
}

func (l *Log) write(title string, format string, v ...any) {

	// Gunakan timezone dari perangkat atau default ke UTC+7 jika tidak ditemukan
	timeNow := time.Now().Local()
	date := timeNow.Format("2006-01-02")
	logFileName := filepath.Join(l.Path, date+".log")

	// Format log
	timestamp := timeNow.Format("2006-01-02 15:04:05 MST")
	logEntry := fmt.Sprintf("[%s] [%s] %s\n", timestamp, title, fmt.Sprintf(format, v...))

	// Gunakan file dalam mode append
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
