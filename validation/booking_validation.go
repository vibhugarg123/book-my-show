package validation

import (
	"github.com/pkg/errors"
	"github.com/vibhugarg123/book-my-show/appcontext"
	"github.com/vibhugarg123/book-my-show/constants"
	"github.com/vibhugarg123/book-my-show/entities"
	"github.com/vibhugarg123/book-my-show/utils"
)

func AddNewBookingValidation(booking entities.Booking) error {
	if booking == (entities.Booking{}) {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.EMPTY_REQUEST_BODY).
			Msg(constants.EMPTY_REQUEST_BODY)
		return utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.EMPTY_REQUEST_BODY)
	}
	if booking.UserId == 0 {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.USER_ID_MANDATORY_IN_BOOKING_REQUEST).
			Msg(constants.USER_ID_MANDATORY_IN_BOOKING_REQUEST)
		return utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.USER_ID_MANDATORY_IN_BOOKING_REQUEST)
	}
	if booking.ShowId == 0 {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.SHOW_ID_MANDATORY_IN_BOOKING_REQUEST).
			Msg(constants.SHOW_ID_MANDATORY_IN_BOOKING_REQUEST)
		return utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.SHOW_ID_MANDATORY_IN_BOOKING_REQUEST)
	}
	if booking.Seats <= 0 {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.NUMBER_OF_SEATS_SHOULD_BE_VALID_BOOKING_REQUEST).
			Msg(constants.NUMBER_OF_SEATS_SHOULD_BE_VALID_BOOKING_REQUEST)
		return utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.NUMBER_OF_SEATS_SHOULD_BE_VALID_BOOKING_REQUEST)
	}
	return nil
}
