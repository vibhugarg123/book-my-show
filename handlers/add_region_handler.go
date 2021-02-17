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

type AddRegionHandler struct {
	service service.RegionService
}

func NewAddRegionHandler(regionService service.RegionService) *AddRegionHandler {
	return &AddRegionHandler{
		service: regionService,
	}
}

// swagger:route POST /region region addRegionRequest
// Creates a new region with given id, name , region_type [1 for Country, 2 for State, 3 for District, 4 Town, 5 Village] and respective-parent id
// parameters: addRegionRequest
// Responses:
//	200: addRegionResponse
//  404: errorResponse
//  500: errorResponse

func (arh *AddRegionHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var region entities.Region
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&region)
	if err != nil {
		appcontext.Logger.Error().Err(err).Msg("failed to decode the request body to add new region")
		utils.CommonResponse(writer, request, http.StatusBadRequest, domain.Error{constants.REQUEST_DECODING_FAILED, constants.DECODING_REQUEST_FAILED})
		return
	}
	appcontext.Logger.Info().Msgf("request received to add new region %v", region)
	region, err = arh.service.Add(region)
	if utils.IsValidationError(err) {
		utils.CommonResponse(writer, request, http.StatusBadRequest, domain.Error{constants.REGION_CREATION_FAILED, err.Error()})
		return
	} else if err != nil {
		utils.CommonResponse(writer, request, http.StatusInternalServerError, domain.Error{constants.REGION_CREATION_FAILED, err.Error()})
		return
	}
	utils.CommonResponse(writer, request, http.StatusOK, region)
}
