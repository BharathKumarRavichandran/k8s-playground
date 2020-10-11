package utils

import (
	"io"
	"os"
	"path"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger is an instance of logrus.Logger
// Logger is to be used for all logging
var Logger *logrus.Logger

// initLogger initializes the logger with appropriate configuration options
func initLogger(config *Config) {
	var (
		fileDir  = config.LogDir
		fileName = config.LogFileName
		maxSize  = config.LogMaxSize
		logLevel = config.LogLevel
	)

	if config.LogFileName == "" {
		fileName = "log.log"
	}

	filePath := path.Join(path.Dir(fileDir), fileName)

	if config.LogMaxSize == 0 {
		maxSize = 50
	}

	if config.LogLevel == "" {
		logLevel = "info"
	}

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		panic(err)
	}

	// Write to file always
	Logger = &logrus.Logger{
		Formatter: &logrus.JSONFormatter{},
		Out: &lumberjack.Logger{
			Filename: filePath,
			MaxSize:  maxSize, // MB
		},
		Level: level,
	}

	// Write to both stdout and file if DEBUG is true
	if config.DEBUG == true {
		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)
		if err == nil {
			Logger.Out = io.MultiWriter(os.Stdout, file)
		} else {
			Logger.Info("Failed to log to file, using default stderr")
		}
	}

	Logger.Info("Logger started")
}

// GetNewFileLogger returns a logger that writes to a file of the given name
func GetNewFileLogger(fileDir string, fileName string, maxSize int, logLevel string, json bool) *logrus.Logger {
	if fileName == "" {
		fileName = "log1.log"
	}
	filePath := path.Join(path.Dir(fileDir), fileName)

	if maxSize == 0 {
		maxSize = 50
	}

	if logLevel == "" {
		logLevel = "info"
	}

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		panic(err)
	}

	logger := &logrus.Logger{
		Out: &lumberjack.Logger{
			Filename: filePath,
			MaxSize:  maxSize, // MB
		},
		Level: level,
	}

	if json {
		logger.Formatter = &logrus.JSONFormatter{}
	} else {
		logger.Formatter = &logrus.TextFormatter{}
	}

	return logger
}
