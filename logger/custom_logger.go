package logger

import "log"

func Log(msg string) {
	log.SetPrefix("[CLOG] ")
	log.Println(msg)
}
