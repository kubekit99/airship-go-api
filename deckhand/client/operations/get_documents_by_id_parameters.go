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

// NewGetDocumentsByIDParams creates a new GetDocumentsByIDParams object
// with the default values initialized.
func NewGetDocumentsByIDParams() *GetDocumentsByIDParams {
	var ()
	return &GetDocumentsByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDocumentsByIDParamsWithTimeout creates a new GetDocumentsByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDocumentsByIDParamsWithTimeout(timeout time.Duration) *GetDocumentsByIDParams {
	var ()
	return &GetDocumentsByIDParams{

		timeout: timeout,
	}
}

// NewGetDocumentsByIDParamsWithContext creates a new GetDocumentsByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDocumentsByIDParamsWithContext(ctx context.Context) *GetDocumentsByIDParams {
	var ()
	return &GetDocumentsByIDParams{

		Context: ctx,
	}
}

// NewGetDocumentsByIDParamsWithHTTPClient creates a new GetDocumentsByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDocumentsByIDParamsWithHTTPClient(client *http.Client) *GetDocumentsByIDParams {
	var ()
	return &GetDocumentsByIDParams{
		HTTPClient: client,
	}
}

/*GetDocumentsByIDParams contains all the parameters to send to the API endpoint
for the get documents by Id operation typically these are written to a http.Request
*/
type GetDocumentsByIDParams struct {

	/*XAuthToken
	  A fernet keystone bearer token used for authentication and authorization

	*/
	XAuthToken *string
	/*RevisionID*/
	RevisionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get documents by Id params
func (o *GetDocumentsByIDParams) WithTimeout(timeout time.Duration) *GetDocumentsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get documents by Id params
func (o *GetDocumentsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get documents by Id params
func (o *GetDocumentsByIDParams) WithContext(ctx context.Context) *GetDocumentsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get documents by Id params
func (o *GetDocumentsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get documents by Id params
func (o *GetDocumentsByIDParams) WithHTTPClient(client *http.Client) *GetDocumentsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get documents by Id params
func (o *GetDocumentsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the get documents by Id params
func (o *GetDocumentsByIDParams) WithXAuthToken(xAuthToken *string) *GetDocumentsByIDParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the get documents by Id params
func (o *GetDocumentsByIDParams) SetXAuthToken(xAuthToken *string) {
	o.XAuthToken = xAuthToken
}

// WithRevisionID adds the revisionID to the get documents by Id params
func (o *GetDocumentsByIDParams) WithRevisionID(revisionID string) *GetDocumentsByIDParams {
	o.SetRevisionID(revisionID)
	return o
}

// SetRevisionID adds the revisionId to the get documents by Id params
func (o *GetDocumentsByIDParams) SetRevisionID(revisionID string) {
	o.RevisionID = revisionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDocumentsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param revision-id
	if err := r.SetPathParam("revision-id", o.RevisionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
