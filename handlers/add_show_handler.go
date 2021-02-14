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

type AddShowHandler struct {
	service service.ShowService
}

func NewAddShowHandler(showService service.ShowService) *AddShowHandler {
	return &AddShowHandler{
		service: showService,
	}
}

func (ash *AddShowHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var show entities.Show
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&show)
	if err != nil {
		appcontext.Logger.Error().Err(err).Msg("failed to decode the request body to add new show")
		utils.CommonResponse(writer, request, http.StatusBadRequest, domain.Error{constants.REQUEST_DECODING_FAILED, constants.DECODING_REQUEST_FAILED})
		return
	}
	appcontext.Logger.Info().Msgf("request received to add new show %v", show)
	showReturned, err := ash.service.Add(show)
	if err != nil {
		utils.CommonResponse(writer, request, http.StatusInternalServerError, domain.Error{constants.SHOW_CREATION_FAILED, err.Error()})
		return
	}
	utils.CommonResponse(writer, request, http.StatusOK, showReturned)
}