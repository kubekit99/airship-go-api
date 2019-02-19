// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ValidateSiteDesignOKCode is the HTTP code returned for type ValidateSiteDesignOK
const ValidateSiteDesignOKCode int = 200

/*ValidateSiteDesignOK Generic String answer

swagger:response validateSiteDesignOK
*/
type ValidateSiteDesignOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewValidateSiteDesignOK creates ValidateSiteDesignOK with default headers values
func NewValidateSiteDesignOK() *ValidateSiteDesignOK {

	return &ValidateSiteDesignOK{}
}

// WithPayload adds the payload to the validate site design o k response
func (o *ValidateSiteDesignOK) WithPayload(payload string) *ValidateSiteDesignOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the validate site design o k response
func (o *ValidateSiteDesignOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ValidateSiteDesignOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ValidateSiteDesignUnauthorizedCode is the HTTP code returned for type ValidateSiteDesignUnauthorized
const ValidateSiteDesignUnauthorizedCode int = 401

/*ValidateSiteDesignUnauthorized 401 Not authorized

swagger:response validateSiteDesignUnauthorized
*/
type ValidateSiteDesignUnauthorized struct {
}

// NewValidateSiteDesignUnauthorized creates ValidateSiteDesignUnauthorized with default headers values
func NewValidateSiteDesignUnauthorized() *ValidateSiteDesignUnauthorized {

	return &ValidateSiteDesignUnauthorized{}
}

// WriteResponse to the client
func (o *ValidateSiteDesignUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ValidateSiteDesignForbiddenCode is the HTTP code returned for type ValidateSiteDesignForbidden
const ValidateSiteDesignForbiddenCode int = 403

/*ValidateSiteDesignForbidden 403 Forbidden

swagger:response validateSiteDesignForbidden
*/
type ValidateSiteDesignForbidden struct {
}

// NewValidateSiteDesignForbidden creates ValidateSiteDesignForbidden with default headers values
func NewValidateSiteDesignForbidden() *ValidateSiteDesignForbidden {

	return &ValidateSiteDesignForbidden{}
}

// WriteResponse to the client
func (o *ValidateSiteDesignForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// ValidateSiteDesignInternalServerErrorCode is the HTTP code returned for type ValidateSiteDesignInternalServerError
const ValidateSiteDesignInternalServerErrorCode int = 500

/*ValidateSiteDesignInternalServerError 500 Internal Server Error

swagger:response validateSiteDesignInternalServerError
*/
type ValidateSiteDesignInternalServerError struct {
}

// NewValidateSiteDesignInternalServerError creates ValidateSiteDesignInternalServerError with default headers values
func NewValidateSiteDesignInternalServerError() *ValidateSiteDesignInternalServerError {

	return &ValidateSiteDesignInternalServerError{}
}

// WriteResponse to the client
func (o *ValidateSiteDesignInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
