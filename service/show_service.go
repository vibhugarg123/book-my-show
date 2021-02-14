package service

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/vibhugarg123/book-my-show/appcontext"
	"github.com/vibhugarg123/book-my-show/constants"
	"github.com/vibhugarg123/book-my-show/entities"
	"github.com/vibhugarg123/book-my-show/repository"
	"github.com/vibhugarg123/book-my-show/validation"
)

type ShowService interface {
	Add(entities.Show) (entities.Show, error)
}

type showService struct {
	showRepository repository.ShowRepository
	hallRepository repository.HallRepository
}

func (s showService) Add(show entities.Show) (entities.Show, error) {
	err := validation.AddNewShowValidation(show)
	if err != nil {
		return entities.Show{}, err
	}
	hallexists, err := s.hallRepository.FetchHallByHallId(show.HallId.Int64)
	if err != nil {
		return entities.Show{}, err
	}
	if len(hallexists) <= 0 {
		appcontext.Logger.Error().
			Str(constants.FAILED_FETCHING_RESULT_FROM_DATABASE, fmt.Sprintf(constants.HALL_WITH_GIVEN_ID_DO_NOT_EXISTS, show.HallId.Int64)).
			Msg(fmt.Sprintf(constants.HALL_WITH_GIVEN_ID_DO_NOT_EXISTS, show.HallId.Int64))
		return entities.Show{}, errors.Wrap(errors.New(constants.FAILED_FETCHING_RESULT_FROM_DATABASE), fmt.Sprintf(constants.HALL_WITH_GIVEN_ID_DO_NOT_EXISTS, show.HallId.Int64))
	}
	if show.AvailableSeats == 0 || (show.AvailableSeats > hallexists[0].Seats) {
		show.AvailableSeats = hallexists[0].Seats
	}
	existingShow, err := s.showRepository.FetchShowByMovieIdHallIdShowDate(show)
	if err != nil {
		appcontext.Logger.Error().
			Str(constants.FAILED_FETCHING_RESULT_FROM_DATABASE, err.Error()).
			Msg(constants.FAILED_GET_DB_CALL)
		return entities.Show{}, errors.Wrap(errors.New(constants.FAILED_FETCHING_RESULT_FROM_DATABASE), err.Error())
	}
	appcontext.Logger.Info().Msgf("existing shows %v", existingShow)
	if len(existingShow) > 0 || (len(existingShow) != 0 && existingShow[0] != (entities.Show{}) && !existingShow[0].HallId.Valid) || (len(existingShow) != 0 && existingShow[0] != (entities.Show{}) && !existingShow[0].MovieId.Valid) {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, fmt.Sprintf(constants.SHOW_ALREADY_EXIST, show.MovieId.Int64, show.HallId.Int64, show.ShowDate)).
			Msg(fmt.Sprintf(constants.SHOW_ALREADY_EXIST, show.MovieId.Int64, show.HallId.Int64, show.ShowDate))
		return entities.Show{}, errors.Wrap(errors.New(constants.REQUEST_INVALID), fmt.Sprintf(constants.SHOW_ALREADY_EXIST, show.MovieId.Int64, show.HallId.Int64, show.ShowDate))
	}
	err = s.showRepository.InsertShow(&show)
	if err != nil {
		appcontext.Logger.Error().
			Str(constants.FAILED_INSERT_DATABASE, err.Error()).
			Msg(err.Error())
		return entities.Show{}, errors.Wrap(errors.New(constants.ADD_SHOW_FAILED), err.Error())
	}
	return show, nil
}

func NewShowService() ShowService {
	return &showService{
		showRepository: repository.NewShowRepository(),
		hallRepository: repository.NewHallRepository(),
	}
}
