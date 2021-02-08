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

type getRegionHandlerTestSuite struct {
	suite.Suite
	regionService    *service.MockRegionService
	getRegionHandler *GetRegionHandler
}

func (suite *getRegionHandlerTestSuite) SetupTest() {
	config.LoadTestConfig()
	appcontext.Init()
	mockController := gomock.NewController(suite.T())
	defer mockController.Finish()
	suite.regionService = service.NewMockRegionService(mockController)
	suite.getRegionHandler = NewGetRegionHandler(suite.regionService)
}

func (suite *getRegionHandlerTestSuite) TestGetRegionHandlerReturnsBadRequestWhenFailedToDecodeRequestBody() {
	request, err := http.NewRequest("GET", "/region/1", nil)
	assert.Nil(suite.T(), err)

	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Handle("/region/{region-id}", suite.getRegionHandler)
	router.ServeHTTP(response, request)

	assert.Equal(suite.T(), http.StatusBadRequest, response.Code)
	assert.Equal(suite.T(), []byte(`{"error_code":"decoding_request_failed","error_message":"failed to decode the request body"}`), bytes.TrimSpace(response.Body.Bytes()))
}

func (suite *getRegionHandlerTestSuite) TestGetRegionHandlerReturnsSuccessResponse() {
	regionId := 1
	request, err := http.NewRequest("GET", "/region/1", nil)
	assert.Nil(suite.T(), err)
	expectedRegion := entities.Region{
		Id:         1,
		Name:       "Delhi",
		RegionType: 1,
	}
	suite.regionService.EXPECT().GetRegionById(regionId).Return(expectedRegion, nil)

	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Handle("/region/{region-id}", suite.getRegionHandler)
	router.ServeHTTP(response, request)

	assert.Equal(suite.T(), http.StatusOK, response.Code)
	assert.Equal(suite.T(), []byte(`{"error_code":"invalid_integer","error_message":"strconv.Atoi: parsing \"\": invalid syntax"}`), bytes.TrimSpace(response.Body.Bytes()))
}

func (suite *getRegionHandlerTestSuite) TestGetRegionHandlerWhenRegionDoesNotExist() {
	regionId := 1
	request, err := http.NewRequest("GET", "/region/1", nil)
	assert.Nil(suite.T(), err)
	suite.regionService.EXPECT().GetRegionById(regionId).Return(entities.Region{}, errors.Wrap(errors.New(constants.REGION_CREATION_FAILED), fmt.Sprintf(constants.REGION_DOES_NOT_EXIST, regionId)))

	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Handle("/region/{region-id}", suite.getRegionHandler)
	router.ServeHTTP(response, request)

	assert.Equal(suite.T(), http.StatusInternalServerError, response.Code)
	assert.Equal(suite.T(), []byte(`{"error_code":"get_region_by_id_failed","error_message":"region with region id- 1 do not exist: region_creation_failed"}`), bytes.TrimSpace(response.Body.Bytes()))
}

func TestGetRegionHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(getRegionHandlerTestSuite))
}
