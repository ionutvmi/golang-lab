package logger

import (
	"io"
	"log"
	"os"
)

var appLogFile *os.File

func init() {
	f, err := os.OpenFile("log_app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	writer := io.MultiWriter(os.Stdout, f)
	log.SetOutput(writer)
	log.SetPrefix("GO LAB: ")

	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)

	appLogFile = f
}

// Close is intended to be called from main
func Close() {
	appLogFile.Close()
}
