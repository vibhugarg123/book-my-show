package appcontext

import (
	"github.com/magiconair/properties/assert"
	"github.com/vibhugarg123/book-my-show/config"
	"testing"
)

func TestMySqlConnectionSuccessfullyEstablished(t *testing.T) {
	config.LoadTestConfig()
	SetupLogger()
	err := InitMySqlConnection()
	assert.Equal(t, nil, err)
}
