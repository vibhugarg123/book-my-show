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
	"net/http"
	"net/http/httptest"
	"testing"
)

type getHallsByTheatreIdHandlerTestSuite struct {
	suite.Suite
	hallService     *service.MockHallService
	getHallsHandler *GetHallHandler
}

func (suite *getHallsByTheatreIdHandlerTestSuite) SetupTest() {
	config.LoadTestConfig()
	appcontext.Init()
	mockController := gomock.NewController(suite.T())
	defer mockController.Finish()
	suite.hallService = service.NewMockHallService(mockController)
	suite.getHallsHandler = NewGetHallHandler(suite.hallService)
}

func (suite *getHallsByTheatreIdHandlerTestSuite) TestGetHallsByTheatreIdHandlerReturnsErrorWhenTheatreIdIsInvalidInRequest() {
	request, err := http.NewRequest("GET", "/hall/A", nil)
	assert.Nil(suite.T(), err)
	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Handle("/hall/{theatre-id}", suite.getHallsHandler)
	router.ServeHTTP(response, request)
	assert.Equal(suite.T(), http.StatusBadRequest, response.Code)
	assert.Equal(suite.T(), []byte(`{"error_code":"invalid_integer","error_message":"strconv.Atoi: parsing \"A\": invalid syntax"}`), bytes.TrimSpace(response.Body.Bytes()))
}

func (suite *getHallsByTheatreIdHandlerTestSuite) TestGetHallsByTheatreIdHandlerReturnsErrorWhenHallsWithGivenTheatreIdDoesNotExist() {
	theatreId := 8
	request, err := http.NewRequest("GET", "/hall/1", nil)
	assert.Nil(suite.T(), err)
	suite.hallService.EXPECT().GetHallByTheatreId(gomock.Any()).Return([]entities.Hall{}, errors.Wrap(errors.New(constants.HALLS_DO_NOT_EXIST), fmt.Sprintf(constants.HALLS_DO_NOT_EXIST_WITH_THEATRE_ID, theatreId)))
	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Handle("/hall/{theatre-id}", suite.getHallsHandler)
	router.ServeHTTP(response, request)
	assert.Equal(suite.T(), http.StatusInternalServerError, response.Code)
	assert.Equal(suite.T(), []byte(`{"error_code":"get_halls_by_theatre_id_failed","error_message":"halls with theatre id- 8 do not exist: halls_do_not_exist"}`), bytes.TrimSpace(response.Body.Bytes()))
}

func (suite *getHallsByTheatreIdHandlerTestSuite) TestGetHallsByTheatreIdHandlerReturnsHallsListSuccessfully() {
	expectedHallList := []entities.Hall{{
		Name:  "HALL_A",
		Seats: 200,
		TheatreId: sql.NullInt64{
			Valid: true,
			Int64: 1,
		}},
	}
	request, err := http.NewRequest("GET", "/hall/1", nil)
	assert.Nil(suite.T(), err)
	suite.hallService.EXPECT().GetHallByTheatreId(gomock.Any()).Return(expectedHallList, nil)
	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Handle("/hall/{theatre-id}", suite.getHallsHandler)
	router.ServeHTTP(response, request)
	assert.Equal(suite.T(), http.StatusOK, response.Code)
	assert.Equal(suite.T(), []byte(`[{"id":0,"name":"HALL_A","seats":200,"theatre_id":{"Int64":1,"Valid":true},"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}]`), bytes.TrimSpace(response.Body.Bytes()))
}

func TestGetHallsByTheatreIdHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(getHallsByTheatreIdHandlerTestSuite))
}