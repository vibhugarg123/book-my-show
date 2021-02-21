package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/urfave/cli/v2"
	"github.com/vibhugarg123/book-my-show/appcontext"
	"github.com/vibhugarg123/book-my-show/config"
	"github.com/vibhugarg123/book-my-show/constants"
	"github.com/vibhugarg123/book-my-show/db"
	"github.com/vibhugarg123/book-my-show/server"
	"os"
	"runtime"
	"sort"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	config.Load()
	appcontext.Init()
	seedData()

	app := &cli.App{
		Name: "Book My Show- An online movie Ticketing System",
		Commands: []*cli.Command{
			{
				Name:    "server",
				Aliases: []string{"start"},
				Usage:   "starts the http server of book my show",
				Action: func(c *cli.Context) error {
					return server.Start()
				},
			},
			{
				Name:        "migrate:run",
				Description: "Running Migration",
				Action: func(c *cli.Context) error {
					appcontext.InitMySqlConnection()
					return db.RunDatabaseMigrations()
				},
			},
		},
	}

	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		appcontext.Logger.Fatal().
			Err(err).
			Msg(constants.FAILED_STARTING_APPLICATION)
	}
}

func seedData() {
}
