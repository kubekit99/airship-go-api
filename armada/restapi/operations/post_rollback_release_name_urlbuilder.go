// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// PostRollbackReleaseNameURL generates an URL for the post rollback release name operation
type PostRollbackReleaseNameURL struct {
	ReleaseName string

	DryRun          *bool
	Force           *bool
	RecreatePods    *bool
	TillerHost      *string
	TillerNamespace *string
	TillerPort      *int64
	Timeout         *int64
	Version         *int64
	Wait            *bool

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PostRollbackReleaseNameURL) WithBasePath(bp string) *PostRollbackReleaseNameURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PostRollbackReleaseNameURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *PostRollbackReleaseNameURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/api/v1.0/rollback/{release_name}"

	releaseName := o.ReleaseName
	if releaseName != "" {
		_path = strings.Replace(_path, "{release_name}", releaseName, -1)
	} else {
		return nil, errors.New("releaseName is required on PostRollbackReleaseNameURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var dryRun string
	if o.DryRun != nil {
		dryRun = swag.FormatBool(*o.DryRun)
	}
	if dryRun != "" {
		qs.Set("dry_run", dryRun)
	}

	var force string
	if o.Force != nil {
		force = swag.FormatBool(*o.Force)
	}
	if force != "" {
		qs.Set("force", force)
	}

	var recreatePods string
	if o.RecreatePods != nil {
		recreatePods = swag.FormatBool(*o.RecreatePods)
	}
	if recreatePods != "" {
		qs.Set("recreate_pods", recreatePods)
	}

	var tillerHost string
	if o.TillerHost != nil {
		tillerHost = *o.TillerHost
	}
	if tillerHost != "" {
		qs.Set("tiller_host", tillerHost)
	}

	var tillerNamespace string
	if o.TillerNamespace != nil {
		tillerNamespace = *o.TillerNamespace
	}
	if tillerNamespace != "" {
		qs.Set("tiller_namespace", tillerNamespace)
	}

	var tillerPort string
	if o.TillerPort != nil {
		tillerPort = swag.FormatInt64(*o.TillerPort)
	}
	if tillerPort != "" {
		qs.Set("tiller_port", tillerPort)
	}

	var timeout string
	if o.Timeout != nil {
		timeout = swag.FormatInt64(*o.Timeout)
	}
	if timeout != "" {
		qs.Set("timeout", timeout)
	}

	var version string
	if o.Version != nil {
		version = swag.FormatInt64(*o.Version)
	}
	if version != "" {
		qs.Set("version", version)
	}

	var wait string
	if o.Wait != nil {
		wait = swag.FormatBool(*o.Wait)
	}
	if wait != "" {
		qs.Set("wait", wait)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *PostRollbackReleaseNameURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *PostRollbackReleaseNameURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *PostRollbackReleaseNameURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on PostRollbackReleaseNameURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on PostRollbackReleaseNameURL")
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
func (o *PostRollbackReleaseNameURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
