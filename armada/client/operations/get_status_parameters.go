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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetStatusParams creates a new GetStatusParams object
// with the default values initialized.
func NewGetStatusParams() *GetStatusParams {
	var (
		tillerHostDefault = string("None")
	)
	return &GetStatusParams{
		TillerHost: &tillerHostDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStatusParamsWithTimeout creates a new GetStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStatusParamsWithTimeout(timeout time.Duration) *GetStatusParams {
	var (
		tillerHostDefault = string("None")
	)
	return &GetStatusParams{
		TillerHost: &tillerHostDefault,

		timeout: timeout,
	}
}

// NewGetStatusParamsWithContext creates a new GetStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStatusParamsWithContext(ctx context.Context) *GetStatusParams {
	var (
		tillerHostDefault = string("None")
	)
	return &GetStatusParams{
		TillerHost: &tillerHostDefault,

		Context: ctx,
	}
}

// NewGetStatusParamsWithHTTPClient creates a new GetStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStatusParamsWithHTTPClient(client *http.Client) *GetStatusParams {
	var (
		tillerHostDefault = string("None")
	)
	return &GetStatusParams{
		TillerHost: &tillerHostDefault,
		HTTPClient: client,
	}
}

/*GetStatusParams contains all the parameters to send to the API endpoint
for the get status operation typically these are written to a http.Request
*/
type GetStatusParams struct {

	/*XAuthToken
	  A fernet keystone bearer token used for authentication and authorization

	*/
	XAuthToken *string
	/*TillerHost
	  Hostname of the Tiller server

	*/
	TillerHost *string
	/*TillerNamespace
	  Tiller namespace. Default is the value of `CONF.tiller_namespace`

	*/
	TillerNamespace *string
	/*TillerPort
	  Port number of the Tiller server. Default is the value of `CONF.tiller_port`.

	*/
	TillerPort *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get status params
func (o *GetStatusParams) WithTimeout(timeout time.Duration) *GetStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get status params
func (o *GetStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get status params
func (o *GetStatusParams) WithContext(ctx context.Context) *GetStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get status params
func (o *GetStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get status params
func (o *GetStatusParams) WithHTTPClient(client *http.Client) *GetStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get status params
func (o *GetStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the get status params
func (o *GetStatusParams) WithXAuthToken(xAuthToken *string) *GetStatusParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the get status params
func (o *GetStatusParams) SetXAuthToken(xAuthToken *string) {
	o.XAuthToken = xAuthToken
}

// WithTillerHost adds the tillerHost to the get status params
func (o *GetStatusParams) WithTillerHost(tillerHost *string) *GetStatusParams {
	o.SetTillerHost(tillerHost)
	return o
}

// SetTillerHost adds the tillerHost to the get status params
func (o *GetStatusParams) SetTillerHost(tillerHost *string) {
	o.TillerHost = tillerHost
}

// WithTillerNamespace adds the tillerNamespace to the get status params
func (o *GetStatusParams) WithTillerNamespace(tillerNamespace *string) *GetStatusParams {
	o.SetTillerNamespace(tillerNamespace)
	return o
}

// SetTillerNamespace adds the tillerNamespace to the get status params
func (o *GetStatusParams) SetTillerNamespace(tillerNamespace *string) {
	o.TillerNamespace = tillerNamespace
}

// WithTillerPort adds the tillerPort to the get status params
func (o *GetStatusParams) WithTillerPort(tillerPort *int64) *GetStatusParams {
	o.SetTillerPort(tillerPort)
	return o
}

// SetTillerPort adds the tillerPort to the get status params
func (o *GetStatusParams) SetTillerPort(tillerPort *int64) {
	o.TillerPort = tillerPort
}

// WriteToRequest writes these params to a swagger request
func (o *GetStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.TillerHost != nil {

		// query param tiller_host
		var qrTillerHost string
		if o.TillerHost != nil {
			qrTillerHost = *o.TillerHost
		}
		qTillerHost := qrTillerHost
		if qTillerHost != "" {
			if err := r.SetQueryParam("tiller_host", qTillerHost); err != nil {
				return err
			}
		}

	}

	if o.TillerNamespace != nil {

		// query param tiller_namespace
		var qrTillerNamespace string
		if o.TillerNamespace != nil {
			qrTillerNamespace = *o.TillerNamespace
		}
		qTillerNamespace := qrTillerNamespace
		if qTillerNamespace != "" {
			if err := r.SetQueryParam("tiller_namespace", qTillerNamespace); err != nil {
				return err
			}
		}

	}

	if o.TillerPort != nil {

		// query param tiller_port
		var qrTillerPort int64
		if o.TillerPort != nil {
			qrTillerPort = *o.TillerPort
		}
		qTillerPort := swag.FormatInt64(qrTillerPort)
		if qTillerPort != "" {
			if err := r.SetQueryParam("tiller_port", qTillerPort); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
