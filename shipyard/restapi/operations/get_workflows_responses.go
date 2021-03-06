// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetWorkflowsOKCode is the HTTP code returned for type GetWorkflowsOK
const GetWorkflowsOKCode int = 200

/*GetWorkflowsOK Generic String answer

swagger:response getWorkflowsOK
*/
type GetWorkflowsOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetWorkflowsOK creates GetWorkflowsOK with default headers values
func NewGetWorkflowsOK() *GetWorkflowsOK {

	return &GetWorkflowsOK{}
}

// WithPayload adds the payload to the get workflows o k response
func (o *GetWorkflowsOK) WithPayload(payload string) *GetWorkflowsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get workflows o k response
func (o *GetWorkflowsOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetWorkflowsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetWorkflowsUnauthorizedCode is the HTTP code returned for type GetWorkflowsUnauthorized
const GetWorkflowsUnauthorizedCode int = 401

/*GetWorkflowsUnauthorized 401 Not authorized

swagger:response getWorkflowsUnauthorized
*/
type GetWorkflowsUnauthorized struct {
}

// NewGetWorkflowsUnauthorized creates GetWorkflowsUnauthorized with default headers values
func NewGetWorkflowsUnauthorized() *GetWorkflowsUnauthorized {

	return &GetWorkflowsUnauthorized{}
}

// WriteResponse to the client
func (o *GetWorkflowsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GetWorkflowsForbiddenCode is the HTTP code returned for type GetWorkflowsForbidden
const GetWorkflowsForbiddenCode int = 403

/*GetWorkflowsForbidden 403 Forbidden

swagger:response getWorkflowsForbidden
*/
type GetWorkflowsForbidden struct {
}

// NewGetWorkflowsForbidden creates GetWorkflowsForbidden with default headers values
func NewGetWorkflowsForbidden() *GetWorkflowsForbidden {

	return &GetWorkflowsForbidden{}
}

// WriteResponse to the client
func (o *GetWorkflowsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// GetWorkflowsInternalServerErrorCode is the HTTP code returned for type GetWorkflowsInternalServerError
const GetWorkflowsInternalServerErrorCode int = 500

/*GetWorkflowsInternalServerError 500 Internal Server Error

swagger:response getWorkflowsInternalServerError
*/
type GetWorkflowsInternalServerError struct {
}

// NewGetWorkflowsInternalServerError creates GetWorkflowsInternalServerError with default headers values
func NewGetWorkflowsInternalServerError() *GetWorkflowsInternalServerError {

	return &GetWorkflowsInternalServerError{}
}

// WriteResponse to the client
func (o *GetWorkflowsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
