// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetPatientsURL generates an URL for the get patients operation
type GetPatientsURL struct {
	Cnic                *string
	DateTime            *strfmt.DateTime
	FatherOrHusbandName *string
	Limit               *int32
	MobileNo            *string
	Name                *string
	Offset              *int32

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetPatientsURL) WithBasePath(bp string) *GetPatientsURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetPatientsURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetPatientsURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/patients"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var cnicQ string
	if o.Cnic != nil {
		cnicQ = *o.Cnic
	}
	if cnicQ != "" {
		qs.Set("cnic", cnicQ)
	}

	var dateTimeQ string
	if o.DateTime != nil {
		dateTimeQ = o.DateTime.String()
	}
	if dateTimeQ != "" {
		qs.Set("dateTime", dateTimeQ)
	}

	var fatherOrHusbandNameQ string
	if o.FatherOrHusbandName != nil {
		fatherOrHusbandNameQ = *o.FatherOrHusbandName
	}
	if fatherOrHusbandNameQ != "" {
		qs.Set("fatherOrHusbandName", fatherOrHusbandNameQ)
	}

	var limitQ string
	if o.Limit != nil {
		limitQ = swag.FormatInt32(*o.Limit)
	}
	if limitQ != "" {
		qs.Set("limit", limitQ)
	}

	var mobileNoQ string
	if o.MobileNo != nil {
		mobileNoQ = *o.MobileNo
	}
	if mobileNoQ != "" {
		qs.Set("mobileNo", mobileNoQ)
	}

	var nameQ string
	if o.Name != nil {
		nameQ = *o.Name
	}
	if nameQ != "" {
		qs.Set("name", nameQ)
	}

	var offsetQ string
	if o.Offset != nil {
		offsetQ = swag.FormatInt32(*o.Offset)
	}
	if offsetQ != "" {
		qs.Set("offset", offsetQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetPatientsURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetPatientsURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetPatientsURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetPatientsURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetPatientsURL")
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
func (o *GetPatientsURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
