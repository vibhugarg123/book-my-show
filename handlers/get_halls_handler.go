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

type GetHallHandler struct {
	service service.HallService
}

func NewGetHallHandler(hallService service.HallService) *GetHallHandler {
	return &GetHallHandler{
		service: hallService,
	}
}

func (ghh *GetHallHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	theatreIdInString := vars["theatre-id"]
	theatreIdInInteger, err := strconv.Atoi(theatreIdInString)
	if err != nil {
		utils.CommonResponse(writer, request, http.StatusBadRequest, domain.Error{constants.NOT_VALID_INTEGER, err.Error()})
		return
	}
	err = utils.ValidateIntegerType(theatreIdInInteger)
	if err != nil {
		utils.CommonResponse(writer, request, http.StatusBadRequest, domain.Error{constants.NOT_VALID_INTEGER, err.Error()})
		return
	}
	appcontext.Logger.Info().Msg(fmt.Sprintf("halls to get for theatre-id %v", theatreIdInInteger))
	halls, err := ghh.service.GetHallByTheatreId(theatreIdInInteger)
	if err != nil {
		utils.CommonResponse(writer, request, http.StatusInternalServerError, domain.Error{constants.GET_HALLS_BY_THEATRE_ID_FAILED, err.Error()})
		return
	}
	utils.CommonResponse(writer, request, http.StatusOK, halls)
}
