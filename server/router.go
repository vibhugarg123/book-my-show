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
	appcontext.Logger.Info().Msg("setting up routes")
	router := mux.NewRouter()
	router.Handle("/user", handlers.NewAddUserHandler(userService)).Methods("POST")
	router.Handle("/login", handlers.NewLoginHandler(userService)).Methods("POST")
	return router
}
