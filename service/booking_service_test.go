package service

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
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

type bookingServiceTestSuite struct {
	suite.Suite
	bookingRepository *repository.MockBookingRepository
	movieRepository   *repository.MockMovieRepository
	showRepository    *repository.MockShowRepository
	service           BookingService
}

func (suite *bookingServiceTestSuite) SetupTest() {
	config.LoadTestConfig()
	appcontext.Init()
	mockController := gomock.NewController(suite.T())
	suite.bookingRepository = repository.NewMockBookingRepository(mockController)
	suite.movieRepository = repository.NewMockMovieRepository(mockController)
	suite.showRepository = repository.NewMockShowRepository(mockController)
	suite.service = &bookingService{
		bookingRepository: suite.bookingRepository,
		showRepository:    suite.showRepository,
		movieRepository:   suite.movieRepository,
	}
}

func (suite *bookingServiceTestSuite) TestAddBookingReturnsErrorWhenUserIdIsMissing() {
	booking := []entities.Booking{
		{
			Seats:  5,
			ShowId: 1,
		},
	}
	bookingReturned, err := suite.service.Add(booking[0])
	assert.Equal(suite.T(), entities.Booking{}, bookingReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.USER_ID_MANDATORY_IN_BOOKING_REQUEST).Error(), err.Error())
}

func (suite *bookingServiceTestSuite) TestAddHallReturnsErrorShowIdMissing() {
	booking := []entities.Booking{
		{
			UserId: 1,
			Seats:  5,
		},
	}
	bookingReturned, err := suite.service.Add(booking[0])
	assert.Equal(suite.T(), entities.Booking{}, bookingReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.SHOW_ID_MANDATORY_IN_BOOKING_REQUEST).Error(), err.Error())
}

func (suite *bookingServiceTestSuite) TestAddBookingReturnsErrorWhenSeatsNeededAreMissing() {
	booking := []entities.Booking{
		{
			ShowId: 1,
			UserId: 1,
		},
	}
	bookingReturned, err := suite.service.Add(booking[0])
	assert.Equal(suite.T(), entities.Booking{}, bookingReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.NUMBER_OF_SEATS_SHOULD_BE_VALID_BOOKING_REQUEST).Error(), err.Error())
}

func (suite *bookingServiceTestSuite) TestCreateBookingReturnsErrorWhenShowIdIsNotValid() {
	booking := []entities.Booking{
		{
			ShowId: 1,
			UserId: 1,
			Seats:  5,
		},
	}
	suite.showRepository.EXPECT().FetchShowByShowId(gomock.Any()).Return(nil, nil)
	bookingReturned, err := suite.service.Add(booking[0])
	assert.Equal(suite.T(), entities.Booking{}, bookingReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), utils.WrapValidationError(errors.New(constants.SHOW_DOES_NOT_EXIT), constants.SHOW_DOES_NOT_EXIST_FOR_GIVEN_SHOW_ID).Error(), err.Error())
}

func (suite *bookingServiceTestSuite) TestCreateBookingReturnsErrorSeatsAreLessComparedToSeatsRequested() {
	booking := []entities.Booking{
		{
			ShowId: 1,
			UserId: 1,
			Seats:  5,
		},
	}
	expectedShow := []entities.Show{
		{
			Id: 1,
			MovieId: sql.NullInt64{
				Valid: true,
				Int64: 5,
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
			AvailableSeats: 4,
		},
	}
	suite.showRepository.EXPECT().FetchShowByShowId(gomock.Any()).Return(expectedShow, nil)
	bookingReturned, err := suite.service.Add(booking[0])
	assert.Equal(suite.T(), entities.Booking{}, bookingReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), utils.WrapValidationError(errors.New(constants.AVAILABLE_SEATS_LESS_COMPARED_TO_REQUESTED), fmt.Sprintf(constants.LESS_SEATS_AVAILABLE, expectedShow[0].AvailableSeats, booking[0].Seats)).Error(), err.Error())
}

func (suite *bookingServiceTestSuite) TestCreateBookingReturnsErrorWhenForeignKeyConstraintUserIdFails() {
	booking := []entities.Booking{
		{
			ShowId: 1,
			UserId: 1,
			Seats:  5,
		},
	}
	expectedShow := []entities.Show{
		{
			Id: 1,
			MovieId: sql.NullInt64{
				Valid: true,
				Int64: 5,
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
			AvailableSeats: 195,
		},
	}
	suite.showRepository.EXPECT().FetchShowByShowId(gomock.Any()).Return(expectedShow, nil)
	suite.showRepository.EXPECT().UpdateSeatsByShowId(gomock.Any(), gomock.Any()).Return(nil)
	expectedMySqlError := error(&mysql.MySQLError{
		Number:  1452,
		Message: "Error 1452: Cannot add or update a child row: a foreign key constraint fails (`BOOK_MY_SHOW`.`bookings`, CONSTRAINT `fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`))",
	})
	suite.bookingRepository.EXPECT().InsertBooking(gomock.Any()).Return(expectedMySqlError)
	bookingReturned, err := suite.service.Add(booking[0])
	assert.Equal(suite.T(), entities.Booking{}, bookingReturned)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), utils.WrapValidationError(errors.New(constants.FOREIGN_KEY_VIOLATION), fmt.Sprintf(constants.USER_ID_FOREIGN_KEY_VIOLATION_IN_CREATE_BOOKING, booking[0].UserId)).Error(), err.Error())
}

func (suite *bookingServiceTestSuite) TestBookingIsSuccessfullyCreated() {
	booking := []entities.Booking{
		{
			ShowId:     1,
			UserId:     1,
			Seats:      5,
			MovieId:    5,
			TotalPrice: 954.25,
		},
	}
	expectedShow := []entities.Show{
		{
			Id: 1,
			MovieId: sql.NullInt64{
				Valid: true,
				Int64: 5,
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
			AvailableSeats: 195,
		},
	}
	suite.showRepository.EXPECT().FetchShowByShowId(gomock.Any()).Return(expectedShow, nil)
	suite.showRepository.EXPECT().UpdateSeatsByShowId(gomock.Any(), gomock.Any()).Return(nil)

	suite.bookingRepository.EXPECT().InsertBooking(gomock.Any()).Return(nil)
	bookingReturned, err := suite.service.Add(booking[0])
	assert.Equal(suite.T(), booking[0], bookingReturned)
	assert.Nil(suite.T(), err)
}

func TestBookingServiceTestSuite(t *testing.T) {
	suite.Run(t, new(bookingServiceTestSuite))
}
