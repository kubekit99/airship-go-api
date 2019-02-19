// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ValidateDesignHandlerFunc turns a function with the right signature into a validate design handler
type ValidateDesignHandlerFunc func(ValidateDesignParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ValidateDesignHandlerFunc) Handle(params ValidateDesignParams) middleware.Responder {
	return fn(params)
}

// ValidateDesignHandler interface for that can handle valid validate design params
type ValidateDesignHandler interface {
	Handle(ValidateDesignParams) middleware.Responder
}

// NewValidateDesign creates a new http.Handler for the validate design operation
func NewValidateDesign(ctx *middleware.Context, handler ValidateDesignHandler) *ValidateDesign {
	return &ValidateDesign{Context: ctx, Handler: handler}
}

/*ValidateDesign swagger:route POST /api/v1.0/validatedesign validateDesign

Validate documents

*/
type ValidateDesign struct {
	Context *middleware.Context
	Handler ValidateDesignHandler
}

func (o *ValidateDesign) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewValidateDesignParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}