package validation

import (
	"github.com/pkg/errors"
	"github.com/vibhugarg123/book-my-show/appcontext"
	"github.com/vibhugarg123/book-my-show/constants"
	"github.com/vibhugarg123/book-my-show/entities"
	"github.com/vibhugarg123/book-my-show/utils"
)

func ValidateForNewUser(user entities.User) error {
	if user == (entities.User{}) {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.EMPTY_REQUEST_BODY).
			Msg(constants.EMPTY_REQUEST_BODY)
		return utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.EMPTY_REQUEST_BODY)
	}
	if len(user.EmailId) == 0 {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.EMAIL_ID_MANDATORY).
			Msg(constants.EMAIL_ID_MANDATORY)
		return utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.EMAIL_ID_MANDATORY)
	}
	if len(user.Password) == 0 {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.PASSWORD_MANDATORY).
			Msg(constants.PASSWORD_MANDATORY)
		return utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.PASSWORD_MANDATORY)
	}
	if len(user.FirstName) == 0 {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.FIRST_NAME_MANDATORY).
			Msg(constants.FIRST_NAME_MANDATORY)
		return utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.FIRST_NAME_MANDATORY)
	}
	return nil
}

func ValidateForLogin(user entities.User) error {
	if user == (entities.User{}) {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.EMPTY_REQUEST_BODY).
			Msg(constants.EMPTY_REQUEST_BODY)
		return utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.EMPTY_REQUEST_BODY)
	}
	if len(user.EmailId) == 0 {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.EMAIL_ID_MANDATORY).
			Msg(constants.EMAIL_ID_MANDATORY)
		return utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.EMAIL_ID_MANDATORY)
	}
	if len(user.Password) == 0 {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.PASSWORD_MANDATORY).
			Msg(constants.PASSWORD_MANDATORY)
		return utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.PASSWORD_MANDATORY)
	}
	return nil
}
