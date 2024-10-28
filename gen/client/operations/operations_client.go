// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new operations API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new operations API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeletePatientsPatientID(params *DeletePatientsPatientIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeletePatientsPatientIDCreated, error)

	GetDashboard(params *GetDashboardParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDashboardOK, error)

	GetDiseases(params *GetDiseasesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDiseasesOK, error)

	GetMedicines(params *GetMedicinesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetMedicinesOK, error)

	GetPatients(params *GetPatientsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPatientsOK, error)

	PostDiseases(params *PostDiseasesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostDiseasesCreated, error)

	PostForgetPassword(params *PostForgetPasswordParams, opts ...ClientOption) (*PostForgetPasswordOK, error)

	PostLogin(params *PostLoginParams, opts ...ClientOption) (*PostLoginOK, error)

	PostMedicines(params *PostMedicinesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostMedicinesCreated, error)

	PostPatients(params *PostPatientsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostPatientsCreated, error)

	PutUpdatePassword(params *PutUpdatePasswordParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutUpdatePasswordOK, error)

	PutUpdatePrice(params *PutUpdatePriceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutUpdatePriceOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeletePatientsPatientID deletes patient
*/
func (a *Client) DeletePatientsPatientID(params *DeletePatientsPatientIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeletePatientsPatientIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePatientsPatientIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeletePatientsPatientID",
		Method:             "DELETE",
		PathPattern:        "/patients/{patient_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeletePatientsPatientIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeletePatientsPatientIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeletePatientsPatientID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDashboard gets dashboard details
*/
func (a *Client) GetDashboard(params *GetDashboardParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDashboardOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDashboardParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDashboard",
		Method:             "GET",
		PathPattern:        "/dashboard",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDashboardReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDashboardOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDashboard: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDiseases gets diseases
*/
func (a *Client) GetDiseases(params *GetDiseasesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDiseasesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDiseasesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDiseases",
		Method:             "GET",
		PathPattern:        "/diseases",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDiseasesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDiseasesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDiseases: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetMedicines gets medicines
*/
func (a *Client) GetMedicines(params *GetMedicinesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetMedicinesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMedicinesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetMedicines",
		Method:             "GET",
		PathPattern:        "/medicines",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetMedicinesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetMedicinesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetMedicines: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPatients gets patients
*/
func (a *Client) GetPatients(params *GetPatientsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPatientsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPatientsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetPatients",
		Method:             "GET",
		PathPattern:        "/patients",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPatientsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPatientsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPatients: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostDiseases adds medicine
*/
func (a *Client) PostDiseases(params *PostDiseasesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostDiseasesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostDiseasesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostDiseases",
		Method:             "POST",
		PathPattern:        "/diseases",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostDiseasesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostDiseasesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostDiseases: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostForgetPassword forgets password
*/
func (a *Client) PostForgetPassword(params *PostForgetPasswordParams, opts ...ClientOption) (*PostForgetPasswordOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostForgetPasswordParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostForgetPassword",
		Method:             "POST",
		PathPattern:        "/forget-password",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostForgetPasswordReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostForgetPasswordOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostForgetPassword: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostLogin logins user
*/
func (a *Client) PostLogin(params *PostLoginParams, opts ...ClientOption) (*PostLoginOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLoginParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostLogin",
		Method:             "POST",
		PathPattern:        "/login",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostLoginReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostLoginOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostLogin: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostMedicines adds medicine
*/
func (a *Client) PostMedicines(params *PostMedicinesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostMedicinesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostMedicinesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostMedicines",
		Method:             "POST",
		PathPattern:        "/medicines",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostMedicinesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostMedicinesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostMedicines: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostPatients adds patient
*/
func (a *Client) PostPatients(params *PostPatientsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostPatientsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPatientsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostPatients",
		Method:             "POST",
		PathPattern:        "/patients",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostPatientsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostPatientsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostPatients: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PutUpdatePassword updates password
*/
func (a *Client) PutUpdatePassword(params *PutUpdatePasswordParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutUpdatePasswordOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutUpdatePasswordParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutUpdatePassword",
		Method:             "PUT",
		PathPattern:        "/update-password",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutUpdatePasswordReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutUpdatePasswordOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutUpdatePassword: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PutUpdatePrice updates price
*/
func (a *Client) PutUpdatePrice(params *PutUpdatePriceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutUpdatePriceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutUpdatePriceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutUpdatePrice",
		Method:             "PUT",
		PathPattern:        "/update-price",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutUpdatePriceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutUpdatePriceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutUpdatePrice: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
