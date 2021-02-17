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
)

type MovieService interface {
	Add(entities.Movie) (entities.Movie, error)
	GetActiveMovies() ([]entities.Movie, error)
}

type movieService struct {
	movieRepository repository.MovieRepository
}

func (m movieService) Add(movie entities.Movie) (entities.Movie, error) {
	err := validation.ValidateForNewMovie(movie)
	if err != nil {
		return entities.Movie{}, err
	}
	existingMovie, err := m.movieRepository.FetchMovieByName(movie.Name)
	if err != nil {
		appcontext.Logger.Error().
			Str(constants.FAILED_FETCHING_RESULT_FROM_DATABASE, err.Error()).
			Msg(constants.FAILED_GET_DB_CALL)
		return entities.Movie{}, utils.WrapValidationError(errors.New(constants.FAILED_FETCHING_RESULT_FROM_DATABASE), err.Error())
	}
	if len(existingMovie) > 0 {
		appcontext.Logger.Error().
			Str(constants.MOVIE_CREATION_FAILED, fmt.Sprintf(constants.MOVIE_ALREADY_EXISTS, movie.Name)).
			Msg(fmt.Sprintf(constants.MOVIE_ALREADY_EXISTS, movie.Name))
		return entities.Movie{}, utils.WrapValidationError(errors.New(constants.MOVIE_CREATION_FAILED), fmt.Sprintf(constants.MOVIE_ALREADY_EXISTS, movie.Name))
	}
	err = m.movieRepository.InsertMovie(movie)
	if err != nil {
		appcontext.Logger.Error().
			Str(constants.FAILED_INSERT_DATABASE, err.Error()).
			Msg(err.Error())
		return entities.Movie{}, utils.WrapValidationError(errors.New(constants.ADD_MOVIE_FAILED), err.Error())
	}
	return movie, nil
}

func (m movieService) GetActiveMovies() ([]entities.Movie, error) {
	movies, err := m.movieRepository.FetchActiveMovies()
	if err != nil {
		appcontext.Logger.Error().
			Str(constants.FAILED_FETCHING_RESULT_FROM_DATABASE, err.Error()).
			Msg(constants.FAILED_GET_DB_CALL)
		return []entities.Movie{}, utils.WrapValidationError(errors.New(constants.FAILED_FETCHING_RESULT_FROM_DATABASE), err.Error())
	}
	if len(movies) == 0 {
		appcontext.Logger.Error().
			Str(constants.NO_ACTIVE_MOVIES, constants.ACTIVE_MOVIES_NOT_PRESENT).
			Msg(constants.ACTIVE_MOVIES_NOT_PRESENT)
		return []entities.Movie{}, utils.WrapValidationError(errors.New(constants.NO_ACTIVE_MOVIES), constants.ACTIVE_MOVIES_NOT_PRESENT)
	}
	return movies, nil
}

func NewMovieService() MovieService {
	return &movieService{
		movieRepository: repository.NewMovieRepository(),
	}
}
