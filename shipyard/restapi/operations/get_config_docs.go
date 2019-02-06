// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetConfigDocsHandlerFunc turns a function with the right signature into a get config docs handler
type GetConfigDocsHandlerFunc func(GetConfigDocsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetConfigDocsHandlerFunc) Handle(params GetConfigDocsParams) middleware.Responder {
	return fn(params)
}

// GetConfigDocsHandler interface for that can handle valid get config docs params
type GetConfigDocsHandler interface {
	Handle(GetConfigDocsParams) middleware.Responder
}

// NewGetConfigDocs creates a new http.Handler for the get config docs operation
func NewGetConfigDocs(ctx *middleware.Context, handler GetConfigDocsHandler) *GetConfigDocs {
	return &GetConfigDocs{Context: ctx, Handler: handler}
}

/*GetConfigDocs swagger:route GET /api/v1.0/configdocs getConfigDocs

Returns the '/configdocs'

*/
type GetConfigDocs struct {
	Context *middleware.Context
	Handler GetConfigDocsHandler
}

func (o *GetConfigDocs) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetConfigDocsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}