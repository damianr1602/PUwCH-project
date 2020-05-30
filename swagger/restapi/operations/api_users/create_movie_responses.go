// Code generated by go-swagger; DO NOT EDIT.

package api_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/damianr1602/chmuryrest/swagger/models"
)

// CreateMovieCreatedCode is the HTTP code returned for type CreateMovieCreated
const CreateMovieCreatedCode int = 201

/*CreateMovieCreated item created

swagger:response createMovieCreated
*/
type CreateMovieCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Movie `json:"body,omitempty"`
}

// NewCreateMovieCreated creates CreateMovieCreated with default headers values
func NewCreateMovieCreated() *CreateMovieCreated {

	return &CreateMovieCreated{}
}

// WithPayload adds the payload to the create movie created response
func (o *CreateMovieCreated) WithPayload(payload *models.Movie) *CreateMovieCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create movie created response
func (o *CreateMovieCreated) SetPayload(payload *models.Movie) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateMovieCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateMovieUnauthorizedCode is the HTTP code returned for type CreateMovieUnauthorized
const CreateMovieUnauthorizedCode int = 401

/*CreateMovieUnauthorized Unauthorized

swagger:response createMovieUnauthorized
*/
type CreateMovieUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Result `json:"body,omitempty"`
}

// NewCreateMovieUnauthorized creates CreateMovieUnauthorized with default headers values
func NewCreateMovieUnauthorized() *CreateMovieUnauthorized {

	return &CreateMovieUnauthorized{}
}

// WithPayload adds the payload to the create movie unauthorized response
func (o *CreateMovieUnauthorized) WithPayload(payload *models.Result) *CreateMovieUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create movie unauthorized response
func (o *CreateMovieUnauthorized) SetPayload(payload *models.Result) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateMovieUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateMovieConflictCode is the HTTP code returned for type CreateMovieConflict
const CreateMovieConflictCode int = 409

/*CreateMovieConflict an existing item already exists

swagger:response createMovieConflict
*/
type CreateMovieConflict struct {
}

// NewCreateMovieConflict creates CreateMovieConflict with default headers values
func NewCreateMovieConflict() *CreateMovieConflict {

	return &CreateMovieConflict{}
}

// WriteResponse to the client
func (o *CreateMovieConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}
