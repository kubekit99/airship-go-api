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

// NewShowRevisionDeepDiffParams creates a new ShowRevisionDeepDiffParams object
// with the default values initialized.
func NewShowRevisionDeepDiffParams() *ShowRevisionDeepDiffParams {
	var ()
	return &ShowRevisionDeepDiffParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewShowRevisionDeepDiffParamsWithTimeout creates a new ShowRevisionDeepDiffParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewShowRevisionDeepDiffParamsWithTimeout(timeout time.Duration) *ShowRevisionDeepDiffParams {
	var ()
	return &ShowRevisionDeepDiffParams{

		timeout: timeout,
	}
}

// NewShowRevisionDeepDiffParamsWithContext creates a new ShowRevisionDeepDiffParams object
// with the default values initialized, and the ability to set a context for a request
func NewShowRevisionDeepDiffParamsWithContext(ctx context.Context) *ShowRevisionDeepDiffParams {
	var ()
	return &ShowRevisionDeepDiffParams{

		Context: ctx,
	}
}

// NewShowRevisionDeepDiffParamsWithHTTPClient creates a new ShowRevisionDeepDiffParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewShowRevisionDeepDiffParamsWithHTTPClient(client *http.Client) *ShowRevisionDeepDiffParams {
	var ()
	return &ShowRevisionDeepDiffParams{
		HTTPClient: client,
	}
}

/*ShowRevisionDeepDiffParams contains all the parameters to send to the API endpoint
for the show revision deep diff operation typically these are written to a http.Request
*/
type ShowRevisionDeepDiffParams struct {

	/*XAuthToken
	  A fernet keystone bearer token used for authentication and authorization

	*/
	XAuthToken *string
	/*ComparisonRevisionID*/
	ComparisonRevisionID string
	/*RevisionID*/
	RevisionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the show revision deep diff params
func (o *ShowRevisionDeepDiffParams) WithTimeout(timeout time.Duration) *ShowRevisionDeepDiffParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the show revision deep diff params
func (o *ShowRevisionDeepDiffParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the show revision deep diff params
func (o *ShowRevisionDeepDiffParams) WithContext(ctx context.Context) *ShowRevisionDeepDiffParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the show revision deep diff params
func (o *ShowRevisionDeepDiffParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the show revision deep diff params
func (o *ShowRevisionDeepDiffParams) WithHTTPClient(client *http.Client) *ShowRevisionDeepDiffParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the show revision deep diff params
func (o *ShowRevisionDeepDiffParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the show revision deep diff params
func (o *ShowRevisionDeepDiffParams) WithXAuthToken(xAuthToken *string) *ShowRevisionDeepDiffParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the show revision deep diff params
func (o *ShowRevisionDeepDiffParams) SetXAuthToken(xAuthToken *string) {
	o.XAuthToken = xAuthToken
}

// WithComparisonRevisionID adds the comparisonRevisionID to the show revision deep diff params
func (o *ShowRevisionDeepDiffParams) WithComparisonRevisionID(comparisonRevisionID string) *ShowRevisionDeepDiffParams {
	o.SetComparisonRevisionID(comparisonRevisionID)
	return o
}

// SetComparisonRevisionID adds the comparisonRevisionId to the show revision deep diff params
func (o *ShowRevisionDeepDiffParams) SetComparisonRevisionID(comparisonRevisionID string) {
	o.ComparisonRevisionID = comparisonRevisionID
}

// WithRevisionID adds the revisionID to the show revision deep diff params
func (o *ShowRevisionDeepDiffParams) WithRevisionID(revisionID string) *ShowRevisionDeepDiffParams {
	o.SetRevisionID(revisionID)
	return o
}

// SetRevisionID adds the revisionId to the show revision deep diff params
func (o *ShowRevisionDeepDiffParams) SetRevisionID(revisionID string) {
	o.RevisionID = revisionID
}

// WriteToRequest writes these params to a swagger request
func (o *ShowRevisionDeepDiffParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param comparison-revision-id
	if err := r.SetPathParam("comparison-revision-id", o.ComparisonRevisionID); err != nil {
		return err
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
