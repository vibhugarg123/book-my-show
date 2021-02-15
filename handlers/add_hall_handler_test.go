package handlers

import (
	"bytes"
	"database/sql"
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

type addHallHandlerTestSuite struct {
	suite.Suite
	hallService    *service.MockHallService
	addHallHandler *AddHallHandler
}

func (suite *addHallHandlerTestSuite) SetupTest() {
	config.LoadTestConfig()
	appcontext.Init()
	mockController := gomock.NewController(suite.T())
	defer mockController.Finish()
	suite.hallService = service.NewMockHallService(mockController)
	suite.addHallHandler = NewAddHallHandler(suite.hallService)
}

func (suite *addHallHandlerTestSuite) TestAddHallHandlerReturnsBadRequestWhenFailedToDecodeRequestBody() {
	requestBody := ``
	request, err := http.NewRequest("POST", "/hall", bytes.NewBuffer([]byte(requestBody)))
	assert.Nil(suite.T(), err)

	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Handle("/hall", suite.addHallHandler)
	router.ServeHTTP(response, request)

	assert.Equal(suite.T(), http.StatusBadRequest, response.Code)
	assert.Equal(suite.T(), []byte(`{"error_code":"decoding_request_failed","error_message":"failed to decode the request body"}`), bytes.TrimSpace(response.Body.Bytes()))
}

func (suite *addHallHandlerTestSuite) TestAddHallHandlerReturnsSuccessResponse() {
	requestBody := `{"name":"HALL_A","seats":200,"theatre_id":{"valid":true,"int64":7}}`
	request, err := http.NewRequest("POST", "/hall", bytes.NewBuffer([]byte(requestBody)))
	assert.Nil(suite.T(), err)
	expectedHall := entities.Hall{
		Name:  "HALL_A",
		Seats: 200,
		TheatreId: sql.NullInt64{
			Valid: true,
			Int64: 7,
		},
	}
	suite.hallService.EXPECT().Add(gomock.Any()).Return(expectedHall, nil)

	response := httptest.NewRecorder()
	handler := http.Handler(suite.addHallHandler)
	handler.ServeHTTP(response, request)

	assert.Equal(suite.T(), http.StatusOK, response.Code)
	assert.Equal(suite.T(), []byte(`{"id":0,"name":"HALL_A","seats":200,"theatre_id":{"Int64":7,"Valid":true},"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`), bytes.TrimSpace(response.Body.Bytes()))
}

func (suite *addHallHandlerTestSuite) TestAddHallHandlerWhenHallAlreadyExists() {
	requestBody := `{"name":"HALL_A","seats":200,"theatre_id":{"valid":true,"int64":7}}`
	request, err := http.NewRequest("POST", "/hall", bytes.NewBuffer([]byte(requestBody)))
	assert.Nil(suite.T(), err)
	expectedHall := entities.Hall{
		Name:  "HALL_A",
		Seats: 200,
		TheatreId: sql.NullInt64{
			Valid: true,
			Int64: 7,
		},
	}
	suite.hallService.EXPECT().Add(gomock.Any()).Return(entities.Hall{}, utils.WrapValidationError(errors.New(constants.HALL_CREATION_FAILED), fmt.Sprintf(constants.HALL_ALREADY_EXISTS, expectedHall.Name, expectedHall.TheatreId.Int64)))

	response := httptest.NewRecorder()
	handler := http.Handler(suite.addHallHandler)
	handler.ServeHTTP(response, request)

	assert.Equal(suite.T(), http.StatusBadRequest, response.Code)
	assert.Equal(suite.T(), []byte(`{"error_code":"hall_creation_failed","error_message":"hall with hall name- HALL_A \u0026 theatre-id 7 already exists: hall_creation_failed"}`), bytes.TrimSpace(response.Body.Bytes()))
}

func TestAddHallHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(addHallHandlerTestSuite))
}
