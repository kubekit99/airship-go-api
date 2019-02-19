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

// NewRetrieveRenderedCleartextConfigDocsParams creates a new RetrieveRenderedCleartextConfigDocsParams object
// with the default values initialized.
func NewRetrieveRenderedCleartextConfigDocsParams() *RetrieveRenderedCleartextConfigDocsParams {
	var ()
	return &RetrieveRenderedCleartextConfigDocsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRetrieveRenderedCleartextConfigDocsParamsWithTimeout creates a new RetrieveRenderedCleartextConfigDocsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRetrieveRenderedCleartextConfigDocsParamsWithTimeout(timeout time.Duration) *RetrieveRenderedCleartextConfigDocsParams {
	var ()
	return &RetrieveRenderedCleartextConfigDocsParams{

		timeout: timeout,
	}
}

// NewRetrieveRenderedCleartextConfigDocsParamsWithContext creates a new RetrieveRenderedCleartextConfigDocsParams object
// with the default values initialized, and the ability to set a context for a request
func NewRetrieveRenderedCleartextConfigDocsParamsWithContext(ctx context.Context) *RetrieveRenderedCleartextConfigDocsParams {
	var ()
	return &RetrieveRenderedCleartextConfigDocsParams{

		Context: ctx,
	}
}

// NewRetrieveRenderedCleartextConfigDocsParamsWithHTTPClient creates a new RetrieveRenderedCleartextConfigDocsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRetrieveRenderedCleartextConfigDocsParamsWithHTTPClient(client *http.Client) *RetrieveRenderedCleartextConfigDocsParams {
	var ()
	return &RetrieveRenderedCleartextConfigDocsParams{
		HTTPClient: client,
	}
}

/*RetrieveRenderedCleartextConfigDocsParams contains all the parameters to send to the API endpoint
for the retrieve rendered cleartext config docs operation typically these are written to a http.Request
*/
type RetrieveRenderedCleartextConfigDocsParams struct {

	/*XAuthToken
	  A fernet keystone bearer token used for authentication and authorization

	*/
	XAuthToken *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the retrieve rendered cleartext config docs params
func (o *RetrieveRenderedCleartextConfigDocsParams) WithTimeout(timeout time.Duration) *RetrieveRenderedCleartextConfigDocsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retrieve rendered cleartext config docs params
func (o *RetrieveRenderedCleartextConfigDocsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retrieve rendered cleartext config docs params
func (o *RetrieveRenderedCleartextConfigDocsParams) WithContext(ctx context.Context) *RetrieveRenderedCleartextConfigDocsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retrieve rendered cleartext config docs params
func (o *RetrieveRenderedCleartextConfigDocsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retrieve rendered cleartext config docs params
func (o *RetrieveRenderedCleartextConfigDocsParams) WithHTTPClient(client *http.Client) *RetrieveRenderedCleartextConfigDocsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retrieve rendered cleartext config docs params
func (o *RetrieveRenderedCleartextConfigDocsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the retrieve rendered cleartext config docs params
func (o *RetrieveRenderedCleartextConfigDocsParams) WithXAuthToken(xAuthToken *string) *RetrieveRenderedCleartextConfigDocsParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the retrieve rendered cleartext config docs params
func (o *RetrieveRenderedCleartextConfigDocsParams) SetXAuthToken(xAuthToken *string) {
	o.XAuthToken = xAuthToken
}

// WriteToRequest writes these params to a swagger request
func (o *RetrieveRenderedCleartextConfigDocsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
