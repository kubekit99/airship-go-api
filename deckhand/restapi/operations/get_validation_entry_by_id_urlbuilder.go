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

// GetValidationEntryByIDURL generates an URL for the get validation entry by Id operation
type GetValidationEntryByIDURL struct {
	EntryID        string
	RevisionID     string
	ValidationName string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetValidationEntryByIDURL) WithBasePath(bp string) *GetValidationEntryByIDURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetValidationEntryByIDURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetValidationEntryByIDURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/api/v1.0/revisions/{revision-id}/validations/{validation-name}/entries/{entry-id}"

	entryID := o.EntryID
	if entryID != "" {
		_path = strings.Replace(_path, "{entry-id}", entryID, -1)
	} else {
		return nil, errors.New("entryId is required on GetValidationEntryByIDURL")
	}

	revisionID := o.RevisionID
	if revisionID != "" {
		_path = strings.Replace(_path, "{revision-id}", revisionID, -1)
	} else {
		return nil, errors.New("revisionId is required on GetValidationEntryByIDURL")
	}

	validationName := o.ValidationName
	if validationName != "" {
		_path = strings.Replace(_path, "{validation-name}", validationName, -1)
	} else {
		return nil, errors.New("validationName is required on GetValidationEntryByIDURL")
	}

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetValidationEntryByIDURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetValidationEntryByIDURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetValidationEntryByIDURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetValidationEntryByIDURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetValidationEntryByIDURL")
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
func (o *GetValidationEntryByIDURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
