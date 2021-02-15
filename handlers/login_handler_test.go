package handlers

import (
	"bytes"
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

type loginHandlerTestSuite struct {
	suite.Suite
	userService  *service.MockUserService
	loginHandler *LoginHandler
}

func (suite *loginHandlerTestSuite) SetupTest() {
	config.LoadTestConfig()
	appcontext.Init()
	mockController := gomock.NewController(suite.T())
	defer mockController.Finish()
	suite.userService = service.NewMockUserService(mockController)
	suite.loginHandler = NewLoginHandler(suite.userService)
}

func (suite *loginHandlerTestSuite) TestLoginHandlerReturnsBadRequestWhenFailedToDecodeRequestBody() {
	requestBody := ``
	request, err := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(requestBody)))
	assert.Nil(suite.T(), err)

	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.Handle("/login", suite.loginHandler)
	router.ServeHTTP(response, request)

	assert.Equal(suite.T(), []byte(`{"error_code":"decoding_request_failed","error_message":"failed to decode the request body"}`), bytes.TrimSpace(response.Body.Bytes()))
	assert.Equal(suite.T(), http.StatusBadRequest, response.Code)
}

func (suite *loginHandlerTestSuite) TestLoginHandlerReturnsSuccessResponse() {
	requestBody := `{"email_id":"bhambri.lakshay@gmail.com","password":"helo@123"}`
	request, err := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(requestBody)))
	assert.Nil(suite.T(), err)

	expectedUser := entities.User{
		EmailId:  "bhambri.lakshay@gmail.com",
		Password: "helo@123",
	}
	suite.userService.EXPECT().Login(gomock.Any()).Return(expectedUser, nil)

	response := httptest.NewRecorder()
	handler := http.Handler(suite.loginHandler)
	handler.ServeHTTP(response, request)

	assert.Equal(suite.T(), http.StatusOK, response.Code)
	assert.Equal(suite.T(), []byte(`{"login_status":"login successful"}`), bytes.TrimSpace(response.Body.Bytes()))
}

func (suite *loginHandlerTestSuite) TestLoginUserDoesNotExist() {
	requestBody := `{"email_id":"cbbc@ghuil.com","password":"helo@123"}`
	request, err := http.NewRequest("POST", "/user", bytes.NewBuffer([]byte(requestBody)))
	assert.Nil(suite.T(), err)
	suite.userService.EXPECT().Login(gomock.Any()).Return(entities.User{}, utils.WrapValidationError(errors.New(constants.USER_DO_NOT_EXIST), constants.USER_DOES_NOT_EXIST))

	response := httptest.NewRecorder()
	handler := http.Handler(suite.loginHandler)
	handler.ServeHTTP(response, request)

	assert.Equal(suite.T(), http.StatusBadRequest, response.Code)
	assert.Equal(suite.T(), []byte(`{"error_code":"login_failed","error_message":"user does not exist: user_not_present"}`), bytes.TrimSpace(response.Body.Bytes()))
}

func TestLoginHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(loginHandlerTestSuite))
}
