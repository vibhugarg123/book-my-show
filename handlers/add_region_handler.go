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
	if err != nil {
		utils.CommonResponse(writer, request, http.StatusInternalServerError, domain.Error{constants.REGION_CREATION_FAILED, err.Error()})
		return
	}
	utils.CommonResponse(writer, request, http.StatusOK, region)
}
