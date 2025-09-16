package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func InitLogs() {
	logFile, err := os.OpenFile("proxy.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("failed to open log file:", err)
		return
	}

	multiWriter := io.MultiWriter(os.Stdout, logFile)

	log.SetOutput(multiWriter)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
