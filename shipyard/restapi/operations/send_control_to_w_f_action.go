// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// SendControlToWFActionHandlerFunc turns a function with the right signature into a send control to w f action handler
type SendControlToWFActionHandlerFunc func(SendControlToWFActionParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SendControlToWFActionHandlerFunc) Handle(params SendControlToWFActionParams) middleware.Responder {
	return fn(params)
}

// SendControlToWFActionHandler interface for that can handle valid send control to w f action params
type SendControlToWFActionHandler interface {
	Handle(SendControlToWFActionParams) middleware.Responder
}

// NewSendControlToWFAction creates a new http.Handler for the send control to w f action operation
func NewSendControlToWFAction(ctx *middleware.Context, handler SendControlToWFActionHandler) *SendControlToWFAction {
	return &SendControlToWFAction{Context: ctx, Handler: handler}
}

/*SendControlToWFAction swagger:route POST /api/v1.0/actions/{action-id}/control/{control-verb} sendControlToWFAction

Send a control to an action

*/
type SendControlToWFAction struct {
	Context *middleware.Context
	Handler SendControlToWFActionHandler
}

func (o *SendControlToWFAction) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSendControlToWFActionParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}