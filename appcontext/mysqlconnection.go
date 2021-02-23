package appcontext

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
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
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&multiStatements=true", config.DatabaseUserName(), config.DatabasePassword(), config.DatabaseHostIP(), config.DatabasePort(), "BOOK_MY_SHOW")
	Logger.Info().Msgf("connection url %s", dbUrl)
	connection, err := sqlx.Connect(config.DatabaseDriverName(), dbUrl)
	if err != nil {
		Logger.Error().
			Str("mysqlconnectionfailed", err.Error()).
			Msg(constants.FAILED_MYSQL_CONNECTION)
		return errors.Wrap(err, constants.FAILED_MYSQL_CONNECTION)
	}
	Db = connection
	Logger.Info().Msg("connection to my sql established successfully")
	return nil
}
