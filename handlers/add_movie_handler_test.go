package handlers

import (
	"bytes"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vibhugarg123/book-my-show/appcontext"
	"github.com/vibhugarg123/book-my-show/config"
	"github.com/vibhugarg123/book-my-show/entities"
	"github.com/vibhugarg123/book-my-show/service"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type addMovieHandlerTestSuite struct {
	suite.Suite
	movieService    *service.MockMovieService
	addMovieHandler *AddMovieHandler
}

func (suite *addMovieHandlerTestSuite) SetupTest() {
	config.LoadTestConfig()
	appcontext.Init()
	mockController := gomock.NewController(suite.T())
	defer mockController.Finish()
	suite.movieService = service.NewMockMovieService(mockController)
	suite.addMovieHandler = NewAddMovieHandler(suite.movieService)
}

func (suite *addMovieHandlerTestSuite) TestAddMovieHandlerReturnsBadRequestWhenFailedToDecodeRequestBody() {
	requestBody := ``
	request, err := http.NewRequest("POST", "/movie", bytes.NewBuffer([]byte(requestBody)))
	assert.Nil(suite.T(), err)

	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Handle("/movie", suite.addMovieHandler)
	router.ServeHTTP(response, request)

	assert.Equal(suite.T(), http.StatusBadRequest, response.Code)
	assert.Equal(suite.T(), []byte(`{"error_code":"decoding_request_failed","error_message":"failed to decode the request body"}`), bytes.TrimSpace(response.Body.Bytes()))
}

func (suite *addMovieHandlerTestSuite) TestAddMovieHandlerReturnsSuccessResponse() {
	requestBody := `{"name":"ABCD","director_name":"Salman Khan","release_date":"2021-02-14T00:00:00+05:30","is_active":true}`
	request, err := http.NewRequest("POST", "/movie", bytes.NewBuffer([]byte(requestBody)))
	assert.Nil(suite.T(), err)
	expectedMovie := []entities.Movie{
		{
			DirectorName: "Salman Khan",
			ReleaseDate:  time.Date(2021, time.February, 14, 0, 0, 0, 0, time.UTC),
			IsActive:     true,
		},
	}
	suite.movieService.EXPECT().Add(gomock.Any()).Return(expectedMovie[0], nil)
	response := httptest.NewRecorder()
	handler := http.Handler(suite.addMovieHandler)
	handler.ServeHTTP(response, request)

	assert.Equal(suite.T(), http.StatusOK, response.Code)
	assert.Equal(suite.T(), []byte(`{"id":0,"name":"","director_name":"Salman Khan","release_date":"2021-02-14T00:00:00Z","is_active":true,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`), bytes.TrimSpace(response.Body.Bytes()))
}

func (suite *addMovieHandlerTestSuite) TestAddMovieHandlerReturnsErrorWhenReleaseDateIsIncorrect() {
	requestBody := `{"name":"ABCD","director_name":"Salman Khan","release_date":"2021-02-29T00:00:00+05:30","is_active":true}`
	request, err := http.NewRequest("POST", "/movie", bytes.NewBuffer([]byte(requestBody)))
	assert.Nil(suite.T(), err)
	response := httptest.NewRecorder()
	handler := http.Handler(suite.addMovieHandler)
	handler.ServeHTTP(response, request)

	assert.Equal(suite.T(), http.StatusBadRequest, response.Code)
	assert.Equal(suite.T(), []byte(`{"error_code":"decoding_request_failed","error_message":"failed to decode the request body"}`), bytes.TrimSpace(response.Body.Bytes()))
}

func TestAddMovieHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(addMovieHandlerTestSuite))
}
