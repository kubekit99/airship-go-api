// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetActionsOKCode is the HTTP code returned for type GetActionsOK
const GetActionsOKCode int = 200

/*GetActionsOK Generic String answer

swagger:response getActionsOK
*/
type GetActionsOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetActionsOK creates GetActionsOK with default headers values
func NewGetActionsOK() *GetActionsOK {

	return &GetActionsOK{}
}

// WithPayload adds the payload to the get actions o k response
func (o *GetActionsOK) WithPayload(payload string) *GetActionsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get actions o k response
func (o *GetActionsOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetActionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetActionsUnauthorizedCode is the HTTP code returned for type GetActionsUnauthorized
const GetActionsUnauthorizedCode int = 401

/*GetActionsUnauthorized 401 Not authorized

swagger:response getActionsUnauthorized
*/
type GetActionsUnauthorized struct {
}

// NewGetActionsUnauthorized creates GetActionsUnauthorized with default headers values
func NewGetActionsUnauthorized() *GetActionsUnauthorized {

	return &GetActionsUnauthorized{}
}

// WriteResponse to the client
func (o *GetActionsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GetActionsForbiddenCode is the HTTP code returned for type GetActionsForbidden
const GetActionsForbiddenCode int = 403

/*GetActionsForbidden 403 Forbidden

swagger:response getActionsForbidden
*/
type GetActionsForbidden struct {
}

// NewGetActionsForbidden creates GetActionsForbidden with default headers values
func NewGetActionsForbidden() *GetActionsForbidden {

	return &GetActionsForbidden{}
}

// WriteResponse to the client
func (o *GetActionsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// GetActionsInternalServerErrorCode is the HTTP code returned for type GetActionsInternalServerError
const GetActionsInternalServerErrorCode int = 500

/*GetActionsInternalServerError 500 Internal Server Error

swagger:response getActionsInternalServerError
*/
type GetActionsInternalServerError struct {
}

// NewGetActionsInternalServerError creates GetActionsInternalServerError with default headers values
func NewGetActionsInternalServerError() *GetActionsInternalServerError {

	return &GetActionsInternalServerError{}
}

// WriteResponse to the client
func (o *GetActionsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
