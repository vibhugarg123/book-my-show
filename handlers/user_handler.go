package handlers

import (
	"encoding/json"
	"github.com/vibhugarg123/book-my-show/appcontext"
	"github.com/vibhugarg123/book-my-show/constants"
	"github.com/vibhugarg123/book-my-show/domain"
	"github.com/vibhugarg123/book-my-show/entities"
	"github.com/vibhugarg123/book-my-show/service"
	"github.com/vibhugarg123/book-my-show/utils"
	"net/http"
)

type AddUserHandler struct {
	service service.UserService
}

func NewAddUserHandler(userService service.UserService) *AddUserHandler {
	return &AddUserHandler{
		service: userService,
	}
}

func (auh *AddUserHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var user entities.User
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&user)
	if err != nil {
		appcontext.Logger.Error().Err(err).Msg("failed to decode the request body to add new user")
		utils.CommonResponse(writer, request, http.StatusBadRequest, domain.Error{constants.REQUEST_DECODING_FAILED, constants.DECODING_REQUEST_FAILED})
		return
	}
	appcontext.Logger.Info().Msgf("request received to add user %v", user)
	user, err = auh.service.Add(user)
	if err != nil {
		utils.CommonResponse(writer, request, http.StatusInternalServerError, domain.Error{constants.FAILED_CREATING_USER, err.Error()})
		return
	}
	utils.CommonResponse(writer, request, http.StatusOK, user)
}
