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
	"github.com/vibhugarg123/book-my-show/utils"
	"testing"
	"time"
)

type showServiceTestSuite struct {
	suite.Suite
	mockShowRepository *repository.MockShowRepository
	mockHallRepository *repository.MockHallRepository
	service            ShowService
}

func (suite *showServiceTestSuite) SetupTest() {
	config.LoadTestConfig()
	appcontext.Init()
	mockController := gomock.NewController(suite.T())
	suite.mockShowRepository = repository.NewMockShowRepository(mockController)
	suite.mockHallRepository = repository.NewMockHallRepository(mockController)
	suite.service = &showService{
		showRepository: suite.mockShowRepository,
		hallRepository: suite.mockHallRepository,
	}
}

func (suite *showServiceTestSuite) TestAddShowReturnsErrorWhenShowObjectIsMissing() {
	show := []entities.Show{
		{
		},
	}
	showReturned, err := suite.service.Add(show[0])
	assert.Equal(suite.T(), entities.Show{}, showReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.EMPTY_REQUEST_BODY).Error(), err.Error())
}

func (suite *showServiceTestSuite) TestAddShowReturnsErrorWhenShowMovieIdIsNotValid() {
	show := []entities.Show{
		{
			MovieId: sql.NullInt64{
				Valid: false,
				Int64: 1,
			},
			HallId: sql.NullInt64{
				Valid: true,
				Int64: 1,
			},
			ShowDate: time.Date(2021, time.February, 14, 9, 30, 0, 0, time.UTC),
			TimingId: entities.Timing{
				Name:      "Morning",
				StartTime: time.Date(2021, time.February, 14, 9, 30, 0, 0, time.UTC),
				EndTime:   time.Date(2021, time.February, 14, 12, 30, 0, 0, time.UTC),
			},
			SeatPrice: 190.85,
		},
	}
	showReturned, err := suite.service.Add(show[0])
	assert.Equal(suite.T(), entities.Show{}, showReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.MOVIE_ID_MANDATORY_SHOW_CREATION).Error(), err.Error())
}

func (suite *showServiceTestSuite) TestAddShowReturnsErrorWhenHallIdIsNotValid() {
	show := []entities.Show{
		{
			MovieId: sql.NullInt64{
				Valid: true,
				Int64: 1,
			},
			HallId: sql.NullInt64{
				Valid: false,
				Int64: 1,
			},
			ShowDate: time.Date(2021, time.February, 14, 9, 30, 0, 0, time.UTC),
			TimingId: entities.Timing{
				Name:      "Morning",
				StartTime: time.Date(2021, time.February, 14, 9, 30, 0, 0, time.UTC),
				EndTime:   time.Date(2021, time.February, 14, 12, 30, 0, 0, time.UTC),
			},
			SeatPrice: 190.85,
		},
	}
	showReturned, err := suite.service.Add(show[0])
	assert.Equal(suite.T(), entities.Show{}, showReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.HALL_ID_MANDATORY_SHOW_CREATION).Error(), err.Error())
}

func (suite *showServiceTestSuite) TestAddShowReturnsErrorWhenShowDateIsNotPresent() {
	show := []entities.Show{
		{
			MovieId: sql.NullInt64{
				Valid: true,
				Int64: 1,
			},
			HallId: sql.NullInt64{
				Valid: true,
				Int64: 1,
			},
			TimingId: entities.Timing{
				Name:      "Morning",
				StartTime: time.Date(2021, time.February, 14, 9, 30, 0, 0, time.UTC),
				EndTime:   time.Date(2021, time.February, 14, 12, 30, 0, 0, time.UTC),
			},
			SeatPrice: 190.85,
		},
	}
	showReturned, err := suite.service.Add(show[0])
	assert.Equal(suite.T(), entities.Show{}, showReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.SHOW_DATE_MANDATORY_IN_SHOW_CREATION).Error(), err.Error())
}

func (suite *showServiceTestSuite) TestAddShowReturnsErrorWhenSeatPriceIsZero() {
	show := []entities.Show{
		{
			MovieId: sql.NullInt64{
				Valid: true,
				Int64: 1,
			},
			HallId: sql.NullInt64{
				Valid: true,
				Int64: 1,
			},
			ShowDate: time.Date(2021, time.February, 14, 9, 30, 0, 0, time.UTC),
			TimingId: entities.Timing{
				Name:      "Morning",
				StartTime: time.Date(2021, time.February, 14, 9, 30, 0, 0, time.UTC),
				EndTime:   time.Date(2021, time.February, 14, 12, 30, 0, 0, time.UTC),
			},
			SeatPrice: 0.0,
		},
	}
	showReturned, err := suite.service.Add(show[0])
	assert.Equal(suite.T(), entities.Show{}, showReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.SEAT_PRICE_MANDATORY_IN_SHOW_CREATION).Error(), err.Error())
}

func (suite *showServiceTestSuite) TestAddShowReturnsErrorWhenShowTimingInformationIsMissing() {
	show := []entities.Show{
		{
			MovieId: sql.NullInt64{
				Valid: true,
				Int64: 1,
			},
			HallId: sql.NullInt64{
				Valid: true,
				Int64: 1,
			},
			ShowDate:  time.Date(2021, time.February, 14, 9, 30, 0, 0, time.UTC),
			SeatPrice: 190.85,
		},
	}
	showReturned, err := suite.service.Add(show[0])
	assert.Equal(suite.T(), entities.Show{}, showReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.SHOW_TIMING_DETAILS_MISSING_IN_SHOW_CREATION).Error(), err.Error())
}

func (suite *showServiceTestSuite) TestAddShowReturnsErrorWhenShowStartTimeInformationIsMissing() {
	show := []entities.Show{
		{
			MovieId: sql.NullInt64{
				Valid: true,
				Int64: 1,
			},
			HallId: sql.NullInt64{
				Valid: true,
				Int64: 1,
			},
			TimingId: entities.Timing{
				Name:    "Morning",
				EndTime: time.Date(2021, time.February, 14, 12, 30, 0, 0, time.UTC),
			},
			ShowDate:  time.Date(2021, time.February, 14, 9, 30, 0, 0, time.UTC),
			SeatPrice: 190.85,
		},
	}
	showReturned, err := suite.service.Add(show[0])
	assert.Equal(suite.T(), entities.Show{}, showReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.SHOW_START_TIME_MISSING_IN_SHOW_CREATION).Error(), err.Error())
}

func (suite *showServiceTestSuite) TestAddShowReturnsErrorWhenShowEndTimeInformationIsMissing() {
	show := []entities.Show{
		{
			MovieId: sql.NullInt64{
				Valid: true,
				Int64: 1,
			},
			HallId: sql.NullInt64{
				Valid: true,
				Int64: 1,
			},
			TimingId: entities.Timing{
				Name:      "Morning",
				StartTime: time.Date(2021, time.February, 14, 9, 30, 0, 0, time.UTC),
			},
			ShowDate:  time.Date(2021, time.February, 14, 9, 30, 0, 0, time.UTC),
			SeatPrice: 190.85,
		},
	}
	showReturned, err := suite.service.Add(show[0])
	assert.Equal(suite.T(), entities.Show{}, showReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.SHOW_END_TIME_MISSING_IN_SHOW_CREATION).Error(), err.Error())
}

func (suite *showServiceTestSuite) TestAddShowReturnsErrorWhenShowTimingDayClassificationInformationIsMissing() {
	show := []entities.Show{
		{
			MovieId: sql.NullInt64{
				Valid: true,
				Int64: 1,
			},
			HallId: sql.NullInt64{
				Valid: true,
				Int64: 1,
			},
			TimingId: entities.Timing{
				StartTime: time.Date(2021, time.February, 14, 9, 30, 0, 0, time.UTC),
				EndTime:   time.Date(2021, time.February, 14, 12, 30, 0, 0, time.UTC),
			},
			ShowDate:  time.Date(2021, time.February, 14, 9, 30, 0, 0, time.UTC),
			SeatPrice: 190.85,
		},
	}
	showReturned, err := suite.service.Add(show[0])
	assert.Equal(suite.T(), entities.Show{}, showReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.SHOW_NAME_MISSING_IN_SHOW_CREATION).Error(), err.Error())
}

func (suite *showServiceTestSuite) TestAddShowReturnsErrorWhenHallIdDoNotExists() {
	suite.mockHallRepository.EXPECT().FetchHallByHallId(gomock.Any()).Return([]entities.Hall{}, nil)
	show := []entities.Show{
		{
			MovieId: sql.NullInt64{
				Valid: true,
				Int64: 1,
			},
			HallId: sql.NullInt64{
				Valid: true,
				Int64: 1,
			},
			TimingId: entities.Timing{
				Name:      "Morning",
				StartTime: time.Date(2021, time.February, 14, 9, 30, 0, 0, time.UTC),
				EndTime:   time.Date(2021, time.February, 14, 12, 30, 0, 0, time.UTC),
			},
			ShowDate:  time.Date(2021, time.February, 14, 9, 30, 0, 0, time.UTC),
			SeatPrice: 190.85,
		},
	}
	showReturned, err := suite.service.Add(show[0])
	assert.Equal(suite.T(), entities.Show{}, showReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), utils.WrapValidationError(errors.New(constants.FAILED_FETCHING_RESULT_FROM_DATABASE), fmt.Sprintf(constants.HALL_WITH_GIVEN_ID_DO_NOT_EXISTS, show[0].HallId.Int64)).Error(), err.Error())
}

func (suite *showServiceTestSuite) TestAddShowReturnsErrorWhenShowAlreadyExists() {
	expectedhall := []entities.Hall{
		{
			Id:    1,
			Name:  "Hall_A",
			Seats: 200,
			TheatreId: sql.NullInt64{
				Valid: true,
				Int64: 1,
			},
		},
	}
	show := []entities.Show{
		{
			MovieId: sql.NullInt64{
				Valid: true,
				Int64: 1,
			},
			HallId: sql.NullInt64{
				Valid: true,
				Int64: 1,
			},
			TimingId: entities.Timing{
				Name:      "Morning",
				StartTime: time.Date(2021, time.February, 14, 9, 30, 0, 0, time.UTC),
				EndTime:   time.Date(2021, time.February, 14, 12, 30, 0, 0, time.UTC),
			},
			ShowDate:  time.Date(2021, time.February, 14, 9, 30, 0, 0, time.UTC),
			SeatPrice: 190.85,
		},
	}
	suite.mockHallRepository.EXPECT().FetchHallByHallId(gomock.Any()).Return(expectedhall, nil)
	suite.mockShowRepository.EXPECT().FetchShowByMovieIdHallIdShowDate(gomock.Any()).Return(show, nil)
	showReturned, err := suite.service.Add(show[0])
	assert.Equal(suite.T(), entities.Show{}, showReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), fmt.Sprintf(constants.SHOW_ALREADY_EXIST, show[0].MovieId.Int64, show[0].HallId.Int64, show[0].ShowDate)).Error(), err.Error())
}

func (suite *showServiceTestSuite) TestShowIsCreatedSuccessfully() {
	expectedhall := []entities.Hall{
		{
			Id:    1,
			Name:  "Hall_A",
			Seats: 200,
			TheatreId: sql.NullInt64{
				Valid: true,
				Int64: 1,
			},
		},
	}
	show := []entities.Show{
		{
			MovieId: sql.NullInt64{
				Valid: true,
				Int64: 1,
			},
			HallId: sql.NullInt64{
				Valid: true,
				Int64: 1,
			},
			TimingId: entities.Timing{
				Name:      "Morning",
				StartTime: time.Date(2021, time.February, 14, 9, 30, 0, 0, time.UTC),
				EndTime:   time.Date(2021, time.February, 14, 12, 30, 0, 0, time.UTC),
			},
			ShowDate:       time.Date(2021, time.February, 14, 9, 30, 0, 0, time.UTC),
			SeatPrice:      190.85,
			AvailableSeats: 200,
		},
	}
	suite.mockHallRepository.EXPECT().FetchHallByHallId(gomock.Any()).Return(expectedhall, nil)
	suite.mockShowRepository.EXPECT().FetchShowByMovieIdHallIdShowDate(gomock.Any()).Return(nil, nil)
	suite.mockShowRepository.EXPECT().InsertShow(gomock.Any()).Return(nil)
	showReturned, err := suite.service.Add(show[0])
	assert.Equal(suite.T(), show[0], showReturned)
	assert.Nil(suite.T(), err)
}

func TestShowServiceTestSuite(t *testing.T) {
	suite.Run(t, new(showServiceTestSuite))
}
