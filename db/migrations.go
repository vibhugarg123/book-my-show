package db

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/vibhugarg123/book-my-show/appcontext"
)

const (
	migrationFilePath = "db/migration"
)

var appMigrate *migrate.Migrate

func Run() error {
	err := appMigrate.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	fmt.Println("Migrations successful")
	return nil
}

func RunDatabaseMigrations() error {
	driver, err := mysql.WithInstance(appcontext.MySqlConnection().DB, &mysql.Config{})
	if err != nil {
		appcontext.Logger.Info().Msgf("could not start sql migration... %v", err)
		return err
	}
	var migrationDir = flag.String("migration.files", migrationFilePath, "Directory where the migration files are located ?")
	appcontext.Logger.Info().Msgf(fmt.Sprintf("file://%s", *migrationDir))
	appMigrate, err = migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", *migrationDir),
		"mysql", driver)

	if err != nil {
		appcontext.Logger.Info().Msgf("migration failed... %v", err)
		return err
	}
	return Run()
}
