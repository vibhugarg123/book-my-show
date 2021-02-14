package validation

import (
	"github.com/pkg/errors"
	"github.com/vibhugarg123/book-my-show/appcontext"
	"github.com/vibhugarg123/book-my-show/constants"
	"github.com/vibhugarg123/book-my-show/entities"
)

func AddNewShowValidation(show entities.Show) error {
	if show == (entities.Show{}) {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.EMPTY_REQUEST_BODY).
			Msg(constants.EMPTY_REQUEST_BODY)
		return errors.Wrap(errors.New(constants.REQUEST_INVALID), constants.EMPTY_REQUEST_BODY)
	}
	if !show.MovieId.Valid {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.MOVIE_ID_MANDATORY_SHOW_CREATION).
			Msg(constants.MOVIE_ID_MANDATORY_SHOW_CREATION)
		return errors.Wrap(errors.New(constants.REQUEST_INVALID), constants.MOVIE_ID_MANDATORY_SHOW_CREATION)
	}
	if !show.HallId.Valid {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.HALL_ID_MANDATORY_SHOW_CREATION).
			Msg(constants.HALL_ID_MANDATORY_SHOW_CREATION)
		return errors.Wrap(errors.New(constants.REQUEST_INVALID), constants.HALL_ID_MANDATORY_SHOW_CREATION)
	}
	if show.ShowDate.IsZero() {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.SHOW_DATE_MANDATORY_IN_SHOW_CREATION).
			Msg(constants.SHOW_DATE_MANDATORY_IN_SHOW_CREATION)
		return errors.Wrap(errors.New(constants.REQUEST_INVALID), constants.SHOW_DATE_MANDATORY_IN_SHOW_CREATION)
	}
	if show.SeatPrice <= 0.0 {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.SEAT_PRICE_MANDATORY_IN_SHOW_CREATION).
			Msg(constants.SEAT_PRICE_MANDATORY_IN_SHOW_CREATION)
		return errors.Wrap(errors.New(constants.REQUEST_INVALID), constants.SEAT_PRICE_MANDATORY_IN_SHOW_CREATION)
	}
	if show.TimingId == (entities.Timing{}) {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.SHOW_TIMING_DETAILS_MISSING_IN_SHOW_CREATION).
			Msg(constants.SHOW_TIMING_DETAILS_MISSING_IN_SHOW_CREATION)
		return errors.Wrap(errors.New(constants.REQUEST_INVALID), constants.SHOW_TIMING_DETAILS_MISSING_IN_SHOW_CREATION)
	}
	if show.TimingId.StartTime.IsZero() {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.SHOW_START_TIME_MISSING_IN_SHOW_CREATION).
			Msg(constants.SHOW_START_TIME_MISSING_IN_SHOW_CREATION)
		return errors.Wrap(errors.New(constants.REQUEST_INVALID), constants.SHOW_START_TIME_MISSING_IN_SHOW_CREATION)
	}
	if show.TimingId.EndTime.IsZero() {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.SHOW_END_TIME_MISSING_IN_SHOW_CREATION).
			Msg(constants.SHOW_END_TIME_MISSING_IN_SHOW_CREATION)
		return errors.Wrap(errors.New(constants.REQUEST_INVALID), constants.SHOW_END_TIME_MISSING_IN_SHOW_CREATION)
	}
	if len(show.TimingId.Name) == 0 {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.SHOW_NAME_MISSING_IN_SHOW_CREATION).
			Msg(constants.SHOW_NAME_MISSING_IN_SHOW_CREATION)
		return errors.Wrap(errors.New(constants.REQUEST_INVALID), constants.SHOW_NAME_MISSING_IN_SHOW_CREATION)
	}
	return nil
}
