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

// NewRetrieveConfigDocsClearTextByCollectionIDParams creates a new RetrieveConfigDocsClearTextByCollectionIDParams object
// no default values defined in spec.
func NewRetrieveConfigDocsClearTextByCollectionIDParams() RetrieveConfigDocsClearTextByCollectionIDParams {

	return RetrieveConfigDocsClearTextByCollectionIDParams{}
}

// RetrieveConfigDocsClearTextByCollectionIDParams contains all the bound params for the retrieve config docs clear text by collection Id operation
// typically these are obtained from a http.Request
//
// swagger:parameters retrieveConfigDocsClearTextByCollectionId
type RetrieveConfigDocsClearTextByCollectionIDParams struct {

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
	CollectionID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewRetrieveConfigDocsClearTextByCollectionIDParams() beforehand.
func (o *RetrieveConfigDocsClearTextByCollectionIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindXAuthToken(r.Header[http.CanonicalHeaderKey("X-Auth-Token")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	rCollectionID, rhkCollectionID, _ := route.Params.GetOK("collection-id")
	if err := o.bindCollectionID(rCollectionID, rhkCollectionID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindXAuthToken binds and validates parameter XAuthToken from header.
func (o *RetrieveConfigDocsClearTextByCollectionIDParams) bindXAuthToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

// bindCollectionID binds and validates parameter CollectionID from path.
func (o *RetrieveConfigDocsClearTextByCollectionIDParams) bindCollectionID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.CollectionID = raw

	return nil
}
