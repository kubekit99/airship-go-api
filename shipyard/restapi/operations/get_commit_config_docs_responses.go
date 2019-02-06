// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetCommitConfigDocsOKCode is the HTTP code returned for type GetCommitConfigDocsOK
const GetCommitConfigDocsOKCode int = 200

/*GetCommitConfigDocsOK Generic String answer

swagger:response getCommitConfigDocsOK
*/
type GetCommitConfigDocsOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetCommitConfigDocsOK creates GetCommitConfigDocsOK with default headers values
func NewGetCommitConfigDocsOK() *GetCommitConfigDocsOK {

	return &GetCommitConfigDocsOK{}
}

// WithPayload adds the payload to the get commit config docs o k response
func (o *GetCommitConfigDocsOK) WithPayload(payload string) *GetCommitConfigDocsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get commit config docs o k response
func (o *GetCommitConfigDocsOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCommitConfigDocsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetCommitConfigDocsUnauthorizedCode is the HTTP code returned for type GetCommitConfigDocsUnauthorized
const GetCommitConfigDocsUnauthorizedCode int = 401

/*GetCommitConfigDocsUnauthorized 401 Not authorized

swagger:response getCommitConfigDocsUnauthorized
*/
type GetCommitConfigDocsUnauthorized struct {
}

// NewGetCommitConfigDocsUnauthorized creates GetCommitConfigDocsUnauthorized with default headers values
func NewGetCommitConfigDocsUnauthorized() *GetCommitConfigDocsUnauthorized {

	return &GetCommitConfigDocsUnauthorized{}
}

// WriteResponse to the client
func (o *GetCommitConfigDocsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GetCommitConfigDocsForbiddenCode is the HTTP code returned for type GetCommitConfigDocsForbidden
const GetCommitConfigDocsForbiddenCode int = 403

/*GetCommitConfigDocsForbidden 403 Forbidden

swagger:response getCommitConfigDocsForbidden
*/
type GetCommitConfigDocsForbidden struct {
}

// NewGetCommitConfigDocsForbidden creates GetCommitConfigDocsForbidden with default headers values
func NewGetCommitConfigDocsForbidden() *GetCommitConfigDocsForbidden {

	return &GetCommitConfigDocsForbidden{}
}

// WriteResponse to the client
func (o *GetCommitConfigDocsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// GetCommitConfigDocsInternalServerErrorCode is the HTTP code returned for type GetCommitConfigDocsInternalServerError
const GetCommitConfigDocsInternalServerErrorCode int = 500

/*GetCommitConfigDocsInternalServerError 500 Internal Server Error

swagger:response getCommitConfigDocsInternalServerError
*/
type GetCommitConfigDocsInternalServerError struct {
}

// NewGetCommitConfigDocsInternalServerError creates GetCommitConfigDocsInternalServerError with default headers values
func NewGetCommitConfigDocsInternalServerError() *GetCommitConfigDocsInternalServerError {

	return &GetCommitConfigDocsInternalServerError{}
}

// WriteResponse to the client
func (o *GetCommitConfigDocsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
