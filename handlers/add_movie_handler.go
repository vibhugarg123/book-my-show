package handlers

import (
	"encoding/json"
	"github.com/vibhugarg123/book-my-show/appcontext"
	"github.com/vibhugarg123/book-my-show/constants"
	"github.com/vibhugarg123/book-my-show/domain"
	"github.com/vibhugarg123/book-my-show/entities"
	"github.com/vibhugarg123/book-my-show/service"
	"github.com/vibhugarg123/book-my-show/utils"
	"net/http"
)

type AddMovieHandler struct {
	service service.MovieService
}

func NewAddMovieHandler(movieService service.MovieService) *AddMovieHandler {
	return &AddMovieHandler{
		service: movieService,
	}
}

func (amh *AddMovieHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var movie entities.Movie
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&movie)
	if err != nil {
		appcontext.Logger.Error().Err(err).Msg("failed to decode the request body to add new movie")
		utils.CommonResponse(writer, request, http.StatusBadRequest, domain.Error{constants.REQUEST_DECODING_FAILED, constants.DECODING_REQUEST_FAILED})
		return
	}
	appcontext.Logger.Info().Msgf("request received to add new movie %v", movie)
	movie, err = amh.service.Add(movie)
	if err != nil {
		utils.CommonResponse(writer, request, http.StatusInternalServerError, domain.Error{constants.MOVIE_CREATION_FAILED, err.Error()})
		return
	}
	utils.CommonResponse(writer, request, http.StatusOK, movie)
}
