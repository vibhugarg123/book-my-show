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

type AddTheatreHandler struct {
	service service.TheatreService
}

func NewAddTheatreHandler(theatreService service.TheatreService) *AddTheatreHandler {
	return &AddTheatreHandler{
		service: theatreService,
	}
}

func (ath *AddTheatreHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var theatre entities.Theatre
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&theatre)
	if err != nil {
		appcontext.Logger.Error().Err(err).Msg("failed to decode the request body to add new region")
		utils.CommonResponse(writer, request, http.StatusBadRequest, domain.Error{constants.REQUEST_DECODING_FAILED, constants.DECODING_REQUEST_FAILED})
		return
	}
	appcontext.Logger.Info().Msgf("request received to add new theatre %v", theatre)
	theatre, err = ath.service.Add(theatre)
	if err != nil {
		utils.CommonResponse(writer, request, http.StatusInternalServerError, domain.Error{constants.THEATRE_CREATION_FAILED, err.Error()})
		return
	}
	utils.CommonResponse(writer, request, http.StatusOK, theatre)
}
