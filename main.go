package main

import l "github.com/golamee/logger/log"

func NewLog(Path ...string) *l.Log {

	var log *l.Log

	if len(Path) > 0 {
		log = &l.Log{Path: Path[0]}
	} else {
		log = &l.Log{Path: "/tmp/logs/"}
	}

	return log.Init()

}
