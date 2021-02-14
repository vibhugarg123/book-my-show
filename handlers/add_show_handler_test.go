package handlers

import (
	"bytes"
	"database/sql"
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
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type addShowHandlerTestSuite struct {
	suite.Suite
	showService    *service.MockShowService
	addShowHandler *AddShowHandler
}

func (suite *addShowHandlerTestSuite) SetupTest() {
	config.LoadTestConfig()
	appcontext.Init()
	mockController := gomock.NewController(suite.T())
	defer mockController.Finish()
	suite.showService = service.NewMockShowService(mockController)
	suite.addShowHandler = NewAddShowHandler(suite.showService)
}

func (suite *addShowHandlerTestSuite) TestAddMovieHandlerReturnsBadRequestWhenFailedToDecodeRequestBody() {
	requestBody := ``
	request, err := http.NewRequest("POST", "/show", bytes.NewBuffer([]byte(requestBody)))
	assert.Nil(suite.T(), err)

	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Handle("/show", suite.addShowHandler)
	router.ServeHTTP(response, request)

	assert.Equal(suite.T(), http.StatusBadRequest, response.Code)
	assert.Equal(suite.T(), []byte(`{"error_code":"decoding_request_failed","error_message":"failed to decode the request body"}`), bytes.TrimSpace(response.Body.Bytes()))
}

func (suite *addShowHandlerTestSuite) TestAddShowHandlerReturnsSuccessResponse() {
	requestBody := `{"movie_id":{"valid":true,"int64":9},"hall_id":{"valid":true,"int64":1},"show_date":"2021-02-14T09:30:00+05:30","timing_id":{"name":"Morning","start_time":"2021-02-14T09:30:00+05:30","end_time":"2021-02-14T13:30:00+05:30"},"seat_price":190.85,"available_seats":200}`
	request, err := http.NewRequest("POST", "/show", bytes.NewBuffer([]byte(requestBody)))
	assert.Nil(suite.T(), err)
	expectedShow := []entities.Show{
		{
			Id: 7,
			MovieId: sql.NullInt64{
				Valid: true,
				Int64: 9,
			},
			HallId: sql.NullInt64{
				Valid: true,
				Int64: 1,
			},
			TimingId: entities.Timing{
				Name:      "Morning",
				StartTime: time.Date(2021, time.February, 14, 9, 30, 0, 0, time.UTC),
				EndTime:   time.Date(2021, time.February, 14, 13, 30, 0, 0, time.UTC),
			},
			ShowDate:       time.Date(2021, time.February, 14, 9, 30, 0, 0, time.UTC),
			SeatPrice:      190.85,
			AvailableSeats: 200,
		},
	}
	suite.showService.EXPECT().Add(gomock.Any()).Return(expectedShow[0], nil)
	response := httptest.NewRecorder()
	handler := http.Handler(suite.addShowHandler)
	handler.ServeHTTP(response, request)

	assert.Equal(suite.T(), http.StatusOK, response.Code)
	assert.Equal(suite.T(), []byte(`{"id":7,"movie_id":{"Int64":9,"Valid":true},"hall_id":{"Int64":1,"Valid":true},"show_date":"2021-02-14T09:30:00Z","timing_id":{"id":0,"name":"Morning","start_time":"2021-02-14T09:30:00Z","end_time":"2021-02-14T13:30:00Z","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"},"seat_price":190.85,"available_seats":200,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`), bytes.TrimSpace(response.Body.Bytes()))
}

func (suite *addShowHandlerTestSuite) TestAddShowHandlerReturnsErrorWhenMovieIdIsMissing() {
	requestBody := `{"movie_id":{"valid":false,"int64":9},"hall_id":{"valid":true,"int64":1},"show_date":"2021-02-14T09:30:00+05:30","timing_id":{"name":"Morning","start_time":"2021-02-14T09:30:00+05:30","end_time":"2021-02-14T13:30:00+05:30"},"seat_price":190.85,"available_seats":200}`
	request, err := http.NewRequest("POST", "/show", bytes.NewBuffer([]byte(requestBody)))
	assert.Nil(suite.T(), err)
	suite.showService.EXPECT().Add(gomock.Any()).Return(entities.Show{}, errors.Wrap(errors.New(constants.REQUEST_INVALID), constants.MOVIE_ID_MANDATORY_SHOW_CREATION))
	response := httptest.NewRecorder()
	handler := http.Handler(suite.addShowHandler)
	handler.ServeHTTP(response, request)

	assert.Equal(suite.T(), http.StatusInternalServerError, response.Code)
	assert.Equal(suite.T(), []byte(`{"error_code":"show_creation_failed","error_message":"movie id is mandatory in it's show creation request: request_invalid"}`), bytes.TrimSpace(response.Body.Bytes()))
}

func TestAddShowHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(addShowHandlerTestSuite))
}
