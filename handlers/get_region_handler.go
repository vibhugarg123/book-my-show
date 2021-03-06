package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vibhugarg123/book-my-show/appcontext"
	"github.com/vibhugarg123/book-my-show/constants"
	"github.com/vibhugarg123/book-my-show/domain"
	"github.com/vibhugarg123/book-my-show/service"
	"github.com/vibhugarg123/book-my-show/utils"
	"net/http"
	"strconv"
)

type GetRegionHandler struct {
	service service.RegionService
}

func NewGetRegionHandler(regionService service.RegionService) *GetRegionHandler {
	return &GetRegionHandler{
		service: regionService,
	}
}

// swagger:route GET /region/{region-id} region noContent
// Get the regions with the respective region-id
// Parameters:
//  + name: region-id
//    type: string
//    in: path
//    required: true
// Responses:
//	200: regionsByIdResponse
//  404: errorResponse
//  500: errorResponse
func (grh *GetRegionHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	regionIdInString := vars["region-id"]
	regionIdInInteger, err := strconv.Atoi(regionIdInString)
	if err != nil {
		utils.CommonResponse(writer, request, http.StatusBadRequest, domain.Error{constants.NOT_VALID_INTEGER, err.Error()})
		return
	}
	err = utils.ValidateIntegerType(regionIdInInteger)
	if err != nil {
		utils.CommonResponse(writer, request, http.StatusBadRequest, domain.Error{constants.NOT_VALID_INTEGER, err.Error()})
		return
	}
	appcontext.Logger.Info().Msg(fmt.Sprintf("region to get for region-id %v", regionIdInInteger))
	region, err := grh.service.GetRegionById(regionIdInInteger)
	if utils.IsValidationError(err) {
		utils.CommonResponse(writer, request, http.StatusBadRequest, domain.Error{constants.GET_REGION_CALL_FAILED, err.Error()})
		return
	} else if err != nil {
		utils.CommonResponse(writer, request, http.StatusInternalServerError, domain.Error{constants.GET_REGION_CALL_FAILED, err.Error()})
		return
	}
	utils.CommonResponse(writer, request, http.StatusOK, region)
}
