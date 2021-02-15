package service

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vibhugarg123/book-my-show/appcontext"
	"github.com/vibhugarg123/book-my-show/config"
	"github.com/vibhugarg123/book-my-show/constants"
	"github.com/vibhugarg123/book-my-show/entities"
	"github.com/vibhugarg123/book-my-show/repository"
	"github.com/vibhugarg123/book-my-show/utils"
	"testing"
)

type userServiceTestSuite struct {
	suite.Suite
	mockUserRepository *repository.MockUserRepository
	service            UserService
}

func (suite *userServiceTestSuite) SetupTest() {
	config.LoadTestConfig()
	appcontext.Init()
	mockController := gomock.NewController(suite.T())
	suite.mockUserRepository = repository.NewMockUserRepository(mockController)
	suite.service = &userService{
		userRepository: suite.mockUserRepository,
	}
}

func (suite *userServiceTestSuite) TestCreateUserReturnsErrorWhenEmailIdAlreadyExists() {
	user := []entities.User{
		{
			FirstName: "power",
			LastName:  "puff",
			EmailId:   "xyz@gmail.com",
			Password:  "one@123",
		},
	}
	suite.mockUserRepository.EXPECT().FetchUserByEmailId(user[0].EmailId).Return(user, nil)
	userExists, err := suite.service.Add(user[0])
	assert.Equal(suite.T(), entities.User{}, userExists)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), utils.WrapValidationError(errors.New(constants.FAILED_CREATING_USER), fmt.Sprintf(constants.USER_ALREADY_EXISTS, user[0].EmailId)).Error(), err.Error())
}

func (suite *userServiceTestSuite) TestCreateUserReturnsErrorWhenEmailIdIsMissing() {
	user := []entities.User{
		{
			FirstName: "power",
			LastName:  "puff",
			EmailId:   "",
			Password:  "one@123",
		},
	}
	userExists, err := suite.service.Add(user[0])
	assert.Equal(suite.T(), entities.User{}, userExists)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.EMAIL_ID_MANDATORY).Error(), err.Error())
}

func (suite *userServiceTestSuite) TestCreateUserReturnsErrorWhenPasswordIsMissing() {
	user := []entities.User{
		{
			FirstName: "power",
			LastName:  "puff",
			EmailId:   "one@gmail.com",
			Password:  "",
		},
	}
	userExists, err := suite.service.Add(user[0])
	assert.Equal(suite.T(), entities.User{}, userExists)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.PASSWORD_MANDATORY).Error(), err.Error())
}

func (suite *userServiceTestSuite) TestCreateUserReturnsErrorWhenFirstNameIsMissing() {
	user := []entities.User{
		{
			FirstName: "",
			LastName:  "puff",
			EmailId:   "one@gmail.com",
			Password:  "one@313@33",
		},
	}
	userExists, err := suite.service.Add(user[0])
	assert.Equal(suite.T(), entities.User{}, userExists)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.FIRST_NAME_MANDATORY).Error(), err.Error())
}

func (suite *userServiceTestSuite) TestUserIsSuccessfullyCreated() {
	user := []entities.User{
		{
			FirstName: "power",
			LastName:  "puff",
			EmailId:   "one@gmail.com",
			Password:  "one@313@33",
		},
	}
	suite.mockUserRepository.EXPECT().FetchUserByEmailId(user[0].EmailId).Return(nil, nil)
	suite.mockUserRepository.EXPECT().InsertUser(user[0]).Return(nil)
	userExists, err := suite.service.Add(user[0])
	assert.Equal(suite.T(), user[0], userExists)
	assert.Nil(suite.T(), err)
}

func (suite *userServiceTestSuite) TestLoginWhenUserDoesNotExist() {
	user := []entities.User{
		{
			FirstName: "power",
			LastName:  "puff",
			EmailId:   "one@gmail.com",
			Password:  "one@313@33",
		},
	}
	suite.mockUserRepository.EXPECT().FetchUserByEmailIdAndPassword(user[0].EmailId, user[0].Password).Return(nil, nil)
	userExists, err := suite.service.Login(user[0])
	assert.Equal(suite.T(), entities.User{}, userExists)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), utils.WrapValidationError(errors.New(constants.USER_DO_NOT_EXIST), constants.USER_DOES_NOT_EXIST).Error(), err.Error())
}

func (suite *userServiceTestSuite) TestLoginIsSuccessful() {
	user := []entities.User{
		{
			FirstName: "power",
			LastName:  "puff",
			EmailId:   "one@gmail.com",
			Password:  "one@313@33",
		},
	}
	suite.mockUserRepository.EXPECT().FetchUserByEmailIdAndPassword(user[0].EmailId, user[0].Password).Return(user, nil)
	userExists, err := suite.service.Login(user[0])
	assert.Equal(suite.T(), user[0], userExists)
	assert.Nil(suite.T(), err)
}

func TestClientServiceTestSuite(t *testing.T) {
	suite.Run(t, new(userServiceTestSuite))
}
