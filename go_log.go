package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type LogEntry struct {
	Timestamp  string `json:"timestamp"`
	LogNumber  int    `json:"log_number"`
}

func createLogFile() {
	currentTime := time.Now()
	logFilename := "logtimes.json"

	logEntry := LogEntry{
		Timestamp: currentTime.Format("2006-01-02 15:04:05"),
		LogNumber: getLogNumber(logFilename),
	}

	logFile, err := os.OpenFile(logFilename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer logFile.Close()

	logEncoder := json.NewEncoder(logFile)
	if err := logEncoder.Encode(logEntry); err != nil {
		fmt.Println("Error writing log entry:", err)
		return
	}

	fmt.Printf("Log entry added to '%s'.\n", logFilename)
}

func getLogNumber(filename string) int {
	logEntries := []LogEntry{}
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return 1
		}
		fmt.Println("Error opening log file:", err)
		return 0
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	for decoder.More() {
		var entry LogEntry
		if err := decoder.Decode(&entry); err != nil {
			fmt.Println("Error decoding log entry:", err)
			return 0
		}
		logEntries = append(logEntries, entry)
	}

	return len(logEntries) + 1
}

func main() {
	createLogFile()
}