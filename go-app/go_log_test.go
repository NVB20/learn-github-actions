package main

import (
	"encoding/json"
	"os"
	"testing"
	"time"
)

func TestGetLogNumberFileNotFound(t *testing.T) {
	filename := "nonexistent_log.json"
	logNumber := getLogNumber(filename)
	if logNumber != 1 {
		t.Errorf("Expected log number to be 1, got %d", logNumber)
	}
}

func TestGetLogNumberExistingFile(t *testing.T) {
	// Create a temporary log file for testing
	tempLogFile := "temp_log.json"
	defer os.Remove(tempLogFile)
	createTempLogFile(tempLogFile)

	logNumber := getLogNumber(tempLogFile)
	expectedLogNumber := 2
	if logNumber != expectedLogNumber {
		t.Errorf("Expected log number to be %d, got %d", expectedLogNumber, logNumber)
	}
}

func createTempLogFile(filename string) {
	currentTime := time.Now()
	logEntry := LogEntry{
		Timestamp: currentTime.Format("2006-01-02 15:04:05"),
		LogNumber: 1,
	}

	logFile, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	defer logFile.Close()

	logEncoder := json.NewEncoder(logFile)
	logEncoder.Encode(logEntry)
}

func TestCreateLogFile(t *testing.T) {
	// Call the function and check if it runs without errors
	createLogFile()
}