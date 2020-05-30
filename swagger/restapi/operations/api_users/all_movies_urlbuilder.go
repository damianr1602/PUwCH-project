// Code generated by go-swagger; DO NOT EDIT.

package api_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// AllMoviesURL generates an URL for the all movies operation
type AllMoviesURL struct {
	Limit        *int32
	SearchString *string
	Skip         *int32

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *AllMoviesURL) WithBasePath(bp string) *AllMoviesURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *AllMoviesURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *AllMoviesURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/movies"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/damianr1602/chmury/1.0.0"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var limitQ string
	if o.Limit != nil {
		limitQ = swag.FormatInt32(*o.Limit)
	}
	if limitQ != "" {
		qs.Set("limit", limitQ)
	}

	var searchStringQ string
	if o.SearchString != nil {
		searchStringQ = *o.SearchString
	}
	if searchStringQ != "" {
		qs.Set("searchString", searchStringQ)
	}

	var skipQ string
	if o.Skip != nil {
		skipQ = swag.FormatInt32(*o.Skip)
	}
	if skipQ != "" {
		qs.Set("skip", skipQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *AllMoviesURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *AllMoviesURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *AllMoviesURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on AllMoviesURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on AllMoviesURL")
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
func (o *AllMoviesURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}