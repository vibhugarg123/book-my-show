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

type RegionService interface {
	Add(entities.Region) (entities.Region, error)
	GetRegionById(int) (entities.Region, error)
}

type regionService struct {
	regionRepository repository.RegionRepository
}

func (r regionService) Add(region entities.Region) (entities.Region, error) {
	err := validation.CreateNewRegionValidator(region)
	if err != nil {
		return entities.Region{}, err
	}
	existingRegion, err := r.regionRepository.FetchRegionById(int64(region.Id))
	if err != nil {
		appcontext.Logger.Error().
			Str(constants.FAILED_FETCHING_RESULT_FROM_DATABASE, err.Error()).
			Msg(constants.FAILED_GET_DB_CALL)
		return entities.Region{}, utils.WrapValidationError(errors.New(constants.FAILED_FETCHING_RESULT_FROM_DATABASE), err.Error())
	}
	if len(existingRegion) > 0 {
		appcontext.Logger.Error().
			Str(constants.REGION_CREATION_FAILED, fmt.Sprintf(constants.REGION_ALREADY_EXISTS, region.Id)).
			Msg(fmt.Sprintf(constants.REGION_ALREADY_EXISTS, region.Id))
		return entities.Region{}, utils.WrapValidationError(errors.New(constants.REGION_CREATION_FAILED), fmt.Sprintf(constants.REGION_ALREADY_EXISTS, region.Id))
	}
	err = r.regionRepository.InsertRegion(region)
	if err != nil {
		appcontext.Logger.Error().
			Str(constants.FAILED_INSERT_DATABASE, err.Error()).
			Msg(err.Error())
		return entities.Region{}, utils.WrapValidationError(errors.New(constants.ADD_REGION_FAILED), err.Error())
	}
	return region, nil
}

func (r regionService) GetRegionById(regionId int) (entities.Region, error) {
	region, err := r.regionRepository.FetchRegionById(int64(regionId))
	if err != nil {
		appcontext.Logger.Error().
			Str(constants.FAILED_FETCHING_RESULT_FROM_DATABASE, err.Error()).
			Msg(constants.FAILED_GET_DB_CALL)
		return entities.Region{}, utils.WrapValidationError(errors.New(constants.FAILED_FETCHING_RESULT_FROM_DATABASE), err.Error())
	}
	if len(region) == 0 {
		appcontext.Logger.Error().
			Str(constants.REGION_DO_NOT_EXIST, fmt.Sprintf(constants.REGION_DOES_NOT_EXIST, regionId)).
			Msg(fmt.Sprintf(constants.REGION_DOES_NOT_EXIST, regionId))
		return entities.Region{}, utils.WrapValidationError(errors.New(constants.REGION_CREATION_FAILED), fmt.Sprintf(constants.REGION_DOES_NOT_EXIST, regionId))
	}
	return region[0], nil
}

func NewRegionService() RegionService {
	return &regionService{
		regionRepository: repository.NewRegionRepository(),
	}
}
