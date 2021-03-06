// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/damianr1602/chmuryrest/gen/models"
)

// AllMoviesOKCode is the HTTP code returned for type AllMoviesOK
const AllMoviesOKCode int = 200

/*AllMoviesOK search results matching criteria

swagger:response allMoviesOK
*/
type AllMoviesOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Movie `json:"body,omitempty"`
}

// NewAllMoviesOK creates AllMoviesOK with default headers values
func NewAllMoviesOK() *AllMoviesOK {

	return &AllMoviesOK{}
}

// WithPayload adds the payload to the all movies o k response
func (o *AllMoviesOK) WithPayload(payload []*models.Movie) *AllMoviesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the all movies o k response
func (o *AllMoviesOK) SetPayload(payload []*models.Movie) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AllMoviesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Movie, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// AllMoviesNotFoundCode is the HTTP code returned for type AllMoviesNotFound
const AllMoviesNotFoundCode int = 404

/*AllMoviesNotFound The specified resource was not found

swagger:response allMoviesNotFound
*/
type AllMoviesNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Result `json:"body,omitempty"`
}

// NewAllMoviesNotFound creates AllMoviesNotFound with default headers values
func NewAllMoviesNotFound() *AllMoviesNotFound {

	return &AllMoviesNotFound{}
}

// WithPayload adds the payload to the all movies not found response
func (o *AllMoviesNotFound) WithPayload(payload *models.Result) *AllMoviesNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the all movies not found response
func (o *AllMoviesNotFound) SetPayload(payload *models.Result) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AllMoviesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
