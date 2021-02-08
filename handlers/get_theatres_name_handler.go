package handlers

import (
	"github.com/gorilla/mux"
	"github.com/vibhugarg123/book-my-show/appcontext"
	"github.com/vibhugarg123/book-my-show/constants"
	"github.com/vibhugarg123/book-my-show/domain"
	"github.com/vibhugarg123/book-my-show/service"
	"github.com/vibhugarg123/book-my-show/utils"
	"net/http"
)

type GetTheatreByNameHandler struct {
	service service.TheatreService
}

func NewGetTheatreByNameHandler(theatreService service.TheatreService) *GetTheatreByNameHandler {
	return &GetTheatreByNameHandler{
		service: theatreService,
	}
}

func (ngt *GetTheatreByNameHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	theatreName := vars["theatre-name"]
	appcontext.Logger.Info().Msgf("request received for fetching theatres with name %s", theatreName)
	theatres, err := ngt.service.GetTheatreByName(theatreName)
	if err != nil {
		utils.CommonResponse(writer, request, http.StatusInternalServerError, domain.Error{constants.GET_THEATRES_CALL_FAILED, err.Error()})
		return
	}
	utils.CommonResponse(writer, request, http.StatusOK, theatres)
}
