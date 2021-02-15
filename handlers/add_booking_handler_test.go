package handlers

import (
	"bytes"
	"fmt"
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
)

type addBookingHandlerTestSuite struct {
	suite.Suite
	bookingService    *service.MockBookingService
	addBookingHandler *AddBookingHandler
}

func (suite *addBookingHandlerTestSuite) SetupTest() {
	config.LoadTestConfig()
	appcontext.Init()
	mockController := gomock.NewController(suite.T())
	defer mockController.Finish()
	suite.bookingService = service.NewMockBookingService(mockController)
	suite.addBookingHandler = NewAddBookingHandler(suite.bookingService)
}

func (suite *addBookingHandlerTestSuite) TestAddBookingHandlerReturnsBadRequestWhenFailedToDecodeRequestBody() {
	requestBody := ``
	request, err := http.NewRequest("POST", "/booking", bytes.NewBuffer([]byte(requestBody)))
	assert.Nil(suite.T(), err)

	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Handle("/booking", suite.addBookingHandler)
	router.ServeHTTP(response, request)

	assert.Equal(suite.T(), http.StatusBadRequest, response.Code)
	assert.Equal(suite.T(), []byte(`{"error_code":"decoding_request_failed","error_message":"failed to decode the request body"}`), bytes.TrimSpace(response.Body.Bytes()))
}

func (suite *addBookingHandlerTestSuite) TestAddBookingHandlerReturnsSuccessResponse() {
	requestBody := `{"user_id":10,"show_id":1,"seats":5}`
	request, err := http.NewRequest("POST", "/booking", bytes.NewBuffer([]byte(requestBody)))
	assert.Nil(suite.T(), err)
	booking := []entities.Booking{
		{
			ShowId:     1,
			UserId:     10,
			Seats:      5,
			MovieId:    5,
			TotalPrice: 954.25,
		},
	}
	suite.bookingService.EXPECT().Add(gomock.Any()).Return(booking[0], nil)

	response := httptest.NewRecorder()
	handler := http.Handler(suite.addBookingHandler)
	handler.ServeHTTP(response, request)

	assert.Equal(suite.T(), http.StatusOK, response.Code)
	assert.Equal(suite.T(), []byte(`{"id":0,"user_id":10,"show_id":1,"seats":5,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","total_price":954.25,"movie_id":5}`), bytes.TrimSpace(response.Body.Bytes()))
}

func (suite *addBookingHandlerTestSuite) TestAddBookingHandlerReturnsErrorResponseWhenUserIdViolatesForeignKeyConstraint() {
	requestBody := `{"user_id":10,"show_id":6,"seats":1}`
	request, err := http.NewRequest("POST", "/booking", bytes.NewBuffer([]byte(requestBody)))
	assert.Nil(suite.T(), err)
	booking := []entities.Booking{
		{
			ShowId:     6,
			UserId:     10,
			Seats:      1,
			MovieId:    5,
			TotalPrice: 105.50,
		},
	}
	suite.bookingService.EXPECT().Add(gomock.Any()).Return(entities.Booking{}, utils.WrapValidationError(errors.New(constants.FOREIGN_KEY_VIOLATION), fmt.Sprintf(constants.USER_ID_FOREIGN_KEY_VIOLATION_IN_CREATE_BOOKING, booking[0].UserId)))

	response := httptest.NewRecorder()
	handler := http.Handler(suite.addBookingHandler)
	handler.ServeHTTP(response, request)

	assert.Equal(suite.T(), http.StatusBadRequest, response.Code)
	assert.Equal(suite.T(), []byte(`{"error_code":"booking_creation_failed","error_message":"[user-id - 10] in create booking request does not exist: foreign_key_violation"}`), bytes.TrimSpace(response.Body.Bytes()))
}

func TestAddBookingHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(addBookingHandlerTestSuite))
}
