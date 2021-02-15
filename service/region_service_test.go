package service

import (
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

type regionServiceTestSuite struct {
	suite.Suite
	mockRegionRepository *repository.MockRegionRepository
	service              RegionService
}

func (suite *regionServiceTestSuite) SetupTest() {
	config.LoadTestConfig()
	appcontext.Init()
	mockController := gomock.NewController(suite.T())
	suite.mockRegionRepository = repository.NewMockRegionRepository(mockController)
	suite.service = &regionService{
		regionRepository: suite.mockRegionRepository,
	}
}

func (suite *regionServiceTestSuite) TestAddRegionReturnsErrorWhenRegionIdIsMissing() {
	region := []entities.Region{
		{
			Name:       "puff",
			RegionType: 1,
		},
	}
	regionReturned, err := suite.service.Add(region[0])
	assert.Equal(suite.T(), entities.Region{}, regionReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.REGION_ID_MANDATORY).Error(), err.Error())
}

func (suite *regionServiceTestSuite) TestAddRegionReturnsErrorWhenRegionNameIsMissing() {
	region := []entities.Region{
		{
			Id:         1,
			Name:       "",
			RegionType: 1,
		},
	}
	regionReturned, err := suite.service.Add(region[0])
	assert.Equal(suite.T(), entities.Region{}, regionReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.REGION_NAME_MANDATORY).Error(), err.Error())
}

func (suite *regionServiceTestSuite) TestAddRegionReturnsErrorWhenRegionTypeIsMissing() {
	region := []entities.Region{
		{
			Id:   1,
			Name: "Bangalore",
		},
	}
	regionReturned, err := suite.service.Add(region[0])
	assert.Equal(suite.T(), entities.Region{}, regionReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.REGION_TYPE_MANDATORY).Error(), err.Error())
}

func (suite *regionServiceTestSuite) TestAddRegionReturnsErrorWhenRegionIdAlreadyExists() {
	region := []entities.Region{
		{
			Id:         1,
			Name:       "Bangalore",
			RegionType: 1,
		},
	}
	suite.mockRegionRepository.EXPECT().FetchRegionById(int64(region[0].Id)).Return(region, nil)
	regionReturned, err := suite.service.Add(region[0])
	assert.Equal(suite.T(), entities.Region{}, regionReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "region with region id- 1 already exists: region_creation_failed", err.Error())
}

func (suite *regionServiceTestSuite) TestAddRegionWhenRegionIsCreatedSuccessfully() {
	region := []entities.Region{
		{
			Id:         1,
			Name:       "Bangalore",
			RegionType: 1,
		},
	}
	suite.mockRegionRepository.EXPECT().FetchRegionById(int64(region[0].Id)).Return(nil, nil)
	suite.mockRegionRepository.EXPECT().InsertRegion(region[0]).Return(nil)
	regionReturned, err := suite.service.Add(region[0])
	assert.Equal(suite.T(), region[0], regionReturned)
	assert.Nil(suite.T(), err)
}

func TestRegionServiceTestSuite(t *testing.T) {
	suite.Run(t, new(regionServiceTestSuite))
}
