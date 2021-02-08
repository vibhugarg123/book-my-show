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
	"net/http"
	"net/http/httptest"
	"testing"
)

type addRegionHandlerTestSuite struct {
	suite.Suite
	regionService    *service.MockRegionService
	addRegionHandler *AddRegionHandler
}

func (suite *addRegionHandlerTestSuite) SetupTest() {
	config.LoadTestConfig()
	appcontext.Init()
	mockController := gomock.NewController(suite.T())
	defer mockController.Finish()
	suite.regionService = service.NewMockRegionService(mockController)
	suite.addRegionHandler = NewAddRegionHandler(suite.regionService)
}

func (suite *addRegionHandlerTestSuite) TestAddRegionHandlerReturnsBadRequestWhenFailedToDecodeRequestBody() {
	requestBody := ``
	request, err := http.NewRequest("POST", "/region", bytes.NewBuffer([]byte(requestBody)))
	assert.Nil(suite.T(), err)

	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Handle("/region", suite.addRegionHandler)
	router.ServeHTTP(response, request)

	assert.Equal(suite.T(), http.StatusBadRequest, response.Code)
	assert.Equal(suite.T(), []byte(`{"error_code":"decoding_request_failed","error_message":"failed to decode the request body"}`), bytes.TrimSpace(response.Body.Bytes()))
}

func (suite *addRegionHandlerTestSuite) TestAddRegionHandlerReturnsSuccessResponse() {
	requestBody := `{"id":2,"name":"Shahabad","region_type":1}`
	request, err := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(requestBody)))
	assert.Nil(suite.T(), err)
	expectedRegion := entities.Region{
		Id:         2,
		Name:       "puff",
		RegionType: 1,
	}
	suite.regionService.EXPECT().Add(gomock.Any()).Return(expectedRegion, nil)

	response := httptest.NewRecorder()
	handler := http.Handler(suite.addRegionHandler)
	handler.ServeHTTP(response, request)

	assert.Equal(suite.T(), http.StatusOK, response.Code)
	assert.Equal(suite.T(), []byte(`{"id":2,"name":"puff","region_type":1,"parent_id":{"Int64":0,"Valid":false},"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`), bytes.TrimSpace(response.Body.Bytes()))
}

func (suite *addRegionHandlerTestSuite) TestAddRegionHandlerWhenRegionAlreadyExists() {
	requestBody := `{"id":1,"name":"Delhi","region_type":1}`
	request, err := http.NewRequest("POST", "/user", bytes.NewBuffer([]byte(requestBody)))
	assert.Nil(suite.T(), err)
	region := entities.Region{
		Id:         1,
		Name:       "Delhi",
		RegionType: 1,
	}
	suite.regionService.EXPECT().Add(region).Return(entities.Region{}, errors.Wrap(errors.New(constants.REGION_CREATION_FAILED), fmt.Sprintf(constants.REGION_ALREADY_EXISTS, region.Id)))

	response := httptest.NewRecorder()
	handler := http.Handler(suite.addRegionHandler)
	handler.ServeHTTP(response, request)

	assert.Equal(suite.T(), http.StatusInternalServerError, response.Code)
	assert.Equal(suite.T(), []byte(`{"error_code":"region_creation_failed","error_message":"region with region id- 1 already exists: region_creation_failed"}`), bytes.TrimSpace(response.Body.Bytes()))
}

func TestAddRegionHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(addRegionHandlerTestSuite))
}
