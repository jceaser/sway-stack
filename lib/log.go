/************************************************************************************************100
Project: sway-stack
Package: logs
File: log.go code to write logs
License: Private, use by explicit agreement only
Copyright 2025 - all rights reserved by thomas.cherry@gmail.com
***************************************************************************************************/

package lib

import (
	"io"
	"log"
	"os"
)

/* ************************************************************************** */
// MARK: - Data Types

// loggers
type LogType struct {
	Report *log.Logger
	Error  *log.Logger
	Warn   *log.Logger
	Info   *log.Logger
	Debug  *log.Logger
	Stats  *log.Logger
}

var Log LogType

func init() {
	file := os.Stderr
	settings := log.Ldate | log.Ltime | log.Lshortfile
	Log = LogType{
		Report: log.New(file, "REPORT: ", settings),
		Error:  log.New(file, "ERROR: ", settings),
		Warn:   log.New(file, "WARNING: ", settings),
		Info:   log.New(io.Discard, "INFO: ", settings),
		Debug:  log.New(io.Discard, "DEBUG: ", settings),
		Stats:  log.New(io.Discard, "STAT: ", settings)}
}

func (self *LogType) SetLevel(value, flag int, logger *log.Logger) {
	if value&flag == flag {
		logger.SetOutput(os.Stderr)
	} else {
		logger.SetOutput(io.Discard)
	}
}

func (self *LogType) Level(value int) {
	//self.SetLevel(value, 1, Log.Report) // should always be on
	self.SetLevel(value, 2, Log.Error)
	self.SetLevel(value, 4, Log.Warn)
	self.SetLevel(value, 8, Log.Info)
	self.SetLevel(value, 16, Log.Debug)
	//self.SetLevel(value, 32, Log.Stats) // should always be controlled by file function
}

func (self *LogType) EnableInfo() {
	self.Info.SetOutput(os.Stderr)
}

func (self *LogType) EnableDebug() {
	self.Debug.SetOutput(os.Stderr)
}

func (self *LogType) EnableStats(file *os.File) {
	self.Stats.SetOutput(file)
}
