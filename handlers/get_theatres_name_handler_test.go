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

type getTheatreByNameHandlerTestSuite struct {
	suite.Suite
	theatreService          *service.MockTheatreService
	getTheatreByNameHandler *GetTheatreByNameHandler
}

func (suite *getTheatreByNameHandlerTestSuite) SetupTest() {
	config.LoadTestConfig()
	appcontext.Init()
	mockController := gomock.NewController(suite.T())
	defer mockController.Finish()
	suite.theatreService = service.NewMockTheatreService(mockController)
	suite.getTheatreByNameHandler = NewGetTheatreByNameHandler(suite.theatreService)
}

func (suite *getTheatreByNameHandlerTestSuite) TestGetTheatreByNameHandlerReturnsErrorWhenTheatreNameIsMissingInRequest() {
	request, err := http.NewRequest("GET", "/theatre/ ", nil)
	assert.Nil(suite.T(), err)
	suite.theatreService.EXPECT().GetTheatreByName(gomock.Any()).Return([]entities.Theatre{}, utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.THEATRE_NAME_MANDATORY))
	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Handle("/theatre/{theatre-name}", suite.getTheatreByNameHandler)
	router.ServeHTTP(response, request)
	assert.Equal(suite.T(), http.StatusBadRequest, response.Code)
	assert.Equal(suite.T(), []byte(`{"error_code":"get_theatres_call_failed","error_message":"theatre name is missing in request: request_invalid"}`), bytes.TrimSpace(response.Body.Bytes()))
}

func (suite *getTheatreByNameHandlerTestSuite) TestGetTheatreByNameHandlerReturnsErrorWhenTheatreNameDoesNotExist() {
	theatreName := "PVR_CINEMAS"
	request, err := http.NewRequest("GET", "/theatre/PVR_CINEMAS", nil)
	assert.Nil(suite.T(), err)
	suite.theatreService.EXPECT().GetTheatreByName(gomock.Any()).Return([]entities.Theatre{}, utils.WrapValidationError(errors.New(constants.THEATRE_DO_NOT_EXIST), fmt.Sprintf(constants.THEATRE_DOES_NOT_EXIST, theatreName)))
	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Handle("/theatre/{theatre-name}", suite.getTheatreByNameHandler)
	router.ServeHTTP(response, request)
	assert.Equal(suite.T(), http.StatusBadRequest, response.Code)
	assert.Equal(suite.T(), []byte(`{"error_code":"get_theatres_call_failed","error_message":"theatre- PVR_CINEMAS do not exist: theatre_do_not_exist"}`), bytes.TrimSpace(response.Body.Bytes()))
}

func (suite *getTheatreByNameHandlerTestSuite) TestGetTheatreByNameHandlerReturnsTheatresListSuccessfully() {
	expectedTheatreList := []entities.Theatre{
		{
			Id:       1,
			Name:     "PVR Cinemas",
			Address:  "MG Road, Gurgaon",
			RegionId: sql.NullInt64{Int64: 4, Valid: true},
		},
	}
	request, err := http.NewRequest("GET", "/theatre/PVR_CINEMAS", nil)
	assert.Nil(suite.T(), err)
	suite.theatreService.EXPECT().GetTheatreByName(gomock.Any()).Return(expectedTheatreList, nil)
	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Handle("/theatre/{theatre-name}", suite.getTheatreByNameHandler)
	router.ServeHTTP(response, request)
	assert.Equal(suite.T(), http.StatusOK, response.Code)
	assert.Equal(suite.T(), []byte(`[{"id":1,"name":"PVR Cinemas","address":"MG Road, Gurgaon","region_id":{"Int64":4,"Valid":true},"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}]`), bytes.TrimSpace(response.Body.Bytes()))
}

func TestGetTheatreByNameHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(getTheatreByNameHandlerTestSuite))
}
