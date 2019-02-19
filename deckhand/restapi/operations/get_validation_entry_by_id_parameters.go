// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetValidationEntryByIDParams creates a new GetValidationEntryByIDParams object
// no default values defined in spec.
func NewGetValidationEntryByIDParams() GetValidationEntryByIDParams {

	return GetValidationEntryByIDParams{}
}

// GetValidationEntryByIDParams contains all the bound params for the get validation entry by Id operation
// typically these are obtained from a http.Request
//
// swagger:parameters getValidationEntryById
type GetValidationEntryByIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*A fernet keystone bearer token used for authentication and authorization
	  In: header
	*/
	XAuthToken *string
	/*
	  Required: true
	  In: path
	*/
	EntryID string
	/*
	  Required: true
	  In: path
	*/
	RevisionID string
	/*
	  Required: true
	  In: path
	*/
	ValidationName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetValidationEntryByIDParams() beforehand.
func (o *GetValidationEntryByIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindXAuthToken(r.Header[http.CanonicalHeaderKey("X-Auth-Token")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	rEntryID, rhkEntryID, _ := route.Params.GetOK("entry-id")
	if err := o.bindEntryID(rEntryID, rhkEntryID, route.Formats); err != nil {
		res = append(res, err)
	}

	rRevisionID, rhkRevisionID, _ := route.Params.GetOK("revision-id")
	if err := o.bindRevisionID(rRevisionID, rhkRevisionID, route.Formats); err != nil {
		res = append(res, err)
	}

	rValidationName, rhkValidationName, _ := route.Params.GetOK("validation-name")
	if err := o.bindValidationName(rValidationName, rhkValidationName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindXAuthToken binds and validates parameter XAuthToken from header.
func (o *GetValidationEntryByIDParams) bindXAuthToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.XAuthToken = &raw

	return nil
}

// bindEntryID binds and validates parameter EntryID from path.
func (o *GetValidationEntryByIDParams) bindEntryID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.EntryID = raw

	return nil
}

// bindRevisionID binds and validates parameter RevisionID from path.
func (o *GetValidationEntryByIDParams) bindRevisionID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.RevisionID = raw

	return nil
}

// bindValidationName binds and validates parameter ValidationName from path.
func (o *GetValidationEntryByIDParams) bindValidationName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ValidationName = raw

	return nil
}
