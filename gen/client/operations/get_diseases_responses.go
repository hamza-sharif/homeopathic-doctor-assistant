// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hamza-sharif/homeopathic-doctor-assistant/gen/models"
)

// GetDiseasesReader is a Reader for the GetDiseases structure.
type GetDiseasesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDiseasesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDiseasesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDiseasesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /diseases] GetDiseases", response, response.Code())
	}
}

// NewGetDiseasesOK creates a GetDiseasesOK with default headers values
func NewGetDiseasesOK() *GetDiseasesOK {
	return &GetDiseasesOK{}
}

/*
GetDiseasesOK describes a response with status code 200, with default header values.

List of diseases
*/
type GetDiseasesOK struct {
	Payload []*models.Disease
}

// IsSuccess returns true when this get diseases o k response has a 2xx status code
func (o *GetDiseasesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get diseases o k response has a 3xx status code
func (o *GetDiseasesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get diseases o k response has a 4xx status code
func (o *GetDiseasesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get diseases o k response has a 5xx status code
func (o *GetDiseasesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get diseases o k response a status code equal to that given
func (o *GetDiseasesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get diseases o k response
func (o *GetDiseasesOK) Code() int {
	return 200
}

func (o *GetDiseasesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /diseases][%d] getDiseasesOK %s", 200, payload)
}

func (o *GetDiseasesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /diseases][%d] getDiseasesOK %s", 200, payload)
}

func (o *GetDiseasesOK) GetPayload() []*models.Disease {
	return o.Payload
}

func (o *GetDiseasesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDiseasesBadRequest creates a GetDiseasesBadRequest with default headers values
func NewGetDiseasesBadRequest() *GetDiseasesBadRequest {
	return &GetDiseasesBadRequest{}
}

/*
GetDiseasesBadRequest describes a response with status code 400, with default header values.

bad request
*/
type GetDiseasesBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this get diseases bad request response has a 2xx status code
func (o *GetDiseasesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get diseases bad request response has a 3xx status code
func (o *GetDiseasesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get diseases bad request response has a 4xx status code
func (o *GetDiseasesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get diseases bad request response has a 5xx status code
func (o *GetDiseasesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get diseases bad request response a status code equal to that given
func (o *GetDiseasesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get diseases bad request response
func (o *GetDiseasesBadRequest) Code() int {
	return 400
}

func (o *GetDiseasesBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /diseases][%d] getDiseasesBadRequest %s", 400, payload)
}

func (o *GetDiseasesBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /diseases][%d] getDiseasesBadRequest %s", 400, payload)
}

func (o *GetDiseasesBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *GetDiseasesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
