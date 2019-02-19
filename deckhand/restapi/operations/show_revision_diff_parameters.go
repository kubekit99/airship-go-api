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

// NewShowRevisionDiffParams creates a new ShowRevisionDiffParams object
// no default values defined in spec.
func NewShowRevisionDiffParams() ShowRevisionDiffParams {

	return ShowRevisionDiffParams{}
}

// ShowRevisionDiffParams contains all the bound params for the show revision diff operation
// typically these are obtained from a http.Request
//
// swagger:parameters showRevisionDiff
type ShowRevisionDiffParams struct {

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
	ComparisonRevisionID string
	/*
	  Required: true
	  In: path
	*/
	RevisionID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewShowRevisionDiffParams() beforehand.
func (o *ShowRevisionDiffParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindXAuthToken(r.Header[http.CanonicalHeaderKey("X-Auth-Token")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	rComparisonRevisionID, rhkComparisonRevisionID, _ := route.Params.GetOK("comparison-revision-id")
	if err := o.bindComparisonRevisionID(rComparisonRevisionID, rhkComparisonRevisionID, route.Formats); err != nil {
		res = append(res, err)
	}

	rRevisionID, rhkRevisionID, _ := route.Params.GetOK("revision-id")
	if err := o.bindRevisionID(rRevisionID, rhkRevisionID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindXAuthToken binds and validates parameter XAuthToken from header.
func (o *ShowRevisionDiffParams) bindXAuthToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

// bindComparisonRevisionID binds and validates parameter ComparisonRevisionID from path.
func (o *ShowRevisionDiffParams) bindComparisonRevisionID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ComparisonRevisionID = raw

	return nil
}

// bindRevisionID binds and validates parameter RevisionID from path.
func (o *ShowRevisionDiffParams) bindRevisionID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.RevisionID = raw

	return nil
}
