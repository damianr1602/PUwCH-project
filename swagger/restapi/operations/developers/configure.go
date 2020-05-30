// Code generated by go-swagger; DO NOT EDIT.

package developers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ConfigureHandlerFunc turns a function with the right signature into a configure handler
type ConfigureHandlerFunc func(ConfigureParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ConfigureHandlerFunc) Handle(params ConfigureParams) middleware.Responder {
	return fn(params)
}

// ConfigureHandler interface for that can handle valid configure params
type ConfigureHandler interface {
	Handle(ConfigureParams) middleware.Responder
}

// NewConfigure creates a new http.Handler for the configure operation
func NewConfigure(ctx *middleware.Context, handler ConfigureHandler) *Configure {
	return &Configure{Context: ctx, Handler: handler}
}

/*Configure swagger:route GET /configure developers configure

configure device

By passing in the appropriate options, you can configure
available devices


*/
type Configure struct {
	Context *middleware.Context
	Handler ConfigureHandler
}

func (o *Configure) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewConfigureParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}