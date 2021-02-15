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

type HallService interface {
	Add(entities.Hall) (entities.Hall, error)
	GetHallByTheatreId(int) ([]entities.Hall, error)
}

type hallService struct {
	hallRepository repository.HallRepository
}

func (h hallService) Add(hall entities.Hall) (entities.Hall, error) {
	err := validation.CreateNewHallValidator(hall)
	if err != nil {
		return entities.Hall{}, err
	}
	existingHall, err := h.hallRepository.FetchHallByNameAndTheatreId(hall)
	if err != nil {
		appcontext.Logger.Error().
			Str(constants.FAILED_FETCHING_RESULT_FROM_DATABASE, err.Error()).
			Msg(constants.FAILED_GET_DB_CALL)
		return entities.Hall{}, utils.WrapValidationError(errors.New(constants.FAILED_FETCHING_RESULT_FROM_DATABASE), err.Error())
	}
	if len(existingHall) > 0 {
		appcontext.Logger.Error().
			Str(constants.HALL_CREATION_FAILED, fmt.Sprintf(constants.HALL_ALREADY_EXISTS, hall.Name, hall.TheatreId.Int64)).
			Msg(fmt.Sprintf(constants.HALL_ALREADY_EXISTS, hall.Name, hall.TheatreId.Int64))
		return entities.Hall{}, utils.WrapValidationError(errors.New(constants.HALL_CREATION_FAILED), fmt.Sprintf(constants.HALL_ALREADY_EXISTS, hall.Name, hall.TheatreId.Int64))
	}
	err = h.hallRepository.InsertHall(hall)
	if err != nil {
		appcontext.Logger.Error().
			Str(constants.FAILED_INSERT_DATABASE, err.Error()).
			Msg(err.Error())
		return entities.Hall{}, utils.WrapValidationError(errors.New(constants.ADD_REGION_FAILED), err.Error())
	}
	return hall, nil
}

func (h hallService) GetHallByTheatreId(theatreId int) ([]entities.Hall, error) {
	halls, err := h.hallRepository.FetchHallByTheatreId(theatreId)
	if err != nil {
		appcontext.Logger.Error().
			Str(constants.FAILED_FETCHING_RESULT_FROM_DATABASE, err.Error()).
			Msg(constants.FAILED_GET_DB_CALL)
		return []entities.Hall{}, utils.WrapValidationError(errors.New(constants.FAILED_FETCHING_RESULT_FROM_DATABASE), err.Error())
	}
	if len(halls) == 0 {
		appcontext.Logger.Error().
			Str(constants.HALLS_DO_NOT_EXIST, fmt.Sprintf(constants.HALLS_DO_NOT_EXIST_WITH_THEATRE_ID, theatreId)).
			Msg(fmt.Sprintf(constants.HALLS_DO_NOT_EXIST_WITH_THEATRE_ID, theatreId))
		return []entities.Hall{}, utils.WrapValidationError(errors.New(constants.HALLS_DO_NOT_EXIST), fmt.Sprintf(constants.HALLS_DO_NOT_EXIST_WITH_THEATRE_ID, theatreId))
	}
	return halls, nil
}

func NewHallService() HallService {
	return &hallService{
		hallRepository: repository.NewHallRepository(),
	}
}
