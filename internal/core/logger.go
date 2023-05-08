package core

import (
	"log"
	"os"
)

type Logger struct {
	log *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		log: log.New(os.Stdout, "", log.LstdFlags),
	}
}
