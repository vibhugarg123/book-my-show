package server

import (
	"fmt"
	"github.com/urfave/negroni"
	"github.com/vibhugarg123/book-my-show/appcontext"
	"github.com/vibhugarg123/book-my-show/config"
	"gopkg.in/tylerb/graceful.v1"
)

func Start() error {
	router := Router()
	server := negroni.New()
	server.Use(negroni.NewRecovery())
	server.Use(negroni.NewLogger())
	server.UseHandler(router)
	appcontext.Logger.Info().Msg("starting book my show application")
	return graceful.RunWithErr(fmt.Sprintf(":%d", config.AppPort()),
		config.GracefulShutdownDuration(), server)
}
