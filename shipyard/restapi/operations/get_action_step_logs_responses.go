// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetActionStepLogsOKCode is the HTTP code returned for type GetActionStepLogsOK
const GetActionStepLogsOKCode int = 200

/*GetActionStepLogsOK Generic String answer

swagger:response getActionStepLogsOK
*/
type GetActionStepLogsOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetActionStepLogsOK creates GetActionStepLogsOK with default headers values
func NewGetActionStepLogsOK() *GetActionStepLogsOK {

	return &GetActionStepLogsOK{}
}

// WithPayload adds the payload to the get action step logs o k response
func (o *GetActionStepLogsOK) WithPayload(payload string) *GetActionStepLogsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get action step logs o k response
func (o *GetActionStepLogsOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetActionStepLogsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetActionStepLogsUnauthorizedCode is the HTTP code returned for type GetActionStepLogsUnauthorized
const GetActionStepLogsUnauthorizedCode int = 401

/*GetActionStepLogsUnauthorized 401 Not authorized

swagger:response getActionStepLogsUnauthorized
*/
type GetActionStepLogsUnauthorized struct {
}

// NewGetActionStepLogsUnauthorized creates GetActionStepLogsUnauthorized with default headers values
func NewGetActionStepLogsUnauthorized() *GetActionStepLogsUnauthorized {

	return &GetActionStepLogsUnauthorized{}
}

// WriteResponse to the client
func (o *GetActionStepLogsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GetActionStepLogsForbiddenCode is the HTTP code returned for type GetActionStepLogsForbidden
const GetActionStepLogsForbiddenCode int = 403

/*GetActionStepLogsForbidden 403 Forbidden

swagger:response getActionStepLogsForbidden
*/
type GetActionStepLogsForbidden struct {
}

// NewGetActionStepLogsForbidden creates GetActionStepLogsForbidden with default headers values
func NewGetActionStepLogsForbidden() *GetActionStepLogsForbidden {

	return &GetActionStepLogsForbidden{}
}

// WriteResponse to the client
func (o *GetActionStepLogsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// GetActionStepLogsInternalServerErrorCode is the HTTP code returned for type GetActionStepLogsInternalServerError
const GetActionStepLogsInternalServerErrorCode int = 500

/*GetActionStepLogsInternalServerError 500 Internal Server Error

swagger:response getActionStepLogsInternalServerError
*/
type GetActionStepLogsInternalServerError struct {
}

// NewGetActionStepLogsInternalServerError creates GetActionStepLogsInternalServerError with default headers values
func NewGetActionStepLogsInternalServerError() *GetActionStepLogsInternalServerError {

	return &GetActionStepLogsInternalServerError{}
}

// WriteResponse to the client
func (o *GetActionStepLogsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
