package handlers

import (
	"bytes"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vibhugarg123/book-my-show/appcontext"
	"github.com/vibhugarg123/book-my-show/config"
	"github.com/vibhugarg123/book-my-show/constants"
	"github.com/vibhugarg123/book-my-show/entities"
	"github.com/vibhugarg123/book-my-show/service"
	"github.com/vibhugarg123/book-my-show/utils"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type getMoviesHandlerTestSuite struct {
	suite.Suite
	movieService    *service.MockMovieService
	getMovieHandler *GetMoviesHandler
}

func (suite *getMoviesHandlerTestSuite) SetupTest() {
	config.LoadTestConfig()
	appcontext.Init()
	mockController := gomock.NewController(suite.T())
	defer mockController.Finish()
	suite.movieService = service.NewMockMovieService(mockController)
	suite.getMovieHandler = NewGetMoviesHandler(suite.movieService)
}

func (suite *getMoviesHandlerTestSuite) TestGetMoviesHandlerReturnsActiveMoviesDoNotExist() {
	request, err := http.NewRequest("GET", "/movies", nil)
	assert.Nil(suite.T(), err)

	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Handle("/movies", suite.getMovieHandler)
	suite.movieService.EXPECT().GetActiveMovies().Return(nil, utils.WrapValidationError(errors.New(constants.NO_ACTIVE_MOVIES), constants.ACTIVE_MOVIES_NOT_PRESENT))
	router.ServeHTTP(response, request)
	assert.Equal(suite.T(), http.StatusBadRequest, response.Code)
	assert.Equal(suite.T(), []byte(`{"error_code":"get_movies_active_failed","error_message":"there are no active movies: no_active_movies"}`), bytes.TrimSpace(response.Body.Bytes()))
}

func (suite *getMoviesHandlerTestSuite) TestGetMoviesHandlerReturnsActiveMovies() {
	expectedMovies := []entities.Movie{
		{
			Name:         "ABCD",
			DirectorName: "SVG Production",
			ReleaseDate:  time.Date(2019, time.February, 1, 0, 0, 0, 0, time.UTC),
			IsActive:     true,
		},
	}
	request, err := http.NewRequest("GET", "/movies", nil)
	assert.Nil(suite.T(), err)

	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Handle("/movies", suite.getMovieHandler)
	suite.movieService.EXPECT().GetActiveMovies().Return(expectedMovies, nil)
	router.ServeHTTP(response, request)
	assert.Equal(suite.T(), http.StatusOK, response.Code)
	assert.Equal(suite.T(), []byte(`[{"id":0,"name":"ABCD","director_name":"SVG Production","release_date":"2019-02-01T00:00:00Z","is_active":true,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}]`), bytes.TrimSpace(response.Body.Bytes()))
}

func TestGetMoviesHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(getMoviesHandlerTestSuite))
}
