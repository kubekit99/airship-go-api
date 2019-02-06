// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// GetActionStepByIDURL generates an URL for the get action step by Id operation
type GetActionStepByIDURL struct {
	ActionID string
	StepID   string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetActionStepByIDURL) WithBasePath(bp string) *GetActionStepByIDURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetActionStepByIDURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetActionStepByIDURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/api/v1.0/actions/{action-id}/steps/{step-id}"

	actionID := o.ActionID
	if actionID != "" {
		_path = strings.Replace(_path, "{action-id}", actionID, -1)
	} else {
		return nil, errors.New("actionId is required on GetActionStepByIDURL")
	}

	stepID := o.StepID
	if stepID != "" {
		_path = strings.Replace(_path, "{step-id}", stepID, -1)
	} else {
		return nil, errors.New("stepId is required on GetActionStepByIDURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetActionStepByIDURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetActionStepByIDURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetActionStepByIDURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetActionStepByIDURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetActionStepByIDURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetActionStepByIDURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}