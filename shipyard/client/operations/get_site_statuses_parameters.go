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

// NewGetSiteStatusesParams creates a new GetSiteStatusesParams object
// with the default values initialized.
func NewGetSiteStatusesParams() *GetSiteStatusesParams {
	var ()
	return &GetSiteStatusesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSiteStatusesParamsWithTimeout creates a new GetSiteStatusesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSiteStatusesParamsWithTimeout(timeout time.Duration) *GetSiteStatusesParams {
	var ()
	return &GetSiteStatusesParams{

		timeout: timeout,
	}
}

// NewGetSiteStatusesParamsWithContext creates a new GetSiteStatusesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSiteStatusesParamsWithContext(ctx context.Context) *GetSiteStatusesParams {
	var ()
	return &GetSiteStatusesParams{

		Context: ctx,
	}
}

// NewGetSiteStatusesParamsWithHTTPClient creates a new GetSiteStatusesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSiteStatusesParamsWithHTTPClient(client *http.Client) *GetSiteStatusesParams {
	var ()
	return &GetSiteStatusesParams{
		HTTPClient: client,
	}
}

/*GetSiteStatusesParams contains all the parameters to send to the API endpoint
for the get site statuses operation typically these are written to a http.Request
*/
type GetSiteStatusesParams struct {

	/*XAuthToken
	  A fernet keystone bearer token used for authentication and authorization

	*/
	XAuthToken *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get site statuses params
func (o *GetSiteStatusesParams) WithTimeout(timeout time.Duration) *GetSiteStatusesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get site statuses params
func (o *GetSiteStatusesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get site statuses params
func (o *GetSiteStatusesParams) WithContext(ctx context.Context) *GetSiteStatusesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get site statuses params
func (o *GetSiteStatusesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get site statuses params
func (o *GetSiteStatusesParams) WithHTTPClient(client *http.Client) *GetSiteStatusesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get site statuses params
func (o *GetSiteStatusesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the get site statuses params
func (o *GetSiteStatusesParams) WithXAuthToken(xAuthToken *string) *GetSiteStatusesParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the get site statuses params
func (o *GetSiteStatusesParams) SetXAuthToken(xAuthToken *string) {
	o.XAuthToken = xAuthToken
}

// WriteToRequest writes these params to a swagger request
func (o *GetSiteStatusesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
