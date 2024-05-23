package logger

import (
	"log"
	"os"
)

var Logger *log.Logger

func InitLogger() {
	Logger = log.New(os.Stdout, "", log.LstdFlags)
}

func Info(message string) {
	Logger.Printf("INFO:: %v", message)
}

func Error(message string) {
	Logger.Printf("ERROR:: %v", message)
}
