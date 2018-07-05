package main

import (
	"fmt"
)

// SetLogLevel sets log Level (from 1 to 5, where 5 - show ALL and 1 - show only most important. Default is 3)
func SetLogLevel(level byte) {
	logLevel = level
}

// WriteLog writes to log
func WriteLog(message string, level byte) {
	if level <= logLevel {
		fmt.Println(message)
	}
}
