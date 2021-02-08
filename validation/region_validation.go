package validation

import (
	"github.com/pkg/errors"
	"github.com/vibhugarg123/book-my-show/appcontext"
	"github.com/vibhugarg123/book-my-show/constants"
	"github.com/vibhugarg123/book-my-show/entities"
)

func AddNewRegionValidation(region entities.Region) error {
	if region == (entities.Region{}) {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.EMPTY_REQUEST_BODY).
			Msg(constants.EMPTY_REQUEST_BODY)
		return errors.Wrap(errors.New(constants.REQUEST_INVALID), constants.EMPTY_REQUEST_BODY)
	}
	if region.Id == 0 {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.REGION_ID_MANDATORY).
			Msg(constants.REGION_ID_MANDATORY)
		return errors.Wrap(errors.New(constants.REQUEST_INVALID), constants.REGION_ID_MANDATORY)
	}
	if len(region.Name) == 0 {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.REGION_NAME_MANDATORY).
			Msg(constants.REGION_NAME_MANDATORY)
		return errors.Wrap(errors.New(constants.REQUEST_INVALID), constants.REGION_NAME_MANDATORY)
	}
	if region.RegionType == 0 {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.REGION_TYPE_MANDATORY).
			Msg(constants.REGION_TYPE_MANDATORY)
		return errors.Wrap(errors.New(constants.REQUEST_INVALID), constants.REGION_TYPE_MANDATORY)
	}
	return nil
}
