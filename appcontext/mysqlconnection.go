package appcontext

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/vibhugarg123/book-my-show/config"
	"github.com/vibhugarg123/book-my-show/constants"
)

var Db *sqlx.DB

func MySqlConnection() *sqlx.DB {
	return Db
}

func InitMySqlConnection() error {
	connection, err := sqlx.Connect(config.DatabaseDriverName(), config.DatabaseUserName()+":"+config.DatabasePassword()+"@"+config.DatabaseConnectionType()+"("+config.DatabaseHostIP()+")/"+config.DatabaseSchemaName()+"?parseTime=true")
	if err != nil {
		Logger.Error().
			Str("mysqlconnectionfailed", err.Error()).
			Msg(constants.FAILED_MYSQL_CONNECTION)
		return errors.Wrap(err, constants.FAILED_MYSQL_CONNECTION)
	}
	Db = connection
	Logger.Printf("connection to mysql database established successfully")
	return nil
}
