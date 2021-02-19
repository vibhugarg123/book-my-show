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

type AddHallHandler struct {
	service service.HallService
}

func NewAddHallHandler(hallService service.HallService) *AddHallHandler {
	return &AddHallHandler{
		service: hallService,
	}
}

// swagger:route POST /hall hall addHallRequest
// Adds a new hall to a theatre
// parameters: addHallRequest
// Responses:
//	200: addHallResponse
//  404: errorResponse
//  500: errorResponse
func (ahh *AddHallHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var hall entities.Hall
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&hall)
	if err != nil {
		appcontext.Logger.Error().Err(err).Msg("failed to decode the request body to add new hall")
		utils.CommonResponse(writer, request, http.StatusBadRequest, domain.Error{constants.REQUEST_DECODING_FAILED, constants.DECODING_REQUEST_FAILED})
		return
	}
	appcontext.Logger.Info().Msgf("request received to add new hall %v", hall)
	hall, err = ahh.service.Add(hall)
	if utils.IsValidationError(err) {
		utils.CommonResponse(writer, request, http.StatusBadRequest, domain.Error{constants.HALL_CREATION_FAILED, err.Error()})
		return
	} else if err != nil {
		utils.CommonResponse(writer, request, http.StatusInternalServerError, domain.Error{constants.HALL_CREATION_FAILED, err.Error()})
		return
	}
	utils.CommonResponse(writer, request, http.StatusOK, hall)
}
