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

// NewGetWorkflowByIDParams creates a new GetWorkflowByIDParams object
// with the default values initialized.
func NewGetWorkflowByIDParams() *GetWorkflowByIDParams {
	var ()
	return &GetWorkflowByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkflowByIDParamsWithTimeout creates a new GetWorkflowByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWorkflowByIDParamsWithTimeout(timeout time.Duration) *GetWorkflowByIDParams {
	var ()
	return &GetWorkflowByIDParams{

		timeout: timeout,
	}
}

// NewGetWorkflowByIDParamsWithContext creates a new GetWorkflowByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWorkflowByIDParamsWithContext(ctx context.Context) *GetWorkflowByIDParams {
	var ()
	return &GetWorkflowByIDParams{

		Context: ctx,
	}
}

// NewGetWorkflowByIDParamsWithHTTPClient creates a new GetWorkflowByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWorkflowByIDParamsWithHTTPClient(client *http.Client) *GetWorkflowByIDParams {
	var ()
	return &GetWorkflowByIDParams{
		HTTPClient: client,
	}
}

/*GetWorkflowByIDParams contains all the parameters to send to the API endpoint
for the get workflow by Id operation typically these are written to a http.Request
*/
type GetWorkflowByIDParams struct {

	/*XAuthToken
	  A fernet keystone bearer token used for authentication and authorization

	*/
	XAuthToken *string
	/*WorkflowID*/
	WorkflowID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get workflow by Id params
func (o *GetWorkflowByIDParams) WithTimeout(timeout time.Duration) *GetWorkflowByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workflow by Id params
func (o *GetWorkflowByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workflow by Id params
func (o *GetWorkflowByIDParams) WithContext(ctx context.Context) *GetWorkflowByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workflow by Id params
func (o *GetWorkflowByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workflow by Id params
func (o *GetWorkflowByIDParams) WithHTTPClient(client *http.Client) *GetWorkflowByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workflow by Id params
func (o *GetWorkflowByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the get workflow by Id params
func (o *GetWorkflowByIDParams) WithXAuthToken(xAuthToken *string) *GetWorkflowByIDParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the get workflow by Id params
func (o *GetWorkflowByIDParams) SetXAuthToken(xAuthToken *string) {
	o.XAuthToken = xAuthToken
}

// WithWorkflowID adds the workflowID to the get workflow by Id params
func (o *GetWorkflowByIDParams) WithWorkflowID(workflowID string) *GetWorkflowByIDParams {
	o.SetWorkflowID(workflowID)
	return o
}

// SetWorkflowID adds the workflowId to the get workflow by Id params
func (o *GetWorkflowByIDParams) SetWorkflowID(workflowID string) {
	o.WorkflowID = workflowID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkflowByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param workflow-id
	if err := r.SetPathParam("workflow-id", o.WorkflowID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}