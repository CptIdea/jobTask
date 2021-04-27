// Code generated by go-swagger; DO NOT EDIT.

package petitions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// PetitionsDeleteURL generates an URL for the petitions delete operation
type PetitionsDeleteURL struct {
	PetitionID int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PetitionsDeleteURL) WithBasePath(bp string) *PetitionsDeleteURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PetitionsDeleteURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *PetitionsDeleteURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/petitions/{petitionId}"

	petitionID := swag.FormatInt64(o.PetitionID)
	if petitionID != "" {
		_path = strings.Replace(_path, "{petitionId}", petitionID, -1)
	} else {
		return nil, errors.New("petitionId is required on PetitionsDeleteURL")
	}

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *PetitionsDeleteURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *PetitionsDeleteURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *PetitionsDeleteURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on PetitionsDeleteURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on PetitionsDeleteURL")
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
func (o *PetitionsDeleteURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
