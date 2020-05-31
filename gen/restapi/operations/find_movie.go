// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// FindMovieHandlerFunc turns a function with the right signature into a find movie handler
type FindMovieHandlerFunc func(FindMovieParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FindMovieHandlerFunc) Handle(params FindMovieParams) middleware.Responder {
	return fn(params)
}

// FindMovieHandler interface for that can handle valid find movie params
type FindMovieHandler interface {
	Handle(FindMovieParams) middleware.Responder
}

// NewFindMovie creates a new http.Handler for the find movie operation
func NewFindMovie(ctx *middleware.Context, handler FindMovieHandler) *FindMovie {
	return &FindMovie{Context: ctx, Handler: handler}
}

/*FindMovie swagger:route GET /movies/{id} findMovie

Search best result matching criteria

By passing in the appropriate options, you can search for
available movie in the system


*/
type FindMovie struct {
	Context *middleware.Context
	Handler FindMovieHandler
}

func (o *FindMovie) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewFindMovieParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}