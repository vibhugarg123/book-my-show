package validation

import (
	"github.com/pkg/errors"
	"github.com/vibhugarg123/book-my-show/appcontext"
	"github.com/vibhugarg123/book-my-show/constants"
	"github.com/vibhugarg123/book-my-show/entities"
	"github.com/vibhugarg123/book-my-show/utils"
)

func ValidateForNewTheatre(theatre entities.Theatre) error {
	if theatre == (entities.Theatre{}) {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.EMPTY_REQUEST_BODY).
			Msg(constants.EMPTY_REQUEST_BODY)
		return utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.EMPTY_REQUEST_BODY)
	}
	if len(theatre.Name) == 0 {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.THEATRE_NAME_MANDATORY).
			Msg(constants.THEATRE_NAME_MANDATORY)
		return utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.THEATRE_NAME_MANDATORY)
	}
	if len(theatre.Address) == 0 {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.THEATRE_ADDRESS_MANDATORY).
			Msg(constants.THEATRE_ADDRESS_MANDATORY)
		return utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.THEATRE_ADDRESS_MANDATORY)
	}
	if !theatre.RegionId.Valid {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.THEATRE_REGION_ID_MANDATORY).
			Msg(constants.THEATRE_REGION_ID_MANDATORY)
		return utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.THEATRE_REGION_ID_MANDATORY)
	}
	return nil
}
