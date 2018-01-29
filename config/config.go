package config

import (
	"flag"
	"os"

	"github.com/jinzhu/configor"
)

// Config - config struct for this service.
type Config struct {
	ConfigPath          string
	HTTPAddr            string   `env:"HTTPAddr" default:"\":9167\""`
	DevEnv              bool     `env:"DEV_ENV" default:"false"`
	LogLevel            string   `env:"LOG_LEVEL" default:"info"`
	LogToFile           string   `env:"LOG_TO_FILE"`
	LogFormat           string   `env:"LOG_FORMAT" default:"\"json\""`
	FetchInterval       int64    `env:"FETCH_INTERVAL" default:"2000"`
	FetchTimeout        int64    `env:"FETCH_TIMEOUT" default:"10000"`
	ContainerLabelsList []string `env:"CONTAINER_LABELS_LIST"`
	EnableSelfMetrics   bool     `env:"ENABLE_SELF_METRICS" default:"false"`
}

// ParseConfig - Parse and return config for service.
func ParseConfig() Config {
	var config Config

	configorInstance := configor.New(&configor.Config{
		Debug:   false,
		Verbose: false,
	})
	configorInstance.Load(&config)

	loadConfig := func(filepath string) {
		if _, err := os.Stat(filepath + ".default"); err == nil {
			configorInstance.Load(&config, filepath+".default")
		}

		configorInstance.Load(&config, filepath)
	}

	filePath := flag.String("config.file", "", "path to config file")
	flag.Parse()
	if *filePath != "" {
		loadConfig(*filePath)
	}

	if filePath, ok := os.LookupEnv("CONFIG_FILE"); ok {
		loadConfig(filePath)
	}

	return config
}

// ConfigSingleton - Parsed config for this service.
var ConfigSingleton = ParseConfig()
