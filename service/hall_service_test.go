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

type hallServiceTestSuite struct {
	suite.Suite
	mockHallRepository *repository.MockHallRepository
	service            HallService
}

func (suite *hallServiceTestSuite) SetupTest() {
	config.LoadTestConfig()
	appcontext.Init()
	mockController := gomock.NewController(suite.T())
	suite.mockHallRepository = repository.NewMockHallRepository(mockController)
	suite.service = &hallService{
		hallRepository: suite.mockHallRepository,
	}
}

func (suite *hallServiceTestSuite) TestAddHallReturnsErrorWhenHallNameIsMissing() {
	hall := []entities.Hall{
		{
			Name:  "",
			Seats: 200,
			TheatreId: sql.NullInt64{
				Valid: true,
				Int64: 1,
			},
		},
	}
	hallReturned, err := suite.service.Add(hall[0])
	assert.Equal(suite.T(), entities.Hall{}, hallReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), errors.Wrap(errors.New(constants.REQUEST_INVALID), constants.HALL_NAME_MANDATORY).Error(), err.Error())
}

func (suite *hallServiceTestSuite) TestAddHallReturnsErrorWhenNumberOfSeatsAreMissing() {
	hall := []entities.Hall{
		{
			Name: "Hall_A",
			TheatreId: sql.NullInt64{
				Valid: true,
				Int64: 1,
			},
		},
	}
	hallReturned, err := suite.service.Add(hall[0])
	assert.Equal(suite.T(), entities.Hall{}, hallReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), errors.Wrap(errors.New(constants.REQUEST_INVALID), constants.SEATS_NUMBER_INVALID).Error(), err.Error())
}

func (suite *hallServiceTestSuite) TestAddHallReturnsErrorWhenTheatreIdMissing() {
	hall := []entities.Hall{
		{
			Name:  "Hall_A",
			Seats: 200,
			TheatreId: sql.NullInt64{
				Valid: false,
				Int64: 0,
			},
		},
	}
	hallReturned, err := suite.service.Add(hall[0])
	assert.Equal(suite.T(), entities.Hall{}, hallReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), errors.Wrap(errors.New(constants.REQUEST_INVALID), constants.THEATRE_ID_HALL_CREATION_INVALID).Error(), err.Error())
}

func (suite *hallServiceTestSuite) TestAddHallReturnsErrorWhenTheatreIdViolatesForeignKeyConstraint() {
	hall := []entities.Hall{
		{
			Name:  "Hall_A",
			Seats: 200,
			TheatreId: sql.NullInt64{
				Valid: true,
				Int64: 1,
			},
		},
	}
	suite.mockHallRepository.EXPECT().FetchHallByNameAndTheatreId(gomock.Any()).Return(nil, nil)
	suite.mockHallRepository.EXPECT().InsertHall(gomock.Any()).Return(errors.New("Error 1452: Cannot add or update a child row: a foreign key constraint fails (`BOOK_MY_SHOW`.`halls`, CONSTRAINT `fk_theatre_id` FOREIGN KEY (`theatre_id`) REFERENCES `theatres` (`id`))"))
	hallReturned, err := suite.service.Add(hall[0])
	assert.Equal(suite.T(), entities.Hall{}, hallReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), errors.New("Error 1452: Cannot add or update a child row: a foreign key constraint fails (`BOOK_MY_SHOW`.`halls`, CONSTRAINT `fk_theatre_id` FOREIGN KEY (`theatre_id`) REFERENCES `theatres` (`id`)): failed to add new region").Error(), err.Error())
}

func (suite *hallServiceTestSuite) TestAddHallReturnsErrorWhenHallAlreadyExist() {
	hall := []entities.Hall{
		{
			Name:  "Hall_A",
			Seats: 200,
			TheatreId: sql.NullInt64{
				Valid: true,
				Int64: 1,
			},
		},
	}
	suite.mockHallRepository.EXPECT().FetchHallByNameAndTheatreId(gomock.Any()).Return(hall, nil)
	hallReturned, err := suite.service.Add(hall[0])
	assert.Equal(suite.T(), entities.Hall{}, hallReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), errors.Wrap(errors.New(constants.HALL_CREATION_FAILED), fmt.Sprintf(constants.HALL_ALREADY_EXISTS, hall[0].Name, hall[0].TheatreId.Int64)).Error(), err.Error())
}

func (suite *hallServiceTestSuite) TestHallIsSuccessfullyCreated() {
	hall := []entities.Hall{
		{
			Name:  "Hall_A",
			Seats: 200,
			TheatreId: sql.NullInt64{
				Valid: true,
				Int64: 1,
			},
		},
	}
	suite.mockHallRepository.EXPECT().FetchHallByNameAndTheatreId(gomock.Any()).Return(nil, nil)
	suite.mockHallRepository.EXPECT().InsertHall(gomock.Any()).Return(nil)
	hallReturned, err := suite.service.Add(hall[0])
	assert.Equal(suite.T(), hall[0], hallReturned)
	assert.Nil(suite.T(), err)
}

func TestHallServiceTestSuite(t *testing.T) {
	suite.Run(t, new(hallServiceTestSuite))
}
