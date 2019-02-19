// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetWFActionByIDHandlerFunc turns a function with the right signature into a get w f action by Id handler
type GetWFActionByIDHandlerFunc func(GetWFActionByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetWFActionByIDHandlerFunc) Handle(params GetWFActionByIDParams) middleware.Responder {
	return fn(params)
}

// GetWFActionByIDHandler interface for that can handle valid get w f action by Id params
type GetWFActionByIDHandler interface {
	Handle(GetWFActionByIDParams) middleware.Responder
}

// NewGetWFActionByID creates a new http.Handler for the get w f action by Id operation
func NewGetWFActionByID(ctx *middleware.Context, handler GetWFActionByIDHandler) *GetWFActionByID {
	return &GetWFActionByID{Context: ctx, Handler: handler}
}

/*GetWFActionByID swagger:route GET /api/v1.0/actions/{action-id} getWFActionById

Retrieve an action by its id

*/
type GetWFActionByID struct {
	Context *middleware.Context
	Handler GetWFActionByIDHandler
}

func (o *GetWFActionByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetWFActionByIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
