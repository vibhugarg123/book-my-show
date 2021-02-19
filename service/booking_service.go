package service

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vibhugarg123/book-my-show/appcontext"
	"github.com/vibhugarg123/book-my-show/constants"
	"github.com/vibhugarg123/book-my-show/entities"
	"github.com/vibhugarg123/book-my-show/repository"
	"github.com/vibhugarg123/book-my-show/utils"
	"github.com/vibhugarg123/book-my-show/validation"
	"strings"
)

type BookingService interface {
	Add(entities.Booking) (entities.Booking, error)
	GetBooking(int) ([]entities.Booking, error)
}

type bookingService struct {
	bookingRepository repository.BookingRepository
	showRepository    repository.ShowRepository
	movieRepository   repository.MovieRepository
}

func (b bookingService) Add(booking entities.Booking) (entities.Booking, error) {
	err := validation.ValidateForNewBooking(booking)
	if err != nil {
		return entities.Booking{}, err
	}
	showReturned, err := b.ValidateForShowId(booking)
	if err != nil {
		return entities.Booking{}, err
	}
	if err := b.UpdateAvailableSeats(&booking, showReturned); err != nil {
		return entities.Booking{}, err
	}
	if err := b.bookingRepository.InsertBooking(&booking); err != nil {
		appcontext.Logger.Error().
			Str(constants.CREATE_BOOKING_FAILED, err.Error()).
			Msg(err.Error())
		if utils.SqlError(err).Error() == constants.FOREIGN_KEY_VIOLATION && strings.Contains(err.Error(), "fk_user_id") {
			return entities.Booking{}, utils.WrapValidationError(errors.New(constants.FOREIGN_KEY_VIOLATION), fmt.Sprintf(constants.USER_ID_FOREIGN_KEY_VIOLATION_IN_CREATE_BOOKING, booking.UserId))
		}
		if utils.SqlError(err).Error() == constants.FOREIGN_KEY_VIOLATION && strings.Contains(err.Error(), "fk_show_id") {
			return entities.Booking{}, utils.WrapValidationError(errors.New(constants.FOREIGN_KEY_VIOLATION), fmt.Sprintf(constants.SHOW_ID_FOREIGN_KEY_VIOLATION_IN_CREATE_BOOKING, booking.ShowId))
		}
		return entities.Booking{}, errors.Wrap(errors.New(constants.CREATE_BOOKING_FAILED), err.Error())
	}
	booking.MovieId = int(showReturned[0].MovieId.Int64)
	return booking, nil
}

func (b bookingService) UpdateAvailableSeats(booking *entities.Booking, showReturned []entities.Show) error {
	booking.TotalPrice = showReturned[0].SeatPrice * float64(booking.Seats)
	availableSeats := showReturned[0].AvailableSeats - booking.Seats
	if err := b.showRepository.UpdateSeatsByShowId(availableSeats, booking.ShowId); err != nil {
		return err
	}
	return nil
}

func (b bookingService) ValidateForShowId(booking entities.Booking) ([]entities.Show, error) {
	showReturned, err := b.showRepository.FetchShowByShowId(booking.ShowId)
	if err != nil {
		appcontext.Logger.Error().
			Str(constants.FAILED_GET_DB_CALL, err.Error()).
			Msg(err.Error())
		return []entities.Show{}, utils.WrapValidationError(errors.New(constants.FAILED_GET_DB_CALL), err.Error())
	}
	if len(showReturned) == 0 {
		appcontext.Logger.Error().
			Str(constants.SHOW_DOES_NOT_EXIT, constants.SHOW_DOES_NOT_EXIST_FOR_GIVEN_SHOW_ID).
			Msg(constants.SHOW_DOES_NOT_EXIST_FOR_GIVEN_SHOW_ID)
		return []entities.Show{}, utils.WrapValidationError(errors.New(constants.SHOW_DOES_NOT_EXIT), constants.SHOW_DOES_NOT_EXIST_FOR_GIVEN_SHOW_ID)
	}
	if showReturned[0].AvailableSeats-booking.Seats < 0 {
		appcontext.Logger.Error().
			Str(constants.AVAILABLE_SEATS_LESS_COMPARED_TO_REQUESTED, fmt.Sprintf(constants.LESS_SEATS_AVAILABLE, showReturned[0].AvailableSeats, booking.Seats)).
			Msg(fmt.Sprintf(constants.LESS_SEATS_AVAILABLE, showReturned[0].AvailableSeats, booking.Seats))
		return []entities.Show{}, utils.WrapValidationError(errors.New(constants.AVAILABLE_SEATS_LESS_COMPARED_TO_REQUESTED), fmt.Sprintf(constants.LESS_SEATS_AVAILABLE, showReturned[0].AvailableSeats, booking.Seats))
	}
	return showReturned, nil
}

func (b bookingService) GetBooking(userId int) ([]entities.Booking, error) {
	bookings, err := b.bookingRepository.FetchBookingByUserId(userId)
	if err != nil {
		appcontext.Logger.Error().
			Str(constants.FAILED_FETCHING_RESULT_FROM_DATABASE, err.Error()).
			Msg(constants.FAILED_GET_DB_CALL)
		return []entities.Booking{}, utils.WrapValidationError(errors.New(constants.FAILED_FETCHING_RESULT_FROM_DATABASE), err.Error())
	}
	if len(bookings) == 0 {
		appcontext.Logger.Error().
			Str(constants.BOOKING_DO_NOT_EXIST, fmt.Sprintf(constants.BOOKING_DO_NOT_EXIST_FOR_GIVEN_USER_ID, userId)).
			Msg(fmt.Sprintf(constants.BOOKING_DO_NOT_EXIST_FOR_GIVEN_USER_ID, userId))
		return []entities.Booking{}, utils.WrapValidationError(errors.New(constants.BOOKING_DO_NOT_EXIST), fmt.Sprintf(constants.BOOKING_DO_NOT_EXIST_FOR_GIVEN_USER_ID, userId))
	}
	return bookings, nil
}

func NewBookingService() BookingService {
	return &bookingService{
		bookingRepository: repository.NewBookingRepository(),
		showRepository:    repository.NewShowRepository(),
		movieRepository:   repository.NewMovieRepository(),
	}
}
