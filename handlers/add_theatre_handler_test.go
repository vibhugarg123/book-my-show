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

type addTheatreHandlerTestSuite struct {
	suite.Suite
	theatreService    *service.MockTheatreService
	addTheatreHandler *AddTheatreHandler
}

func (suite *addTheatreHandlerTestSuite) SetupTest() {
	config.LoadTestConfig()
	appcontext.Init()
	mockController := gomock.NewController(suite.T())
	defer mockController.Finish()
	suite.theatreService = service.NewMockTheatreService(mockController)
	suite.addTheatreHandler = NewAddTheatreHandler(suite.theatreService)
}

func (suite *addTheatreHandlerTestSuite) TestAddTheatreHandlerReturnsBadRequestWhenFailedToDecodeRequestBody() {
	requestBody := ``
	request, err := http.NewRequest("POST", "/theatre", bytes.NewBuffer([]byte(requestBody)))
	assert.Nil(suite.T(), err)

	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Handle("/theatre", suite.addTheatreHandler)
	router.ServeHTTP(response, request)

	assert.Equal(suite.T(), http.StatusBadRequest, response.Code)
	assert.Equal(suite.T(), []byte(`{"error_code":"decoding_request_failed","error_message":"failed to decode the request body"}`), bytes.TrimSpace(response.Body.Bytes()))
}

func (suite *addTheatreHandlerTestSuite) TestAddRegionHandlerReturnsSuccessResponse() {
	requestBody := `{"name":"PVR Cinemas","address":"GT Road, Panipat","region_id":{"valid":true,"int64":6}}`
	request, err := http.NewRequest("POST", "/theatre", bytes.NewBuffer([]byte(requestBody)))
	assert.Nil(suite.T(), err)
	expectedTheatre := entities.Theatre{
		Name:     "PVR Cinemas",
		Address:  "GT Road, Panipat",
		RegionId: sql.NullInt64{Valid: true, Int64: 6},
	}
	suite.theatreService.EXPECT().Add(gomock.Any()).Return(expectedTheatre, nil)

	response := httptest.NewRecorder()
	handler := http.Handler(suite.addTheatreHandler)
	handler.ServeHTTP(response, request)

	assert.Equal(suite.T(), http.StatusOK, response.Code)
	assert.Equal(suite.T(), []byte(`{"id":0,"name":"PVR Cinemas","address":"GT Road, Panipat","region_id":{"Int64":6,"Valid":true},"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`), bytes.TrimSpace(response.Body.Bytes()))
}

func (suite *addTheatreHandlerTestSuite) TestAddTheatreHandlerReturnsTheatreAlreadyExists() {
	requestBody := `{"name":"PVR Cinemas","address":"GT Road, Panipat","region_id":{"valid":true,"int64":6}}`
	request, err := http.NewRequest("POST", "/theatre", bytes.NewBuffer([]byte(requestBody)))
	assert.Nil(suite.T(), err)
	expectedTheatre := entities.Theatre{
		Name:     "PVR Cinemas",
		Address:  "GT Road, Panipat",
		RegionId: sql.NullInt64{Valid: true, Int64: 6},
	}
	suite.theatreService.EXPECT().Add(gomock.Any()).Return(entities.Theatre{}, utils.WrapValidationError(errors.New(constants.THEATRE_ALREADY_EXIST), fmt.Sprintf(constants.THEATRE_ALREADY_EXISTS, expectedTheatre)))

	response := httptest.NewRecorder()
	handler := http.Handler(suite.addTheatreHandler)
	handler.ServeHTTP(response, request)

	assert.Equal(suite.T(), http.StatusBadRequest, response.Code)
	assert.Equal(suite.T(), []byte(`{"error_code":"theatre_creation_failed","error_message":"theatre with details- {0 PVR Cinemas GT Road, Panipat {6 true} 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC} already exists: theatre_already_exist"}`), bytes.TrimSpace(response.Body.Bytes()))
}

func TestAddTheatreHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(addTheatreHandlerTestSuite))
}
