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

type getBookingHandlerTestSuite struct {
	suite.Suite
	bookingService    *service.MockBookingService
	getBookingHandler *GetBookingHandler
}

func (suite *getBookingHandlerTestSuite) SetupTest() {
	config.LoadTestConfig()
	appcontext.Init()
	mockController := gomock.NewController(suite.T())
	defer mockController.Finish()
	suite.bookingService = service.NewMockBookingService(mockController)
	suite.getBookingHandler = NewGetBookingHandler(suite.bookingService)
}

func (suite *getBookingHandlerTestSuite) TestGetBookingsByUserIdHandlerReturnsBadRequestWhenFailedToDecodeUserId() {
	request, err := http.NewRequest("GET", "/booking/userid/A", nil)
	assert.Nil(suite.T(), err)
	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Handle("/booking/userid/{user-id}", suite.getBookingHandler)
	router.ServeHTTP(response, request)
	assert.Equal(suite.T(), http.StatusBadRequest, response.Code)
	assert.Equal(suite.T(), []byte(`{"error_code":"invalid_integer","error_message":"strconv.Atoi: parsing \"A\": invalid syntax"}`), bytes.TrimSpace(response.Body.Bytes()))
}

func (suite *getBookingHandlerTestSuite) TestGetBookingsByUserIdHandlerReturnBookingsSuccessfully() {
	request, err := http.NewRequest("GET", "/booking/userid/1", nil)
	assert.Nil(suite.T(), err)
	booking := []entities.Booking{
		{
			ShowId:     1,
			UserId:     1,
			Seats:      5,
			MovieId:    5,
			TotalPrice: 954.25,
		},
	}
	suite.bookingService.EXPECT().GetBooking(gomock.Any()).Return(booking, nil)
	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Handle("/booking/userid/{user-id}", suite.getBookingHandler)
	router.ServeHTTP(response, request)

	assert.Equal(suite.T(), http.StatusOK, response.Code)
	assert.Equal(suite.T(), []byte(`[{"id":0,"user_id":1,"show_id":1,"seats":5,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","total_price":954.25,"movie_id":5}]`), bytes.TrimSpace(response.Body.Bytes()))
}

func (suite *getBookingHandlerTestSuite) TestGetBookingsByUserIdHandlerReturnErrorWhenUserDoesNotExist() {
	request, err := http.NewRequest("GET", "/booking/userid/1", nil)
	assert.Nil(suite.T(), err)
	booking := []entities.Booking{
		{
			ShowId:     1,
			UserId:     1,
			Seats:      5,
			MovieId:    5,
			TotalPrice: 954.25,
		},
	}
	suite.bookingService.EXPECT().GetBooking(gomock.Any()).Return(nil, utils.WrapValidationError(errors.New(constants.BOOKING_DO_NOT_EXIST), fmt.Sprintf(constants.BOOKING_DO_NOT_EXIST_FOR_GIVEN_USER_ID, booking[0].UserId)))
	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Handle("/booking/userid/{user-id}", suite.getBookingHandler)
	router.ServeHTTP(response, request)

	assert.Equal(suite.T(), http.StatusBadRequest, response.Code)
	assert.Equal(suite.T(), []byte(`{"error_code":"get_bookings_by_user_id_failed","error_message":"booking do not exist for user-id 1: bookings_do_not_exist"}`), bytes.TrimSpace(response.Body.Bytes()))
}

func TestGetBookingHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(getBookingHandlerTestSuite))
}
