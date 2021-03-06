// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// PostApplyManifestURL generates an URL for the post apply manifest operation
type PostApplyManifestURL struct {
	DisableUpdatePost  *bool
	DisableUpdatePre   *bool
	DryRun             *bool
	EnableChartCleanup *bool
	TargetManifest     *string
	TillerHost         *string
	TillerNamespace    *string
	TillerPort         *int64
	Timeout            *int64
	Wait               *bool

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PostApplyManifestURL) WithBasePath(bp string) *PostApplyManifestURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PostApplyManifestURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *PostApplyManifestURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/api/v1.0/apply"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var disableUpdatePost string
	if o.DisableUpdatePost != nil {
		disableUpdatePost = swag.FormatBool(*o.DisableUpdatePost)
	}
	if disableUpdatePost != "" {
		qs.Set("disable_update_post", disableUpdatePost)
	}

	var disableUpdatePre string
	if o.DisableUpdatePre != nil {
		disableUpdatePre = swag.FormatBool(*o.DisableUpdatePre)
	}
	if disableUpdatePre != "" {
		qs.Set("disable_update_pre", disableUpdatePre)
	}

	var dryRun string
	if o.DryRun != nil {
		dryRun = swag.FormatBool(*o.DryRun)
	}
	if dryRun != "" {
		qs.Set("dry_run", dryRun)
	}

	var enableChartCleanup string
	if o.EnableChartCleanup != nil {
		enableChartCleanup = swag.FormatBool(*o.EnableChartCleanup)
	}
	if enableChartCleanup != "" {
		qs.Set("enable_chart_cleanup", enableChartCleanup)
	}

	var targetManifest string
	if o.TargetManifest != nil {
		targetManifest = *o.TargetManifest
	}
	if targetManifest != "" {
		qs.Set("target_manifest", targetManifest)
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
func (o *PostApplyManifestURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *PostApplyManifestURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *PostApplyManifestURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on PostApplyManifestURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on PostApplyManifestURL")
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
func (o *PostApplyManifestURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
