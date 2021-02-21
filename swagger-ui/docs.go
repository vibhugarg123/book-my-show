// Package classification Documentation for Book My Show API
//
// Documentation for Book My Show API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//	Contact: Vibhu Garg<vibhu.garg@gojek.com>
//  Host: localhost:8089
//    description: Production server (uses live data)
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package swagger_ui

import (
	"github.com/vibhugarg123/book-my-show/domain"
	"github.com/vibhugarg123/book-my-show/swagger-ui/request"
	"github.com/vibhugarg123/book-my-show/swagger-ui/response"
)

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handlers

// Generic error message returned in form of json

// swagger:response errorResponse
type errorResponseWrapper struct {
	// in: body
	Body domain.Error
}

// swagger:parameters addUserRequest
type addUserRequestWrapper struct {
	// in: body
	Body request.AddUser
}

// swagger:response addUserResponse
type userResponseWrapper struct {
	// in: body
	Body response.AddUserResponse
}

// swagger:parameters addRegionRequest
type addRegionRequestWrapper struct {
	// in: body
	Body request.AddRegion
}

// swagger:response addRegionResponse
type addRegionResponseWrapper struct {
	// in: body
	Body response.AddRegionResponse
}

// swagger:parameters addTheatreRequest
type addTheatreRequestWrapper struct {
	// in: body
	Body request.AddTheatre
}

// swagger:response addTheatreResponse
type addTheatreResponseWrapper struct {
	// in: body
	Body response.AddTheatreResponse
}

// swagger:parameters addHallRequest
type addHallRequestWrapper struct {
	// in: body
	Body request.AddHall
}

// swagger:response addHallResponse
type addHallResponseWrapper struct {
	// in: body
	Body response.AddHallResponse
}

// swagger:parameters addMovieRequest
type addMovieRequestWrapper struct {
	// in: body
	Body request.AddMovie
}

// swagger:response addMovieResponse
type addMovieResponseWrapper struct {
	// in: body
	Body response.AddMovieResponse
}

// swagger:parameters addShowRequest
type addShowRequestWrapper struct {
	// in: body
	Body request.AddShow
}

// swagger:response addShowResponse
type addShowResponseWrapper struct {
	// in: body
	Body response.AddShowResponse
}

// swagger:parameters addBookingRequest
type addBookingRequestWrapper struct {
	// in: body
	Body request.AddBooking
}

// swagger:response addBookingResponse
type addBookingResponseWrapper struct {
	// in: body
	Body response.AddBookingResponse
}

// swagger:parameters loginRequest
type loginRequestWrapper struct {
	// in: body
	Body request.LoginRequest
}

// swagger:response loginResponse
type loginResponseWrapper struct {
	// in: body
	Body response.LoginResponse
}

// swagger:response noContent
type noContentRequest struct {
}

// swagger:response theatresByNameResponse
type theatresByNameResponseWrapper struct {
	// in: body
	Body []response.AddTheatreResponse
}

// swagger:response regionsByIdResponse
type regionsByIdResponseWrapper struct {
	// in: body
	Body response.AddRegionResponse
}
