// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetJoinScriptsOKCode is the HTTP code returned for type GetJoinScriptsOK
const GetJoinScriptsOKCode int = 200

/*GetJoinScriptsOK Generic String answer

swagger:response getJoinScriptsOK
*/
type GetJoinScriptsOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetJoinScriptsOK creates GetJoinScriptsOK with default headers values
func NewGetJoinScriptsOK() *GetJoinScriptsOK {

	return &GetJoinScriptsOK{}
}

// WithPayload adds the payload to the get join scripts o k response
func (o *GetJoinScriptsOK) WithPayload(payload string) *GetJoinScriptsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get join scripts o k response
func (o *GetJoinScriptsOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetJoinScriptsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetJoinScriptsUnauthorizedCode is the HTTP code returned for type GetJoinScriptsUnauthorized
const GetJoinScriptsUnauthorizedCode int = 401

/*GetJoinScriptsUnauthorized 401 Not authorized

swagger:response getJoinScriptsUnauthorized
*/
type GetJoinScriptsUnauthorized struct {
}

// NewGetJoinScriptsUnauthorized creates GetJoinScriptsUnauthorized with default headers values
func NewGetJoinScriptsUnauthorized() *GetJoinScriptsUnauthorized {

	return &GetJoinScriptsUnauthorized{}
}

// WriteResponse to the client
func (o *GetJoinScriptsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GetJoinScriptsForbiddenCode is the HTTP code returned for type GetJoinScriptsForbidden
const GetJoinScriptsForbiddenCode int = 403

/*GetJoinScriptsForbidden 403 Forbidden

swagger:response getJoinScriptsForbidden
*/
type GetJoinScriptsForbidden struct {
}

// NewGetJoinScriptsForbidden creates GetJoinScriptsForbidden with default headers values
func NewGetJoinScriptsForbidden() *GetJoinScriptsForbidden {

	return &GetJoinScriptsForbidden{}
}

// WriteResponse to the client
func (o *GetJoinScriptsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// GetJoinScriptsInternalServerErrorCode is the HTTP code returned for type GetJoinScriptsInternalServerError
const GetJoinScriptsInternalServerErrorCode int = 500

/*GetJoinScriptsInternalServerError 500 Internal Server Error

swagger:response getJoinScriptsInternalServerError
*/
type GetJoinScriptsInternalServerError struct {
}

// NewGetJoinScriptsInternalServerError creates GetJoinScriptsInternalServerError with default headers values
func NewGetJoinScriptsInternalServerError() *GetJoinScriptsInternalServerError {

	return &GetJoinScriptsInternalServerError{}
}

// WriteResponse to the client
func (o *GetJoinScriptsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
