package validation

import (
	"github.com/pkg/errors"
	"github.com/vibhugarg123/book-my-show/appcontext"
	"github.com/vibhugarg123/book-my-show/constants"
	"github.com/vibhugarg123/book-my-show/entities"
	"github.com/vibhugarg123/book-my-show/utils"
)

func ValidateForNewMovie(movie entities.Movie) error {
	if movie == (entities.Movie{}) {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.EMPTY_REQUEST_BODY).
			Msg(constants.EMPTY_REQUEST_BODY)
		return utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.EMPTY_REQUEST_BODY)
	}
	if len(movie.Name) == 0 {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.MOVIE_NAME_MANDATORY).
			Msg(constants.MOVIE_NAME_MANDATORY)
		return utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.MOVIE_NAME_MANDATORY)
	}
	if len(movie.DirectorName) == 0 {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.DIRECTOR_NAME_MANDATORY).
			Msg(constants.DIRECTOR_NAME_MANDATORY)
		return utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.DIRECTOR_NAME_MANDATORY)
	}
	if movie.ReleaseDate.IsZero() {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.RELEASE_DATE_MANDATORY).
			Msg(constants.RELEASE_DATE_MANDATORY)
		return utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.RELEASE_DATE_MANDATORY)
	}
	return nil
}
