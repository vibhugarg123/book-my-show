package config

import (
	"github.com/vibhugarg123/book-my-show/utils"
	"time"

	"github.com/spf13/viper"
)

type Configuration struct {
	appPort                      int
	logLevel                     string
	appName                      string
	gracefulShutdownDurationSecs int
	databaseDriverName           string
	databaseUserName             string
	databasePassword             string
	databaseConnectionType       string
	databaseHostIP               string
	databasePort                 string
	databaseSchemaName           string
}

var appConfiguration *Configuration

func LoadTestConfig() {
	load("application.test")
}

func Load() {
	load("application")
}

func load(configFileName string) {
	viper.AutomaticEnv()
	viper.SetConfigName(configFileName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	viper.ReadInConfig()

	appConfiguration = &Configuration{
		appPort:                      utils.GetIntOrPanic("app_port"),
		appName:                      utils.GetStringOrPanic("app_name"),
		logLevel:                     utils.GetStringOrPanic("log_level"),
		gracefulShutdownDurationSecs: utils.GetIntOrPanic("graceful_shutdown_duration_secs"),
		databaseDriverName:           utils.GetStringOrPanic("database_driver_name"),
		databaseUserName:             utils.GetStringOrPanic("database_user_name"),
		databasePassword:             utils.GetStringOrPanic("database_password"),
		databaseConnectionType:       utils.GetStringOrPanic("database_connection_type"),
		databaseHostIP:               utils.GetStringOrPanic("database_host_ip"),
		databasePort:                 utils.GetStringOrPanic("database_port"),
		databaseSchemaName:           utils.GetStringOrPanic("database_schema_name"),
	}
}

func AppPort() int {
	return appConfiguration.appPort
}

func AppName() string {
	return appConfiguration.appName
}

func LogLevel() string {
	return appConfiguration.logLevel
}

func GracefulShutdownDuration() time.Duration {
	return time.Duration(appConfiguration.gracefulShutdownDurationSecs) * time.Second
}

func DatabaseDriverName() string {
	return appConfiguration.databaseDriverName
}

func DatabaseUserName() string {
	return appConfiguration.databaseUserName
}

func DatabasePassword() string {
	return appConfiguration.databasePassword
}

func DatabaseConnectionType() string {
	return appConfiguration.databaseConnectionType
}

func DatabaseHostIP() string {
	return appConfiguration.databaseHostIP
}

func DatabasePort() string {
	return appConfiguration.databasePort
}

func DatabaseSchemaName() string {
	return appConfiguration.databaseSchemaName
}
