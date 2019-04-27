package clogger

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/hugefiver/GoPSP/conf"
)

const (
	DEBUG = iota
	INFO
	WARNING
	ERROR
)

var (
	isInited = false
	logPath  string
	logLevel int
)

func init() {
	if strings.HasPrefix(conf.Log.Path, "/") {
		logPath = conf.Log.Path
	} else {
		basePath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		logPath = filepath.Join(basePath, conf.Log.Path)
	}

	switch strings.ToLower(conf.Log.Level) {
	case "debug":
		logLevel = DEBUG
	case "info":
		logLevel = INFO
	case "warning":
		logLevel = WARNING
	case "error":
		logLevel = ERROR
	default:
		logLevel = WARNING
	}
	isInited = true
}
