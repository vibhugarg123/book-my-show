package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

func GetIntOrPanic(key string) int {
	intValue, err := strconv.Atoi(GetConfigStringValue(key))
	if err != nil {
		PanicForkey(key, err)
	}
	return intValue
}

func GetBoolOrPanic(key string) bool {
	boolValue, err := strconv.ParseBool(GetConfigStringValue(key))
	if err != nil {
		PanicForkey(key, err)
	}
	return boolValue
}

func GetStringOrPanic(key string) string {
	value := GetConfigStringValue(key)
	if value == "" {
		PanicForkey(key, errors.New("config is not set"))
	}
	return value
}

func GetConfigStringValue(key string) string {
	if !viper.IsSet(key) && os.Getenv(key) == "" {
		fmt.Printf("config %s is not set\n", key)
	}
	value := os.Getenv(key)
	if value == "" {
		value = viper.GetString(key)
	}
	return value
}

func PanicForkey(key string, err error) {
	panic(fmt.Sprintf("error %v occured while reading config %s", err.Error(), key))
}
