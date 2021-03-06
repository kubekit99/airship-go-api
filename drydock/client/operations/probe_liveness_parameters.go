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

// NewProbeLivenessParams creates a new ProbeLivenessParams object
// with the default values initialized.
func NewProbeLivenessParams() *ProbeLivenessParams {

	return &ProbeLivenessParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProbeLivenessParamsWithTimeout creates a new ProbeLivenessParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProbeLivenessParamsWithTimeout(timeout time.Duration) *ProbeLivenessParams {

	return &ProbeLivenessParams{

		timeout: timeout,
	}
}

// NewProbeLivenessParamsWithContext creates a new ProbeLivenessParams object
// with the default values initialized, and the ability to set a context for a request
func NewProbeLivenessParamsWithContext(ctx context.Context) *ProbeLivenessParams {

	return &ProbeLivenessParams{

		Context: ctx,
	}
}

// NewProbeLivenessParamsWithHTTPClient creates a new ProbeLivenessParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProbeLivenessParamsWithHTTPClient(client *http.Client) *ProbeLivenessParams {

	return &ProbeLivenessParams{
		HTTPClient: client,
	}
}

/*ProbeLivenessParams contains all the parameters to send to the API endpoint
for the probe liveness operation typically these are written to a http.Request
*/
type ProbeLivenessParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the probe liveness params
func (o *ProbeLivenessParams) WithTimeout(timeout time.Duration) *ProbeLivenessParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the probe liveness params
func (o *ProbeLivenessParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the probe liveness params
func (o *ProbeLivenessParams) WithContext(ctx context.Context) *ProbeLivenessParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the probe liveness params
func (o *ProbeLivenessParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the probe liveness params
func (o *ProbeLivenessParams) WithHTTPClient(client *http.Client) *ProbeLivenessParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the probe liveness params
func (o *ProbeLivenessParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ProbeLivenessParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
