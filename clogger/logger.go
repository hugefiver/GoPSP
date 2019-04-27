package clogger

import (
	"log"
	"os"
)

func mkLog(prefix string, i ...interface{}) {
	var logFile *os.File

	if logPath == "" {
		logFile = os.Stderr
	} else {
		logFile, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE, 0660)
		if err != nil {
			log.Println("[Error] Open log file failed.")
		} else {
			defer logFile.Close()
		}
	}
	logger := log.New(logFile, prefix, log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println(i)
}

func Debug(i ...interface{}) {
	if logLevel <= DEBUG {
		mkLog("[Debug]", i)
	}
}

func Info(i ...interface{}) {
	if logLevel <= INFO {
		mkLog("[Info]", i)
	}
}

func Warning(i ...interface{}) {
	if logLevel <= WARNING {
		mkLog("[Warning]", i)
	}
}

func Error(i ...interface{}) {
	if logLevel <= ERROR {
		mkLog("[Error]", i)
	}
}
