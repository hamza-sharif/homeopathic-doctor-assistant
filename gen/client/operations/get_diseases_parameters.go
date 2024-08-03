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
	"github.com/go-openapi/strfmt"
)

// NewGetDiseasesParams creates a new GetDiseasesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDiseasesParams() *GetDiseasesParams {
	return &GetDiseasesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDiseasesParamsWithTimeout creates a new GetDiseasesParams object
// with the ability to set a timeout on a request.
func NewGetDiseasesParamsWithTimeout(timeout time.Duration) *GetDiseasesParams {
	return &GetDiseasesParams{
		timeout: timeout,
	}
}

// NewGetDiseasesParamsWithContext creates a new GetDiseasesParams object
// with the ability to set a context for a request.
func NewGetDiseasesParamsWithContext(ctx context.Context) *GetDiseasesParams {
	return &GetDiseasesParams{
		Context: ctx,
	}
}

// NewGetDiseasesParamsWithHTTPClient creates a new GetDiseasesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDiseasesParamsWithHTTPClient(client *http.Client) *GetDiseasesParams {
	return &GetDiseasesParams{
		HTTPClient: client,
	}
}

/*
GetDiseasesParams contains all the parameters to send to the API endpoint

	for the get diseases operation.

	Typically these are written to a http.Request.
*/
type GetDiseasesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get diseases params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDiseasesParams) WithDefaults() *GetDiseasesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get diseases params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDiseasesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get diseases params
func (o *GetDiseasesParams) WithTimeout(timeout time.Duration) *GetDiseasesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get diseases params
func (o *GetDiseasesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get diseases params
func (o *GetDiseasesParams) WithContext(ctx context.Context) *GetDiseasesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get diseases params
func (o *GetDiseasesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get diseases params
func (o *GetDiseasesParams) WithHTTPClient(client *http.Client) *GetDiseasesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get diseases params
func (o *GetDiseasesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetDiseasesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
