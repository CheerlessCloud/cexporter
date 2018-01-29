package logger

import (
	"fmt"
	"os"

	configPackage "github.com/CheerlessCloud/cexporter/config"

	logrus "github.com/CheerlessCloud/logrus"
)

var config = configPackage.ConfigSingleton

// Log - global logger for this service
var Log *logrus.Logger

func init() {
	if config.LogFormat == "json" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else if config.LogFormat == "text" {
		logrus.SetFormatter(&logrus.TextFormatter{})
	} else {
		fmt.Println("Invalid logrus format '" + config.LogFormat + "'")
		fmt.Println("Set by default to JSON")
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}

	if config.LogToFile == "" {
		logrus.SetOutput(os.Stdout)
	} else {
		f, err := os.OpenFile(config.LogToFile, os.O_WRONLY|os.O_APPEND, 0755)
		if err != nil {
			fmt.Println(fmt.Errorf("Error on open logrus file: %#+v", err))
		} else {
			logrus.SetOutput(f)
		}
	}

	lvl, _ := logrus.ParseLevel(config.LogLevel)
	logrus.SetLevel(lvl)

	Log = logrus.New()
	Log.SetLevel(lvl)
}
