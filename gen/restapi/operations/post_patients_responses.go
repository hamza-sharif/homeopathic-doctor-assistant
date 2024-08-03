// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostPatientsCreatedCode is the HTTP code returned for type PostPatientsCreated
const PostPatientsCreatedCode int = 201

/*
PostPatientsCreated Patient added successfully

swagger:response postPatientsCreated
*/
type PostPatientsCreated struct {
}

// NewPostPatientsCreated creates PostPatientsCreated with default headers values
func NewPostPatientsCreated() *PostPatientsCreated {

	return &PostPatientsCreated{}
}

// WriteResponse to the client
func (o *PostPatientsCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// PostPatientsBadRequestCode is the HTTP code returned for type PostPatientsBadRequest
const PostPatientsBadRequestCode int = 400

/*
PostPatientsBadRequest bad request

swagger:response postPatientsBadRequest
*/
type PostPatientsBadRequest struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewPostPatientsBadRequest creates PostPatientsBadRequest with default headers values
func NewPostPatientsBadRequest() *PostPatientsBadRequest {

	return &PostPatientsBadRequest{}
}

// WithPayload adds the payload to the post patients bad request response
func (o *PostPatientsBadRequest) WithPayload(payload interface{}) *PostPatientsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post patients bad request response
func (o *PostPatientsBadRequest) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostPatientsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostPatientsUnauthorizedCode is the HTTP code returned for type PostPatientsUnauthorized
const PostPatientsUnauthorizedCode int = 401

/*
PostPatientsUnauthorized internal server error

swagger:response postPatientsUnauthorized
*/
type PostPatientsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewPostPatientsUnauthorized creates PostPatientsUnauthorized with default headers values
func NewPostPatientsUnauthorized() *PostPatientsUnauthorized {

	return &PostPatientsUnauthorized{}
}

// WithPayload adds the payload to the post patients unauthorized response
func (o *PostPatientsUnauthorized) WithPayload(payload interface{}) *PostPatientsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post patients unauthorized response
func (o *PostPatientsUnauthorized) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostPatientsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
