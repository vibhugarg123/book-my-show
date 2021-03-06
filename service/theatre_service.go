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
	"strings"
)

type TheatreService interface {
	Add(entities.Theatre) (entities.Theatre, error)
	GetTheatreByName(string) ([]entities.Theatre, error)
}

type theatreService struct {
	theatreRepository repository.TheatreRepository
	regionRepository  repository.RegionRepository
}

func (t theatreService) Add(theatre entities.Theatre) (entities.Theatre, error) {
	err := validation.ValidateForNewTheatre(theatre)
	if err != nil {
		return entities.Theatre{}, err
	}
	if err = t.ValidateForRegionIdExists(theatre.RegionId.Int64); err != nil {
		return entities.Theatre{}, err
	}
	if err = t.ValidateForTheatreToBeAddedExists(theatre); err != nil {
		return entities.Theatre{}, err
	}
	err = t.theatreRepository.InsertTheatre(theatre)
	if err != nil {
		appcontext.Logger.Error().
			Str(constants.FAILED_INSERT_DATABASE, err.Error()).
			Msg(err.Error())
		return entities.Theatre{}, utils.WrapValidationError(errors.New(constants.ADD_THEATRE_FAILED), err.Error())
	}
	return theatre, nil
}

func (t theatreService) ValidateForTheatreToBeAddedExists(theatre entities.Theatre) error {
	existingTheatres, err := t.theatreRepository.FetchTheatreByNameRegionIdAndAddress(theatre)
	if err != nil {
		appcontext.Logger.Error().
			Str(constants.FAILED_FETCHING_RESULT_FROM_DATABASE, err.Error()).
			Msg(constants.FAILED_GET_DB_CALL)
		return utils.WrapValidationError(errors.New(constants.FAILED_FETCHING_RESULT_FROM_DATABASE), err.Error())
	}
	if len(existingTheatres) > 0 {
		appcontext.Logger.Error().
			Str(constants.THEATRE_ALREADY_EXIST, fmt.Sprintf(constants.THEATRE_ALREADY_EXISTS, theatre)).
			Msg(fmt.Sprintf(constants.THEATRE_ALREADY_EXISTS, theatre))
		return utils.WrapValidationError(errors.New(constants.THEATRE_ALREADY_EXIST), fmt.Sprintf(constants.THEATRE_ALREADY_EXISTS, theatre))
	}
	return nil
}

func (t theatreService) ValidateForRegionIdExists(regionId int64) error {
	existingRegion, err := t.regionRepository.FetchRegionById(regionId)
	if err != nil {
		appcontext.Logger.Error().
			Str(constants.FAILED_FETCHING_RESULT_FROM_DATABASE, err.Error()).
			Msg(constants.FAILED_GET_DB_CALL)
		return utils.WrapValidationError(errors.New(constants.FAILED_FETCHING_RESULT_FROM_DATABASE), err.Error())
	}
	if len(existingRegion) == 0 {
		appcontext.Logger.Error().
			Str(constants.REGION_DO_NOT_EXIST, fmt.Sprintf(constants.REGION_DOES_NOT_EXIST, regionId)).
			Msg(fmt.Sprintf(constants.REGION_DOES_NOT_EXIST, regionId))
		return utils.WrapValidationError(errors.New(constants.REGION_DO_NOT_EXIST), fmt.Sprintf(constants.REGION_DOES_NOT_EXIST, regionId))
	}
	return nil
}

func (t theatreService) GetTheatreByName(theatreName string) ([]entities.Theatre, error) {
	if len(theatreName) == 0 || len(strings.TrimSpace(theatreName)) == 0 {
		appcontext.Logger.Error().
			Str(constants.REQUEST_INVALID, constants.THEATRE_NAME_MANDATORY).
			Msg(constants.THEATRE_NAME_MANDATORY)
		return []entities.Theatre{}, utils.WrapValidationError(errors.New(constants.REQUEST_INVALID), constants.THEATRE_NAME_MANDATORY)
	}
	theatres, err := t.theatreRepository.FetchTheatreByName(theatreName)
	if err != nil {
		appcontext.Logger.Error().
			Str(constants.FAILED_FETCHING_RESULT_FROM_DATABASE, err.Error()).
			Msg(constants.FAILED_GET_DB_CALL)
		return []entities.Theatre{}, utils.WrapValidationError(errors.New(constants.FAILED_FETCHING_RESULT_FROM_DATABASE), err.Error())
	}
	if len(theatres) == 0 {
		appcontext.Logger.Error().
			Str(constants.THEATRE_DO_NOT_EXIST, fmt.Sprintf(constants.THEATRE_DOES_NOT_EXIST, theatreName)).
			Msg(fmt.Sprintf(constants.THEATRE_DOES_NOT_EXIST, theatreName))
		return []entities.Theatre{}, utils.WrapValidationError(errors.New(constants.THEATRE_DO_NOT_EXIST), fmt.Sprintf(constants.THEATRE_DOES_NOT_EXIST, theatreName))
	}
	return theatres, nil
}

func NewTheatreService() TheatreService {
	return &theatreService{
		theatreRepository: repository.NewTheatreRepository(),
		regionRepository:  repository.NewRegionRepository(),
	}
}
