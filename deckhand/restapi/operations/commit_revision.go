// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CommitRevisionHandlerFunc turns a function with the right signature into a commit revision handler
type CommitRevisionHandlerFunc func(CommitRevisionParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CommitRevisionHandlerFunc) Handle(params CommitRevisionParams) middleware.Responder {
	return fn(params)
}

// CommitRevisionHandler interface for that can handle valid commit revision params
type CommitRevisionHandler interface {
	Handle(CommitRevisionParams) middleware.Responder
}

// NewCommitRevision creates a new http.Handler for the commit revision operation
func NewCommitRevision(ctx *middleware.Context, handler CommitRevisionHandler) *CommitRevision {
	return &CommitRevision{Context: ctx, Handler: handler}
}

/*CommitRevision swagger:route PUT /api/v1.0/buckets/{bucket-name}/documents commitRevision

Create a batch of documents specified in the request body, whereby a new revision is created.

*/
type CommitRevision struct {
	Context *middleware.Context
	Handler CommitRevisionHandler
}

func (o *CommitRevision) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCommitRevisionParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}