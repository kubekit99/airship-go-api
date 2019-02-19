// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// LoadDesignDataPartsHandlerFunc turns a function with the right signature into a load design data parts handler
type LoadDesignDataPartsHandlerFunc func(LoadDesignDataPartsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn LoadDesignDataPartsHandlerFunc) Handle(params LoadDesignDataPartsParams) middleware.Responder {
	return fn(params)
}

// LoadDesignDataPartsHandler interface for that can handle valid load design data parts params
type LoadDesignDataPartsHandler interface {
	Handle(LoadDesignDataPartsParams) middleware.Responder
}

// NewLoadDesignDataParts creates a new http.Handler for the load design data parts operation
func NewLoadDesignDataParts(ctx *middleware.Context, handler LoadDesignDataPartsHandler) *LoadDesignDataParts {
	return &LoadDesignDataParts{Context: ctx, Handler: handler}
}

/*LoadDesignDataParts swagger:route POST /api/v1.0/designs/{design-id}/parts loadDesignDataParts

Load design data

*/
type LoadDesignDataParts struct {
	Context *middleware.Context
	Handler LoadDesignDataPartsHandler
}

func (o *LoadDesignDataParts) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewLoadDesignDataPartsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}