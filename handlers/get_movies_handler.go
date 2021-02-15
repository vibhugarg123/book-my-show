package handlers

import (
	"github.com/vibhugarg123/book-my-show/constants"
	"github.com/vibhugarg123/book-my-show/domain"
	"github.com/vibhugarg123/book-my-show/service"
	"github.com/vibhugarg123/book-my-show/utils"
	"net/http"
)

type GetMoviesHandler struct {
	service service.MovieService
}

func NewGetMoviesHandler(movieService service.MovieService) *GetMoviesHandler {
	return &GetMoviesHandler{
		service: movieService,
	}
}

func (gmh *GetMoviesHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	movies, err := gmh.service.GetActiveMovies()
	if utils.IsValidationError(err) {
		utils.CommonResponse(writer, request, http.StatusBadRequest, domain.Error{constants.GET_MOVIES_CALL_FAILED, err.Error()})
		return
	} else if err != nil {
		utils.CommonResponse(writer, request, http.StatusInternalServerError, domain.Error{constants.GET_MOVIES_CALL_FAILED, err.Error()})
		return
	}
	utils.CommonResponse(writer, request, http.StatusOK, movies)
}
