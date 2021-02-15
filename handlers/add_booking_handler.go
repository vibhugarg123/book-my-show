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

type AddBookingHandler struct {
	service service.BookingService
}

func NewAddBookingHandler(bookingService service.BookingService) *AddBookingHandler {
	return &AddBookingHandler{
		service: bookingService,
	}
}

func (abh *AddBookingHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var booking entities.Booking
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&booking)
	if err != nil {
		appcontext.Logger.Error().Err(err).Msg("failed to decode the request body to create new booking")
		utils.CommonResponse(writer, request, http.StatusBadRequest, domain.Error{constants.REQUEST_DECODING_FAILED, constants.DECODING_REQUEST_FAILED})
		return
	}
	appcontext.Logger.Info().Msgf("request received to create new booking %v", booking)
	booking, err = abh.service.Add(booking)
	if utils.IsValidationError(err) {
		utils.CommonResponse(writer, request, http.StatusBadRequest, domain.Error{constants.BOOKING_CREATION_FAILED, err.Error()})
		return
	} else if err != nil {
		utils.CommonResponse(writer, request, http.StatusInternalServerError, domain.Error{constants.BOOKING_CREATION_FAILED, err.Error()})
		return
	}
	utils.CommonResponse(writer, request, http.StatusOK, booking)
}