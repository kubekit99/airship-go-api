// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetActionStepByIDOKCode is the HTTP code returned for type GetActionStepByIDOK
const GetActionStepByIDOKCode int = 200

/*GetActionStepByIDOK Generic String answer

swagger:response getActionStepByIdOK
*/
type GetActionStepByIDOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetActionStepByIDOK creates GetActionStepByIDOK with default headers values
func NewGetActionStepByIDOK() *GetActionStepByIDOK {

	return &GetActionStepByIDOK{}
}

// WithPayload adds the payload to the get action step by Id o k response
func (o *GetActionStepByIDOK) WithPayload(payload string) *GetActionStepByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get action step by Id o k response
func (o *GetActionStepByIDOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetActionStepByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetActionStepByIDUnauthorizedCode is the HTTP code returned for type GetActionStepByIDUnauthorized
const GetActionStepByIDUnauthorizedCode int = 401

/*GetActionStepByIDUnauthorized 401 Not authorized

swagger:response getActionStepByIdUnauthorized
*/
type GetActionStepByIDUnauthorized struct {
}

// NewGetActionStepByIDUnauthorized creates GetActionStepByIDUnauthorized with default headers values
func NewGetActionStepByIDUnauthorized() *GetActionStepByIDUnauthorized {

	return &GetActionStepByIDUnauthorized{}
}

// WriteResponse to the client
func (o *GetActionStepByIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GetActionStepByIDForbiddenCode is the HTTP code returned for type GetActionStepByIDForbidden
const GetActionStepByIDForbiddenCode int = 403

/*GetActionStepByIDForbidden 403 Forbidden

swagger:response getActionStepByIdForbidden
*/
type GetActionStepByIDForbidden struct {
}

// NewGetActionStepByIDForbidden creates GetActionStepByIDForbidden with default headers values
func NewGetActionStepByIDForbidden() *GetActionStepByIDForbidden {

	return &GetActionStepByIDForbidden{}
}

// WriteResponse to the client
func (o *GetActionStepByIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// GetActionStepByIDInternalServerErrorCode is the HTTP code returned for type GetActionStepByIDInternalServerError
const GetActionStepByIDInternalServerErrorCode int = 500

/*GetActionStepByIDInternalServerError 500 Internal Server Error

swagger:response getActionStepByIdInternalServerError
*/
type GetActionStepByIDInternalServerError struct {
}

// NewGetActionStepByIDInternalServerError creates GetActionStepByIDInternalServerError with default headers values
func NewGetActionStepByIDInternalServerError() *GetActionStepByIDInternalServerError {

	return &GetActionStepByIDInternalServerError{}
}

// WriteResponse to the client
func (o *GetActionStepByIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
