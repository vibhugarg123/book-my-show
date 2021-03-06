package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/vibhugarg123/book-my-show/constants"
	"net/http"
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

func CommonResponse(writer http.ResponseWriter, request *http.Request, httpstatuscode int, result interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(httpstatuscode)
	json.NewEncoder(writer).Encode(result)
}

func ValidateIntegerType(value interface{}) error {
	_, ok := value.(int)
	if !ok {
		return errors.New(constants.NOT_VALID_INTEGER)
	}
	return nil
}

func SqlError(err error) error {
	switch err.(*mysql.MySQLError).Number {
	case 1452:
		return errors.New(constants.FOREIGN_KEY_VIOLATION)
	default:
		return err
	}
}
