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

// NewValidateDesignParams creates a new ValidateDesignParams object
// with the default values initialized.
func NewValidateDesignParams() *ValidateDesignParams {
	var ()
	return &ValidateDesignParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewValidateDesignParamsWithTimeout creates a new ValidateDesignParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewValidateDesignParamsWithTimeout(timeout time.Duration) *ValidateDesignParams {
	var ()
	return &ValidateDesignParams{

		timeout: timeout,
	}
}

// NewValidateDesignParamsWithContext creates a new ValidateDesignParams object
// with the default values initialized, and the ability to set a context for a request
func NewValidateDesignParamsWithContext(ctx context.Context) *ValidateDesignParams {
	var ()
	return &ValidateDesignParams{

		Context: ctx,
	}
}

// NewValidateDesignParamsWithHTTPClient creates a new ValidateDesignParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewValidateDesignParamsWithHTTPClient(client *http.Client) *ValidateDesignParams {
	var ()
	return &ValidateDesignParams{
		HTTPClient: client,
	}
}

/*ValidateDesignParams contains all the parameters to send to the API endpoint
for the validate design operation typically these are written to a http.Request
*/
type ValidateDesignParams struct {

	/*XAuthToken
	  A fernet keystone bearer token used for authentication and authorization

	*/
	XAuthToken *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the validate design params
func (o *ValidateDesignParams) WithTimeout(timeout time.Duration) *ValidateDesignParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate design params
func (o *ValidateDesignParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate design params
func (o *ValidateDesignParams) WithContext(ctx context.Context) *ValidateDesignParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate design params
func (o *ValidateDesignParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate design params
func (o *ValidateDesignParams) WithHTTPClient(client *http.Client) *ValidateDesignParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate design params
func (o *ValidateDesignParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the validate design params
func (o *ValidateDesignParams) WithXAuthToken(xAuthToken *string) *ValidateDesignParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the validate design params
func (o *ValidateDesignParams) SetXAuthToken(xAuthToken *string) {
	o.XAuthToken = xAuthToken
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateDesignParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
