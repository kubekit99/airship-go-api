// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetConfigDocByIDParams creates a new GetConfigDocByIDParams object
// with the default values initialized.
func NewGetConfigDocByIDParams() *GetConfigDocByIDParams {
	var ()
	return &GetConfigDocByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetConfigDocByIDParamsWithTimeout creates a new GetConfigDocByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetConfigDocByIDParamsWithTimeout(timeout time.Duration) *GetConfigDocByIDParams {
	var ()
	return &GetConfigDocByIDParams{

		timeout: timeout,
	}
}

// NewGetConfigDocByIDParamsWithContext creates a new GetConfigDocByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetConfigDocByIDParamsWithContext(ctx context.Context) *GetConfigDocByIDParams {
	var ()
	return &GetConfigDocByIDParams{

		Context: ctx,
	}
}

// NewGetConfigDocByIDParamsWithHTTPClient creates a new GetConfigDocByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetConfigDocByIDParamsWithHTTPClient(client *http.Client) *GetConfigDocByIDParams {
	var ()
	return &GetConfigDocByIDParams{
		HTTPClient: client,
	}
}

/*GetConfigDocByIDParams contains all the parameters to send to the API endpoint
for the get config doc by Id operation typically these are written to a http.Request
*/
type GetConfigDocByIDParams struct {

	/*XAuthToken
	  A fernet keystone bearer token used for authentication and authorization

	*/
	XAuthToken *string
	/*CollectionID*/
	CollectionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get config doc by Id params
func (o *GetConfigDocByIDParams) WithTimeout(timeout time.Duration) *GetConfigDocByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get config doc by Id params
func (o *GetConfigDocByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get config doc by Id params
func (o *GetConfigDocByIDParams) WithContext(ctx context.Context) *GetConfigDocByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get config doc by Id params
func (o *GetConfigDocByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get config doc by Id params
func (o *GetConfigDocByIDParams) WithHTTPClient(client *http.Client) *GetConfigDocByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get config doc by Id params
func (o *GetConfigDocByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the get config doc by Id params
func (o *GetConfigDocByIDParams) WithXAuthToken(xAuthToken *string) *GetConfigDocByIDParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the get config doc by Id params
func (o *GetConfigDocByIDParams) SetXAuthToken(xAuthToken *string) {
	o.XAuthToken = xAuthToken
}

// WithCollectionID adds the collectionID to the get config doc by Id params
func (o *GetConfigDocByIDParams) WithCollectionID(collectionID string) *GetConfigDocByIDParams {
	o.SetCollectionID(collectionID)
	return o
}

// SetCollectionID adds the collectionId to the get config doc by Id params
func (o *GetConfigDocByIDParams) SetCollectionID(collectionID string) {
	o.CollectionID = collectionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetConfigDocByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XAuthToken != nil {

		// header param X-Auth-Token
		if err := r.SetHeaderParam("X-Auth-Token", *o.XAuthToken); err != nil {
			return err
		}

	}

	// path param collection-id
	if err := r.SetPathParam("collection-id", o.CollectionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
