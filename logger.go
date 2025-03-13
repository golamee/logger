package logger

import "github.com/golamee/logger/log"

func NewLog(Path ...string) {
	log.New(Path...)
}
