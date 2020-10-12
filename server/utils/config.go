package utils

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

// Config contains all the configuration options
type Config struct {
	// Environment related options
	// STAGE is the current execution environment. Can be one of "prod", "dev" or "test"
	STAGE string
	DEBUG bool
	HOST  string
	PORT  string
	// ServerPort is the port inwhich the application is running
	ServerPort string

	// Database related options
	// DB_HOST is the host name of the database server
	DB_HOST string
	// DB_PORT is the port of the database server
	DB_PORT int
	// DB_KEYSPACE is the name of the database
	DB_KEYSPACE string
	// DB_USERNAME is the name of the database user
	DB_USERNAME string
	// DB_PASSWORD is the password of the database user
	DB_PASSWORD string

	// KAFKA_SERVICE_NAME is the list of bootstrap servers / brokers
	KAFKA_SERVICE_NAME string
	// KAFKA_CONSUMER_GROUP is the consumer group that this app is part of
	KAFKA_CONSUMER_GROUP string
	// KAFKA_TOPIC is the topic that the app uses to push/read records
	KAFKA_TOPIC    string
	KAFKA_USERNAME string
	KAFKA_PASSWORD string

	// Logging related options
	// LogDir is the path of the log directory
	LogDir string
	// LogFileName is the name of the log file name
	LogFileName string
	// LogMaxSize is the maximum size(MB) of a log file before it gets rotated
	LogMaxSize int
	// LogLevel determines the log level.
	// Can be one of "debug", "info", "warn", "error"
	LogLevel string
}

// Struct to load configurations of all possible modes i.e dev, docker, prod, test
// Only one of them will be selected based on the environment variable ENV
var allConfigurations = struct {

	// Configuration for environment : dev
	Dev Config

	// Configuration for environment : test
	Test Config

	// Configuration for environment : prod
	Prod Config
}{}

// setting config defaults for test, because when running tests
// config.json won't get loaded correctly unless specified by flags
// that gets painful when running individual tests
var config = &Config{
	STAGE:                "test",
	HOST:                 "localhost",
	PORT:                 "8000",
	ServerPort:           ":8000",
	DB_HOST:              "localhost",
	DB_PORT:              9042,
	DB_KEYSPACE:          "k8s_playground",
	DB_USERNAME:          "root",
	DB_PASSWORD:          "",
	KAFKA_SERVICE_NAME:   "",
	KAFKA_CONSUMER_GROUP: "",
	KAFKA_TOPIC:          "",
	KAFKA_USERNAME:       "",
	KAFKA_PASSWORD:       "",
	LogDir:               "storage/logs/",
	LogFileName:          "stdout",
	LogMaxSize:           50,
	LogLevel:             "debug",
}

var configFileName *string

// init reads the config.json file and loads the
// config options into config
func init() {
	configFileName = flag.String("config", "config.json", "Name of the config file")
	flag.Parse()

	stage, exists := os.LookupEnv("ENV")

	if !exists {
		if flag.Lookup("test.v") != nil {
			stage = "Test"
		} else {
			os.Stderr.WriteString("Set environment variable ENV to one of : Dev, Prod, Test. Taking Dev as default.")
			stage = "Dev"
		}
	}

	configFile, err := os.Open(*configFileName)
	if err != nil {
		if stage == "Test" {
			return // config is already set to default value for test. nothing to do.
		}
		log.Fatalf("Failed to open %s. Cannot proceed", *configFileName)
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&allConfigurations)

	if err != nil {
		log.Fatalf("Failed to load configuration. Cannot proceed. Error: %+v", err)
	}

	switch stage {
	case "Dev":
		config = &allConfigurations.Dev
	case "Test":
		config = &allConfigurations.Test
	case "Prod":
		config = &allConfigurations.Prod
	default:
		// Take Dev as default
		config = &allConfigurations.Dev
	}

	log.Printf("Loaded configuration from %s: %+v\n", *configFileName, config)
}

// GetConfiguration returns the configuration loaded from config.json
func GetConfiguration() *Config {
	return config
}

// Init intializes the utils package. The config is accepted as a parameter for helping with testing.
func Init(config *Config) {
	initLogger(config)
}
