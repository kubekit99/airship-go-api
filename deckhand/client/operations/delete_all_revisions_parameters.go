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

// NewDeleteAllRevisionsParams creates a new DeleteAllRevisionsParams object
// with the default values initialized.
func NewDeleteAllRevisionsParams() *DeleteAllRevisionsParams {
	var ()
	return &DeleteAllRevisionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAllRevisionsParamsWithTimeout creates a new DeleteAllRevisionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAllRevisionsParamsWithTimeout(timeout time.Duration) *DeleteAllRevisionsParams {
	var ()
	return &DeleteAllRevisionsParams{

		timeout: timeout,
	}
}

// NewDeleteAllRevisionsParamsWithContext creates a new DeleteAllRevisionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAllRevisionsParamsWithContext(ctx context.Context) *DeleteAllRevisionsParams {
	var ()
	return &DeleteAllRevisionsParams{

		Context: ctx,
	}
}

// NewDeleteAllRevisionsParamsWithHTTPClient creates a new DeleteAllRevisionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAllRevisionsParamsWithHTTPClient(client *http.Client) *DeleteAllRevisionsParams {
	var ()
	return &DeleteAllRevisionsParams{
		HTTPClient: client,
	}
}

/*DeleteAllRevisionsParams contains all the parameters to send to the API endpoint
for the delete all revisions operation typically these are written to a http.Request
*/
type DeleteAllRevisionsParams struct {

	/*XAuthToken
	  A fernet keystone bearer token used for authentication and authorization

	*/
	XAuthToken *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete all revisions params
func (o *DeleteAllRevisionsParams) WithTimeout(timeout time.Duration) *DeleteAllRevisionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete all revisions params
func (o *DeleteAllRevisionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete all revisions params
func (o *DeleteAllRevisionsParams) WithContext(ctx context.Context) *DeleteAllRevisionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete all revisions params
func (o *DeleteAllRevisionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete all revisions params
func (o *DeleteAllRevisionsParams) WithHTTPClient(client *http.Client) *DeleteAllRevisionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete all revisions params
func (o *DeleteAllRevisionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the delete all revisions params
func (o *DeleteAllRevisionsParams) WithXAuthToken(xAuthToken *string) *DeleteAllRevisionsParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the delete all revisions params
func (o *DeleteAllRevisionsParams) SetXAuthToken(xAuthToken *string) {
	o.XAuthToken = xAuthToken
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAllRevisionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}