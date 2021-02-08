package server

import (
	"github.com/gorilla/mux"
	"github.com/vibhugarg123/book-my-show/appcontext"
	"github.com/vibhugarg123/book-my-show/handlers"
	"github.com/vibhugarg123/book-my-show/service"
	"net/http"
)

func Router() http.Handler {
	userService := service.NewUserService()
	regionService := service.NewRegionService()
	theatreService := service.NewTheatreService()
	appcontext.Logger.Info().Msg("setting up routes for book my show application")
	router := mux.NewRouter()
	router.Handle("/user", handlers.NewAddUserHandler(userService)).Methods("POST")
	router.Handle("/login", handlers.NewLoginHandler(userService)).Methods("POST")
	router.Handle("/region", handlers.NewAddRegionHandler(regionService)).Methods("POST")
	router.Handle("/region/{region-id}", handlers.NewGetRegionHandler(regionService)).Methods("GET")
	router.Handle("/theatre", handlers.NewAddTheatreHandler(theatreService)).Methods("POST")
	router.Handle("/theatre/{theatre-name}", handlers.NewGetTheatreByNameHandler(theatreService)).Methods("GET")
	return router
}
