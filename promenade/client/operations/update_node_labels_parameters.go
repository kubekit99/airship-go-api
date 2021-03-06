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

// NewUpdateNodeLabelsParams creates a new UpdateNodeLabelsParams object
// with the default values initialized.
func NewUpdateNodeLabelsParams() *UpdateNodeLabelsParams {
	var ()
	return &UpdateNodeLabelsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNodeLabelsParamsWithTimeout creates a new UpdateNodeLabelsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateNodeLabelsParamsWithTimeout(timeout time.Duration) *UpdateNodeLabelsParams {
	var ()
	return &UpdateNodeLabelsParams{

		timeout: timeout,
	}
}

// NewUpdateNodeLabelsParamsWithContext creates a new UpdateNodeLabelsParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateNodeLabelsParamsWithContext(ctx context.Context) *UpdateNodeLabelsParams {
	var ()
	return &UpdateNodeLabelsParams{

		Context: ctx,
	}
}

// NewUpdateNodeLabelsParamsWithHTTPClient creates a new UpdateNodeLabelsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateNodeLabelsParamsWithHTTPClient(client *http.Client) *UpdateNodeLabelsParams {
	var ()
	return &UpdateNodeLabelsParams{
		HTTPClient: client,
	}
}

/*UpdateNodeLabelsParams contains all the parameters to send to the API endpoint
for the update node labels operation typically these are written to a http.Request
*/
type UpdateNodeLabelsParams struct {

	/*XAuthToken
	  A fernet keystone bearer token used for authentication and authorization

	*/
	XAuthToken *string
	/*NodeName*/
	NodeName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update node labels params
func (o *UpdateNodeLabelsParams) WithTimeout(timeout time.Duration) *UpdateNodeLabelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update node labels params
func (o *UpdateNodeLabelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update node labels params
func (o *UpdateNodeLabelsParams) WithContext(ctx context.Context) *UpdateNodeLabelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update node labels params
func (o *UpdateNodeLabelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update node labels params
func (o *UpdateNodeLabelsParams) WithHTTPClient(client *http.Client) *UpdateNodeLabelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update node labels params
func (o *UpdateNodeLabelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the update node labels params
func (o *UpdateNodeLabelsParams) WithXAuthToken(xAuthToken *string) *UpdateNodeLabelsParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the update node labels params
func (o *UpdateNodeLabelsParams) SetXAuthToken(xAuthToken *string) {
	o.XAuthToken = xAuthToken
}

// WithNodeName adds the nodeName to the update node labels params
func (o *UpdateNodeLabelsParams) WithNodeName(nodeName string) *UpdateNodeLabelsParams {
	o.SetNodeName(nodeName)
	return o
}

// SetNodeName adds the nodeName to the update node labels params
func (o *UpdateNodeLabelsParams) SetNodeName(nodeName string) {
	o.NodeName = nodeName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNodeLabelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param node-name
	if err := r.SetPathParam("node-name", o.NodeName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
