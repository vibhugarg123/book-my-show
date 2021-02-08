package service

import (
	"database/sql"
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
	"testing"
)

type theatreServiceTestSuite struct {
	suite.Suite
	mockRegionRepository  *repository.MockRegionRepository
	mockTheatreRepository *repository.MockTheatreRepository
	service               TheatreService
}

func (suite *theatreServiceTestSuite) SetupTest() {
	config.LoadTestConfig()
	appcontext.Init()
	mockController := gomock.NewController(suite.T())
	suite.mockRegionRepository = repository.NewMockRegionRepository(mockController)
	suite.mockTheatreRepository = repository.NewMockTheatreRepository(mockController)
	suite.service = &theatreService{
		regionRepository:  suite.mockRegionRepository,
		theatreRepository: suite.mockTheatreRepository,
	}
}

func (suite *theatreServiceTestSuite) TestAddTheatreReturnsErrorWhenTheatreNameIsMissing() {
	theatre := entities.Theatre{
		Name:    "",
		Address: "GT Road,Panipat",
	}
	theatreReturned, err := suite.service.Add(theatre)
	assert.Equal(suite.T(), entities.Theatre{}, theatreReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), errors.Wrap(errors.New(constants.REQUEST_INVALID), constants.THEATRE_NAME_MANDATORY).Error(), err.Error())
}

func (suite *theatreServiceTestSuite) TestAddTheatreReturnsErrorWhenTheatreAddressIsMissing() {
	theatre := entities.Theatre{
		Name:    "PVR Cinemas",
		Address: "",
	}
	theatreReturned, err := suite.service.Add(theatre)
	assert.Equal(suite.T(), entities.Theatre{}, theatreReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), errors.Wrap(errors.New(constants.REQUEST_INVALID), constants.THEATRE_ADDRESS_MANDATORY).Error(), err.Error())
}

func (suite *theatreServiceTestSuite) TestAddTheatreReturnsErrorWhenTheatreRegionIdIsMissing() {
	theatre := entities.Theatre{
		Name:    "PVR Cinemas",
		Address: "MG Road, Gurgaon",
	}
	theatreReturned, err := suite.service.Add(theatre)
	assert.Equal(suite.T(), entities.Theatre{}, theatreReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), errors.Wrap(errors.New(constants.REQUEST_INVALID), constants.THEATRE_REGION_ID_MANDATORY).Error(), err.Error())
}

func (suite *theatreServiceTestSuite) TestAddTheatreReturnsErrorWhenTheatreBodyIsMissing() {
	theatre := entities.Theatre{}
	theatreReturned, err := suite.service.Add(theatre)
	assert.Equal(suite.T(), entities.Theatre{}, theatreReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), errors.Wrap(errors.New(constants.REQUEST_INVALID), constants.EMPTY_REQUEST_BODY).Error(), err.Error())
}

func (suite *theatreServiceTestSuite) TestAddTheatreReturnsErrorWhenTheatreRegionIdDoesNotExist() {
	theatre := entities.Theatre{
		Name:     "PVR Cinemas",
		Address:  "MG Road, Gurgaon",
		RegionId: sql.NullInt64{Int64: 6, Valid: true},
	}
	suite.mockRegionRepository.EXPECT().FetchRegionById(theatre.RegionId.Int64).Return([]entities.Region{}, errors.Wrap(errors.New(constants.REGION_DO_NOT_EXIST), fmt.Sprintf(constants.REGION_DOES_NOT_EXIST, theatre.RegionId.Int64)))
	theatreReturned, err := suite.service.Add(theatre)
	assert.Equal(suite.T(), entities.Theatre{}, theatreReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "region with region id- 6 do not exist: region_do_not_exist: select_query_failed", err.Error())
}

func (suite *theatreServiceTestSuite) TestAddTheatreReturnsErrorWhenTheatreAlreadyExist() {
	theatre := []entities.Theatre{
		{
			Id:       1,
			Name:     "PVR Cinemas",
			Address:  "MG Road, Gurgaon",
			RegionId: sql.NullInt64{Int64: 1, Valid: true},
		},
	}
	expectedRegion := []entities.Region{
		{
			Id:         1,
			Name:       "Kurukshetra",
			RegionType: 1,
		},
	}
	suite.mockTheatreRepository.EXPECT().FetchTheatreByNameRegionIdAndAddress(theatre[0]).Return(theatre, nil)
	suite.mockRegionRepository.EXPECT().FetchRegionById(theatre[0].RegionId.Int64).Return(expectedRegion, nil)
	theatreReturned, err := suite.service.Add(theatre[0])
	assert.Equal(suite.T(), entities.Theatre{}, theatreReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "theatre with details- {1 PVR Cinemas MG Road, Gurgaon {1 true} 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC} already exists: theatre_already_exist", err.Error())
}

func (suite *theatreServiceTestSuite) TestAddTheatreCreatedSuccessfully() {
	theatre := []entities.Theatre{
		{
			Name:     "PVR Cinemas",
			Address:  "MG Road, Gurgaon",
			RegionId: sql.NullInt64{Int64: 4, Valid: true},
		},
	}
	expectedRegion := []entities.Region{
		{
			Id:         1,
			Name:       "Kurukshetra",
			RegionType: 1,
		},
	}
	suite.mockRegionRepository.EXPECT().FetchRegionById(theatre[0].RegionId.Int64).Return(expectedRegion, nil)
	suite.mockTheatreRepository.EXPECT().FetchTheatreByNameRegionIdAndAddress(theatre[0]).Return(nil, nil)
	suite.mockTheatreRepository.EXPECT().InsertTheatre(theatre[0]).Return(nil)
	theatreReturned, err := suite.service.Add(theatre[0])
	assert.Equal(suite.T(), theatre[0], theatreReturned)
	assert.Nil(suite.T(), err)
}

func TestTheatreServiceTestSuite(t *testing.T) {
	suite.Run(t, new(theatreServiceTestSuite))
}
