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
	"github.com/vibhugarg123/book-my-show/entities"
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

// swagger:response userResponse
type userResponseWrapper struct {
	// in: body
	Body entities.User
}

//swagger:model
type addUser struct {
	// required: true
	FirstName string `json:"first_name"`
	// required: true
	LastName string `json:"last_name"`
	// required: true
	EmailId string `json:"email_id"`
	// required: true
	Password string `json:"password"`
}

// swagger:parameters addUser
type addUserResponseWrapper struct {
	// in: body
	Body addUser
}
