// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetConfigDocByIDOKCode is the HTTP code returned for type GetConfigDocByIDOK
const GetConfigDocByIDOKCode int = 200

/*GetConfigDocByIDOK Generic String answer

swagger:response getConfigDocByIdOK
*/
type GetConfigDocByIDOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetConfigDocByIDOK creates GetConfigDocByIDOK with default headers values
func NewGetConfigDocByIDOK() *GetConfigDocByIDOK {

	return &GetConfigDocByIDOK{}
}

// WithPayload adds the payload to the get config doc by Id o k response
func (o *GetConfigDocByIDOK) WithPayload(payload string) *GetConfigDocByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config doc by Id o k response
func (o *GetConfigDocByIDOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigDocByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetConfigDocByIDUnauthorizedCode is the HTTP code returned for type GetConfigDocByIDUnauthorized
const GetConfigDocByIDUnauthorizedCode int = 401

/*GetConfigDocByIDUnauthorized 401 Not authorized

swagger:response getConfigDocByIdUnauthorized
*/
type GetConfigDocByIDUnauthorized struct {
}

// NewGetConfigDocByIDUnauthorized creates GetConfigDocByIDUnauthorized with default headers values
func NewGetConfigDocByIDUnauthorized() *GetConfigDocByIDUnauthorized {

	return &GetConfigDocByIDUnauthorized{}
}

// WriteResponse to the client
func (o *GetConfigDocByIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GetConfigDocByIDForbiddenCode is the HTTP code returned for type GetConfigDocByIDForbidden
const GetConfigDocByIDForbiddenCode int = 403

/*GetConfigDocByIDForbidden 403 Forbidden

swagger:response getConfigDocByIdForbidden
*/
type GetConfigDocByIDForbidden struct {
}

// NewGetConfigDocByIDForbidden creates GetConfigDocByIDForbidden with default headers values
func NewGetConfigDocByIDForbidden() *GetConfigDocByIDForbidden {

	return &GetConfigDocByIDForbidden{}
}

// WriteResponse to the client
func (o *GetConfigDocByIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// GetConfigDocByIDInternalServerErrorCode is the HTTP code returned for type GetConfigDocByIDInternalServerError
const GetConfigDocByIDInternalServerErrorCode int = 500

/*GetConfigDocByIDInternalServerError 500 Internal Server Error

swagger:response getConfigDocByIdInternalServerError
*/
type GetConfigDocByIDInternalServerError struct {
}

// NewGetConfigDocByIDInternalServerError creates GetConfigDocByIDInternalServerError with default headers values
func NewGetConfigDocByIDInternalServerError() *GetConfigDocByIDInternalServerError {

	return &GetConfigDocByIDInternalServerError{}
}

// WriteResponse to the client
func (o *GetConfigDocByIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}