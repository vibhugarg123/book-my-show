package validation

import (
	"github.com/pkg/errors"
	"github.com/vibhugarg123/book-my-show/appcontext"
	"github.com/vibhugarg123/book-my-show/constants"
	"github.com/vibhugarg123/book-my-show/entities"
)

func AddNewTheatreValidation(theatre entities.Theatre) error {
	if theatre == (entities.Theatre{}) {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.EMPTY_REQUEST_BODY).
			Msg(constants.EMPTY_REQUEST_BODY)
		return errors.Wrap(errors.New(constants.REQUEST_INVALID), constants.EMPTY_REQUEST_BODY)
	}
	if len(theatre.Name) == 0 {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.THEATRE_NAME_MANDATORY).
			Msg(constants.THEATRE_NAME_MANDATORY)
		return errors.Wrap(errors.New(constants.REQUEST_INVALID), constants.THEATRE_NAME_MANDATORY)
	}
	if len(theatre.Address) == 0 {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.THEATRE_ADDRESS_MANDATORY).
			Msg(constants.THEATRE_ADDRESS_MANDATORY)
		return errors.Wrap(errors.New(constants.REQUEST_INVALID), constants.THEATRE_ADDRESS_MANDATORY)
	}
	if !theatre.RegionId.Valid {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.THEATRE_REGION_ID_MANDATORY).
			Msg(constants.THEATRE_REGION_ID_MANDATORY)
		return errors.Wrap(errors.New(constants.REQUEST_INVALID), constants.THEATRE_REGION_ID_MANDATORY)
	}
	return nil
}
