// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetWFActionStepByIDOKCode is the HTTP code returned for type GetWFActionStepByIDOK
const GetWFActionStepByIDOKCode int = 200

/*GetWFActionStepByIDOK Generic String answer

swagger:response getWFActionStepByIdOK
*/
type GetWFActionStepByIDOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetWFActionStepByIDOK creates GetWFActionStepByIDOK with default headers values
func NewGetWFActionStepByIDOK() *GetWFActionStepByIDOK {

	return &GetWFActionStepByIDOK{}
}

// WithPayload adds the payload to the get w f action step by Id o k response
func (o *GetWFActionStepByIDOK) WithPayload(payload string) *GetWFActionStepByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get w f action step by Id o k response
func (o *GetWFActionStepByIDOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetWFActionStepByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetWFActionStepByIDUnauthorizedCode is the HTTP code returned for type GetWFActionStepByIDUnauthorized
const GetWFActionStepByIDUnauthorizedCode int = 401

/*GetWFActionStepByIDUnauthorized 401 Not authorized

swagger:response getWFActionStepByIdUnauthorized
*/
type GetWFActionStepByIDUnauthorized struct {
}

// NewGetWFActionStepByIDUnauthorized creates GetWFActionStepByIDUnauthorized with default headers values
func NewGetWFActionStepByIDUnauthorized() *GetWFActionStepByIDUnauthorized {

	return &GetWFActionStepByIDUnauthorized{}
}

// WriteResponse to the client
func (o *GetWFActionStepByIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GetWFActionStepByIDForbiddenCode is the HTTP code returned for type GetWFActionStepByIDForbidden
const GetWFActionStepByIDForbiddenCode int = 403

/*GetWFActionStepByIDForbidden 403 Forbidden

swagger:response getWFActionStepByIdForbidden
*/
type GetWFActionStepByIDForbidden struct {
}

// NewGetWFActionStepByIDForbidden creates GetWFActionStepByIDForbidden with default headers values
func NewGetWFActionStepByIDForbidden() *GetWFActionStepByIDForbidden {

	return &GetWFActionStepByIDForbidden{}
}

// WriteResponse to the client
func (o *GetWFActionStepByIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// GetWFActionStepByIDInternalServerErrorCode is the HTTP code returned for type GetWFActionStepByIDInternalServerError
const GetWFActionStepByIDInternalServerErrorCode int = 500

/*GetWFActionStepByIDInternalServerError 500 Internal Server Error

swagger:response getWFActionStepByIdInternalServerError
*/
type GetWFActionStepByIDInternalServerError struct {
}

// NewGetWFActionStepByIDInternalServerError creates GetWFActionStepByIDInternalServerError with default headers values
func NewGetWFActionStepByIDInternalServerError() *GetWFActionStepByIDInternalServerError {

	return &GetWFActionStepByIDInternalServerError{}
}

// WriteResponse to the client
func (o *GetWFActionStepByIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
