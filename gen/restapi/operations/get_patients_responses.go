// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/hamza-sharif/homeopathic-doctor-assistant/gen/models"
)

// GetPatientsOKCode is the HTTP code returned for type GetPatientsOK
const GetPatientsOKCode int = 200

/*
GetPatientsOK List of patients

swagger:response getPatientsOK
*/
type GetPatientsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Patient `json:"body,omitempty"`
}

// NewGetPatientsOK creates GetPatientsOK with default headers values
func NewGetPatientsOK() *GetPatientsOK {

	return &GetPatientsOK{}
}

// WithPayload adds the payload to the get patients o k response
func (o *GetPatientsOK) WithPayload(payload []*models.Patient) *GetPatientsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get patients o k response
func (o *GetPatientsOK) SetPayload(payload []*models.Patient) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPatientsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Patient, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetPatientsBadRequestCode is the HTTP code returned for type GetPatientsBadRequest
const GetPatientsBadRequestCode int = 400

/*
GetPatientsBadRequest bad request

swagger:response getPatientsBadRequest
*/
type GetPatientsBadRequest struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetPatientsBadRequest creates GetPatientsBadRequest with default headers values
func NewGetPatientsBadRequest() *GetPatientsBadRequest {

	return &GetPatientsBadRequest{}
}

// WithPayload adds the payload to the get patients bad request response
func (o *GetPatientsBadRequest) WithPayload(payload interface{}) *GetPatientsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get patients bad request response
func (o *GetPatientsBadRequest) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPatientsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetPatientsUnauthorizedCode is the HTTP code returned for type GetPatientsUnauthorized
const GetPatientsUnauthorizedCode int = 401

/*
GetPatientsUnauthorized internal server error

swagger:response getPatientsUnauthorized
*/
type GetPatientsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetPatientsUnauthorized creates GetPatientsUnauthorized with default headers values
func NewGetPatientsUnauthorized() *GetPatientsUnauthorized {

	return &GetPatientsUnauthorized{}
}

// WithPayload adds the payload to the get patients unauthorized response
func (o *GetPatientsUnauthorized) WithPayload(payload interface{}) *GetPatientsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get patients unauthorized response
func (o *GetPatientsUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPatientsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
