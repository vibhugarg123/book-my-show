package validation

import (
	"github.com/pkg/errors"
	"github.com/vibhugarg123/book-my-show/appcontext"
	"github.com/vibhugarg123/book-my-show/constants"
	"github.com/vibhugarg123/book-my-show/entities"
	"github.com/vibhugarg123/book-my-show/utils"
)

func ValidateForNewHall(hall entities.Hall) error {
	if hall == (entities.Hall{}) {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.EMPTY_REQUEST_BODY).
			Msg(constants.EMPTY_REQUEST_BODY)
		return utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.EMPTY_REQUEST_BODY)
	}
	if len(hall.Name) == 0 {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.HALL_NAME_MANDATORY).
			Msg(constants.HALL_NAME_MANDATORY)
		return utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.HALL_NAME_MANDATORY)
	}
	if hall.Seats <= 0 {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.SEATS_NUMBER_INVALID).
			Msg(constants.SEATS_NUMBER_INVALID)
		return utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.SEATS_NUMBER_INVALID)
	}
	if !hall.TheatreId.Valid {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.THEATRE_ID_HALL_CREATION_INVALID).
			Msg(constants.THEATRE_ID_HALL_CREATION_INVALID)
		return utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.THEATRE_ID_HALL_CREATION_INVALID)
	}
	return nil
}
