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

type LoginHandler struct {
	service service.UserService
}

func NewLoginHandler(userService service.UserService) *LoginHandler {
	return &LoginHandler{
		service: userService,
	}
}

func (lh *LoginHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var user entities.User
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&user)
	if err != nil {
		appcontext.Logger.Error().Err(err).Msg("failed to decode the request body to add new configuration mapping")
		utils.CommonResponse(writer, request, http.StatusBadRequest, domain.Error{constants.REQUEST_DECODING_FAILED, constants.DECODING_REQUEST_FAILED})
		return
	}
	appcontext.Logger.Info().Msgf("request received for login %v", user)
	user, err = lh.service.Login(user)
	if err != nil {
		utils.CommonResponse(writer, request, http.StatusInternalServerError, domain.Error{constants.LOGIN_FAILED, err.Error()})
		return
	}
	m := make(map[string]string)
	m["login_status"] = "login successful"
	utils.CommonResponse(writer, request, http.StatusOK, m)
}
