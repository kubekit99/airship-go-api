// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// LoadDesignDataPartsOKCode is the HTTP code returned for type LoadDesignDataPartsOK
const LoadDesignDataPartsOKCode int = 200

/*LoadDesignDataPartsOK Generic String answer

swagger:response loadDesignDataPartsOK
*/
type LoadDesignDataPartsOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewLoadDesignDataPartsOK creates LoadDesignDataPartsOK with default headers values
func NewLoadDesignDataPartsOK() *LoadDesignDataPartsOK {

	return &LoadDesignDataPartsOK{}
}

// WithPayload adds the payload to the load design data parts o k response
func (o *LoadDesignDataPartsOK) WithPayload(payload string) *LoadDesignDataPartsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the load design data parts o k response
func (o *LoadDesignDataPartsOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoadDesignDataPartsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// LoadDesignDataPartsUnauthorizedCode is the HTTP code returned for type LoadDesignDataPartsUnauthorized
const LoadDesignDataPartsUnauthorizedCode int = 401

/*LoadDesignDataPartsUnauthorized 401 Not authorized

swagger:response loadDesignDataPartsUnauthorized
*/
type LoadDesignDataPartsUnauthorized struct {
}

// NewLoadDesignDataPartsUnauthorized creates LoadDesignDataPartsUnauthorized with default headers values
func NewLoadDesignDataPartsUnauthorized() *LoadDesignDataPartsUnauthorized {

	return &LoadDesignDataPartsUnauthorized{}
}

// WriteResponse to the client
func (o *LoadDesignDataPartsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// LoadDesignDataPartsForbiddenCode is the HTTP code returned for type LoadDesignDataPartsForbidden
const LoadDesignDataPartsForbiddenCode int = 403

/*LoadDesignDataPartsForbidden 403 Forbidden

swagger:response loadDesignDataPartsForbidden
*/
type LoadDesignDataPartsForbidden struct {
}

// NewLoadDesignDataPartsForbidden creates LoadDesignDataPartsForbidden with default headers values
func NewLoadDesignDataPartsForbidden() *LoadDesignDataPartsForbidden {

	return &LoadDesignDataPartsForbidden{}
}

// WriteResponse to the client
func (o *LoadDesignDataPartsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// LoadDesignDataPartsInternalServerErrorCode is the HTTP code returned for type LoadDesignDataPartsInternalServerError
const LoadDesignDataPartsInternalServerErrorCode int = 500

/*LoadDesignDataPartsInternalServerError 500 Internal Server Error

swagger:response loadDesignDataPartsInternalServerError
*/
type LoadDesignDataPartsInternalServerError struct {
}

// NewLoadDesignDataPartsInternalServerError creates LoadDesignDataPartsInternalServerError with default headers values
func NewLoadDesignDataPartsInternalServerError() *LoadDesignDataPartsInternalServerError {

	return &LoadDesignDataPartsInternalServerError{}
}

// WriteResponse to the client
func (o *LoadDesignDataPartsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
