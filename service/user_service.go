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

type UserService interface {
	Add(user entities.User) (entities.User, error)
	Login(user entities.User) (entities.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService() UserService {
	return &userService{
		userRepository: repository.NewUserRepository(),
	}
}

func (u userService) Add(user entities.User) (entities.User, error) {
	err := validation.ValidateForNewUser(user)
	if err != nil {
		return entities.User{}, err
	}
	existingUser, err := u.userRepository.FetchUserByEmailId(user.EmailId)
	if err != nil {
		appcontext.Logger.Error().
			Str(constants.FAILED_FETCHING_RESULT_FROM_DATABASE, err.Error()).
			Msg(constants.FAILED_GET_DB_CALL)
		return entities.User{}, utils.WrapValidationError(errors.New(constants.FAILED_FETCHING_RESULT_FROM_DATABASE), err.Error())
	}
	if len(existingUser) > 0 {
		appcontext.Logger.Error().
			Str(constants.FAILED_CREATING_USER, fmt.Sprintf(constants.USER_ALREADY_EXISTS, user.EmailId)).
			Msg(fmt.Sprintf(constants.USER_ALREADY_EXISTS, user.EmailId))
		return entities.User{}, utils.WrapValidationError(errors.New(constants.FAILED_CREATING_USER), fmt.Sprintf(constants.USER_ALREADY_EXISTS, user.EmailId))
	}
	err = u.userRepository.InsertUser(user)
	if err != nil {
		appcontext.Logger.Error().
			Str(constants.FAILED_INSERT_DATABASE, err.Error()).
			Msg(err.Error())
		return entities.User{}, utils.WrapValidationError(errors.New(constants.USER_CREATION_FAILED), err.Error())
	}
	return user, nil
}

func (u userService) Login(user entities.User) (entities.User, error) {
	err := validation.ValidateForLogin(user)
	if err != nil {
		return entities.User{}, err
	}
	existingUser, err := u.userRepository.FetchUserByEmailIdAndPassword(user.EmailId, user.Password)
	if err != nil {
		appcontext.Logger.Error().
			Str(constants.FAILED_FETCHING_RESULT_FROM_DATABASE, err.Error()).
			Msg(err.Error())
		return entities.User{}, utils.WrapValidationError(errors.New(constants.FAILED_FETCHING_RESULT_FROM_DATABASE), err.Error())
	}
	if len(existingUser) == 0 {
		appcontext.Logger.Error().
			Str(constants.USER_DO_NOT_EXIST, constants.USER_DOES_NOT_EXIST).
			Msg(constants.USER_DOES_NOT_EXIST)
		return entities.User{}, utils.WrapValidationError(errors.New(constants.USER_DO_NOT_EXIST), constants.USER_DOES_NOT_EXIST)
	}
	return user, nil
}
