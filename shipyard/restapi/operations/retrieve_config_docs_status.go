// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// RetrieveConfigDocsStatusHandlerFunc turns a function with the right signature into a retrieve config docs status handler
type RetrieveConfigDocsStatusHandlerFunc func(RetrieveConfigDocsStatusParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RetrieveConfigDocsStatusHandlerFunc) Handle(params RetrieveConfigDocsStatusParams) middleware.Responder {
	return fn(params)
}

// RetrieveConfigDocsStatusHandler interface for that can handle valid retrieve config docs status params
type RetrieveConfigDocsStatusHandler interface {
	Handle(RetrieveConfigDocsStatusParams) middleware.Responder
}

// NewRetrieveConfigDocsStatus creates a new http.Handler for the retrieve config docs status operation
func NewRetrieveConfigDocsStatus(ctx *middleware.Context, handler RetrieveConfigDocsStatusHandler) *RetrieveConfigDocsStatus {
	return &RetrieveConfigDocsStatus{Context: ctx, Handler: handler}
}

/*RetrieveConfigDocsStatus swagger:route GET /api/v1.0/configdocs retrieveConfigDocsStatus

Retrieve the status of the configdocs

*/
type RetrieveConfigDocsStatus struct {
	Context *middleware.Context
	Handler RetrieveConfigDocsStatusHandler
}

func (o *RetrieveConfigDocsStatus) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRetrieveConfigDocsStatusParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
