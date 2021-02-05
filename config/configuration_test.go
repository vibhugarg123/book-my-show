package config

import (
	"github.com/magiconair/properties/assert"
	"testing"
	"time"
)

func TestApplicationConfigurations(t *testing.T) {
	LoadTestConfig()
	expectedAppPort := 8080
	expectedLogLevel := "debug"
	expectedAppName := "book_my_show"
	expectedGracefulShutdownDurationSecs := time.Duration(10 * time.Second)
	expectedDatabaseDriverName := "mysql"
	expectedDatabaseUserName := "root"
	expectedDatabasePassword := "ons_vg"
	expectedDatabaseConnectionType := "tcp"
	expectedDatabaseHostIP := "127.0.0.1"
	expectedDatabasePort := "3306"
	expectedDatabaseSchemaName := "BOOK_MY_SHOW"

	assert.Equal(t, expectedAppPort, AppPort())
	assert.Equal(t, expectedAppName, AppName())
	assert.Equal(t, expectedLogLevel, LogLevel())
	assert.Equal(t, expectedGracefulShutdownDurationSecs, GracefulShutdownDuration())
	assert.Equal(t, expectedDatabaseDriverName, DatabaseDriverName())
	assert.Equal(t, expectedDatabaseUserName, DatabaseUserName())
	assert.Equal(t, expectedDatabasePassword, DatabasePassword())
	assert.Equal(t, expectedDatabaseConnectionType, DatabaseConnectionType())
	assert.Equal(t, expectedDatabaseHostIP, DatabaseHostIP())
	assert.Equal(t, expectedDatabasePort, DatabasePort())
	assert.Equal(t, expectedDatabaseSchemaName, DatabaseSchemaName())
}
