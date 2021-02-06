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

type addUserHandlerTestSuite struct {
	suite.Suite
	userService    *service.MockUserService
	addUserHandler *AddUserHandler
}

func (suite *addUserHandlerTestSuite) SetupTest() {
	config.LoadTestConfig()
	appcontext.Init()
	mockController := gomock.NewController(suite.T())
	defer mockController.Finish()
	suite.userService = service.NewMockUserService(mockController)
	suite.addUserHandler = NewAddUserHandler(suite.userService)
}

func (suite *addUserHandlerTestSuite) TestAddUserHandlerReturnsBadRequestWhenFailedToDecodeRequestBody() {
	requestBody := ``
	request, err := http.NewRequest("POST", "/user", bytes.NewBuffer([]byte(requestBody)))
	assert.Nil(suite.T(), err)

	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Handle("/user", suite.addUserHandler)
	router.ServeHTTP(response, request)

	assert.Equal(suite.T(), []byte(`{"error_code":"decoding_request_failed","error_message":"failed to decode the request body"}`), bytes.TrimSpace(response.Body.Bytes()))
	assert.Equal(suite.T(), http.StatusBadRequest, response.Code)
}

func (suite *addUserHandlerTestSuite) TestAddUserHandlerReturnsSuccessResponse() {
	requestBody := `{"first_name":"cbbc","last_name":"cbbc","email_id":"cbbc@ghuil.com","password":"helo@123"}`
	request, err := http.NewRequest("POST", "/user", bytes.NewBuffer([]byte(requestBody)))
	assert.Nil(suite.T(), err)

	expectedUser := entities.User{
		FirstName: "cbbc",
		LastName:  "cbbc",
		EmailId:   "cbbc@ghuil.com",
		Password:  "helo@123",
	}
	suite.userService.EXPECT().Add(gomock.Any()).Return(expectedUser, nil)

	response := httptest.NewRecorder()
	handler := http.Handler(suite.addUserHandler)
	handler.ServeHTTP(response, request)

	assert.Equal(suite.T(), http.StatusOK, response.Code)
	assert.Equal(suite.T(), []byte(`{"id":0,"first_name":"cbbc","last_name":"cbbc","email_id":"cbbc@ghuil.com","password":"helo@123","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`), bytes.TrimSpace(response.Body.Bytes()))
}

func (suite *addUserHandlerTestSuite) TestAddUserHandlerReturnsFailureOnUserCreation() {
	requestBody := `{"first_name":"cbbc","last_name":"cbbc","email_id":"cbbc@ghuil.com","password":"helo@123"}`
	request, err := http.NewRequest("POST", "/user", bytes.NewBuffer([]byte(requestBody)))
	assert.Nil(suite.T(), err)
	user := entities.User{
		FirstName: "cbbc",
		LastName:  "cbbc",
		EmailId:   "cbbc@ghuil.com",
		Password:  "helo@123",
	}
	suite.userService.EXPECT().Add(gomock.Any()).Return(entities.User{}, errors.Wrap(errors.New(constants.FAILED_CREATING_USER), fmt.Sprintf(constants.USER_ALREADY_EXISTS, user.EmailId)))

	response := httptest.NewRecorder()
	handler := http.Handler(suite.addUserHandler)
	handler.ServeHTTP(response, request)

	assert.Equal(suite.T(), http.StatusInternalServerError, response.Code)
	assert.Equal(suite.T(), []byte(`{"error_code":"user_creation_failed","error_message":"user with email id- cbbc@ghuil.com already exists: user_creation_failed"}`), bytes.TrimSpace(response.Body.Bytes()))
}

func TestAddConfigurationMappingsHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(addUserHandlerTestSuite))
}
