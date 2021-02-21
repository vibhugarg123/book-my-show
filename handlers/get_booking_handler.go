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

type GetBookingHandler struct {
	service service.BookingService
}

func NewGetBookingHandler(bookingService service.BookingService) *GetBookingHandler {
	return &GetBookingHandler{
		service: bookingService,
	}
}

// swagger:route GET /booking/userid/{user-id} booking noContent
// Get the bookings of a particular user
// Parameters:
//  + name: user-id
//    type: string
//    in: path
//    required: true
// Responses:
//	200: bookingsByUserIdResponse
//  404: errorResponse
//  500: errorResponse
func (gbh *GetBookingHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userIdInString := vars["user-id"]
	userIdInInteger, err := strconv.Atoi(userIdInString)
	if err != nil {
		utils.CommonResponse(writer, request, http.StatusBadRequest, domain.Error{constants.NOT_VALID_INTEGER, err.Error()})
		return
	}
	err = utils.ValidateIntegerType(userIdInInteger)
	if err != nil {
		utils.CommonResponse(writer, request, http.StatusBadRequest, domain.Error{constants.NOT_VALID_INTEGER, err.Error()})
		return
	}
	appcontext.Logger.Info().Msg(fmt.Sprintf("booking to get for user-id %d", userIdInInteger))
	bookings, err := gbh.service.GetBooking(userIdInInteger)
	if utils.IsValidationError(err) {
		utils.CommonResponse(writer, request, http.StatusBadRequest, domain.Error{constants.GET_BOOKINGS_BY_USER_ID_FAILED, err.Error()})
		return
	} else if err != nil {
		utils.CommonResponse(writer, request, http.StatusInternalServerError, domain.Error{constants.GET_BOOKINGS_BY_USER_ID_FAILED, err.Error()})
		return
	}
	utils.CommonResponse(writer, request, http.StatusOK, bookings)
}
