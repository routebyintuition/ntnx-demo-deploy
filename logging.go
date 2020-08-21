package main

import (
	"io"
	"log"
	"os"
)

var (
	// Info is informational logs
	Info *log.Logger
	// Warning are warning-level logs
	Warning *log.Logger
	// Error is error-level logs
	Error *log.Logger
)

func logInit(logLocation string) {

	logFileLocation := logLocation + "app_logs"

	//fmt.Println("Application log file: ", logFileLocation)

	logFile, err := os.OpenFile(logFileLocation, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file ", logFileLocation, ":", err)
	}

	infoHandle := io.MultiWriter(logFile, os.Stdout)
	warningHandle := io.MultiWriter(logFile, os.Stdout)
	errorHandle := io.MultiWriter(logFile, os.Stdout)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime)
}
